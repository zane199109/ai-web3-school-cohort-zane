package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"time"

	"x402-caw-demo/caw"

	"github.com/joho/godotenv"
)

const (
	HeaderPaymentAddress = "X-Payment-Address"
	HeaderPaymentAmount  = "X-Payment-Amount"
	HeaderPaymentProof   = "X-Payment-Proof"
)

// 执行 cast send 调用链上 SimpleCAW 合约
func executePaymentViaCast(cawAddr, recipient, amount, ownerPK string) (string, error) {
	cmd := exec.Command(
		"cast", "send",
		cawAddr,
		"executePayment(address,uint256)",
		recipient, amount,
		"--private-key", ownerPK,
		"--rpc-url", "http://127.0.0.1:8545",
		"--json",
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("cast error: %v, output: %s", err, string(output))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(output, &result); err == nil {
		if hash, ok := result["transactionHash"].(string); ok {
			return hash, nil
		}
	}
	return "", fmt.Errorf("failed to parse tx hash from cast output: %s", string(output))
}

// 执行一个场景的完整 x402 流程
func runScenario(apiURL string, client *http.Client, localPolicy caw.Policy, cawAddr, ownerPK string) {
	fmt.Printf("\n━━━ 🔄 场景: %s ━━━\n", apiURL)

	// 1. 发起初始请求
	fmt.Println("👉 Agent: 请求 API 数据...")
	req, _ := http.NewRequest("GET", apiURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ 请求失败: %v\n", err)
		return
	}

	// 2. 识别 HTTP 402
	if resp.StatusCode == http.StatusPaymentRequired {
		fmt.Println("🛑 Agent: 遭遇 402 Paywall！解析付款要求...")

		payee := resp.Header.Get(HeaderPaymentAddress)
		amountStr := resp.Header.Get(HeaderPaymentAmount)
		resp.Body.Close()

		// 将金额字符串转为 big.Int
		reqAmount := new(big.Int)
		reqAmount.SetString(amountStr, 10)

		// 3. 调用本地 CAW 引擎进行预检审计
		approved, audit := caw.CheckPolicy(reqAmount, payee, localPolicy)
		caw.PrintAudit(audit)

		if !approved {
			fmt.Println("🛑 CAW 本地拦截成功！拒绝支付，节省 Gas。")
			return
		}

		// 4. 链下预检通过，调用链上 SimpleCAW 合约执行硬性扣款
		fmt.Println("⚙️ Agent: CAW 审批通过，正在执行链上支付...")
		txHash, err := executePaymentViaCast(cawAddr, payee, amountStr, ownerPK)
		if err != nil {
			fmt.Printf("❌ 链上执行失败 (合约策略拦截): %v\n", err)
			return
		}
		fmt.Printf("✅ Agent: 链上支付成功！TxHash: %s\n", txHash)

		// 5. 带上 TxHash 作为支付凭证重试请求
		fmt.Println("👉 Agent: 携带支付凭证重新请求 API...")
		reqRetry, _ := http.NewRequest("GET", apiURL, nil)
		reqRetry.Header.Set(HeaderPaymentProof, txHash)

		respRetry, err := client.Do(reqRetry)
		if err != nil {
			fmt.Printf("❌ 重试请求失败: %v\n", err)
			return
		}
		defer respRetry.Body.Close()

		bodyRetry, _ := io.ReadAll(respRetry.Body)
		if respRetry.StatusCode == http.StatusOK {
			fmt.Printf("🎉 Agent: 成功获取数据！%s\n", string(bodyRetry))
		} else {
			fmt.Printf("❌ Agent: 付款后仍获取失败: %s\n", string(bodyRetry))
		}
	} else {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("❌ Agent: 期望 402，收到 %d: %s\n", resp.StatusCode, string(body))
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file. Did you run setup.sh?")
	}

	cawAddr := os.Getenv("CAW_ADDR")
	ownerPK := os.Getenv("OWNER_PK")
	providerAddr := os.Getenv("PROVIDER_ADDR")

	// 初始化本地 CAW 策略引擎 (与链上 SimpleCAW 合约保持一致)
	// 100 USDC * 10^6 = 每日最多 100 USDC (USDC decimals=6)
	localPolicy := caw.Policy{
		MaxBudgetPerDay:  new(big.Int).Mul(big.NewInt(100), new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)), // 100 * 10^6
		AllowedRecipient: providerAddr,
	}

	// 三个测试场景
	scenarios := []struct {
		Name string
		URL  string
	}{
		{"合法请求 — 10 USDC 给白名单地址", "http://127.0.0.1:8081/api/success"},
		{"钓鱼攻击 — 10 USDC 给未授权地址", "http://127.0.0.1:8081/api/phishing"},
		{"超额请求 — 200 USDC 超出日预算", "http://127.0.0.1:8081/api/overlimit"},
	}

	fmt.Println("=======================================")
	fmt.Println("🤖 x402-CAW Agent 自动轮询演示启动")
	fmt.Println("=======================================")
	fmt.Println("即将依次执行 3 个场景：")
	for i, s := range scenarios {
		fmt.Printf("  %d. %s\n", i+1, s.Name)
	}
	fmt.Println()

	client := &http.Client{Timeout: 30 * time.Second}

	for i, s := range scenarios {
		fmt.Printf("\n═══════ 场景 %d/%d ═══════\n", i+1, len(scenarios))
		fmt.Printf("📋 %s\n", s.Name)
		runScenario(s.URL, client, localPolicy, cawAddr, ownerPK)
		time.Sleep(1 * time.Second) // 场景间间隔
	}

	fmt.Println("\n=======================================")
	fmt.Println("✅ 全部场景执行完毕！")
	fmt.Println("=======================================")
}
