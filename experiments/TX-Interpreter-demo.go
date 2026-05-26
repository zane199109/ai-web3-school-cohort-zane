package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 配置
const (
	ethRPC     = //填入你的RPC URL
	llmAPIURL  = //填入你的LLM URL
	llmAPIKey  = //填入你的LLM API KEY
	llmModel   = //填入你的LLM MODEL NAME
	scanABIURL = //填入你的Etherscan ABI URL
	scanAPIKey = // 填入你的Etherscan API Key
)

// 链上原始数据
type TxOnChainData struct {
	TxHash        string
	From          common.Address
	To            *common.Address
	Value         string
	TokenValue    string
	InputData     string
	Status        uint64
	GasUsed       uint64
	Logs          []*types.Log
	FuncSignature string
	FuncName      string
	Events        []string
	ContractABI   abi.ABI // 新增合约ABI实例
	ABIRaw        string  // 原始ABI字符串
}

type TxExplain struct {
	UserAction   string   `json:"user_action"`
	Assets       []string `json:"assets"`
	Addresses    []string `json:"addresses"`
	OnChainFacts []string `json:"on_chain_facts"`
	Inferences   []string `json:"inferences"`
	Uncertains   []string `json:"uncertains"`
	CheckList    []string `json:"check_list"`
}

func main() {

	fmt.Println("======= 交易解释器（可重复输入）=======")
	fmt.Println("输入交易哈希解析，输入 exit 退出")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("请输入交易哈希: ")
		if !scanner.Scan() {
			break
		}

		txHash := strings.TrimSpace(scanner.Text())
		if strings.ToLower(txHash) == "exit" {
			fmt.Println("退出程序")
			break
		}

		if len(txHash) != 66 || !strings.HasPrefix(txHash, "0x") {
			fmt.Println("❌ 无效哈希")
			continue
		}

		// txHash := "0x3b207559c44f0b3b6634fce79ed649e58bf9160f52c668644619f4387afce1cf"
		// txHash := "0x62e89e8c1238535329afe0cd3542f825cdce601888064e7678e2de5e8c689b16"
		// txHash := "0xb3c69eec517d4c092ad5e8b7ef07471536949a8df86e3499f7855666bd204f07"

		// 执行解析
		ctx := context.Background()
		data, err := fetchTxData(ctx, txHash)
		if err != nil {
			fmt.Println("❌ 获取失败:", err)
			continue
		}

		fmt.Printf("✅ 函数: %s | ETH: %s | 代币: %s\n",
			data.FuncName, data.Value, data.TokenValue)

		explain, err := callLLM(data)
		if err != nil {
			fmt.Println("❌ LLM失败:", err)
			continue
		}

		printResult(explain)
	}
}

// func main() {
// 	// txHash := "0x3b207559c44f0b3b6634fce79ed649e58bf9160f52c668644619f4387afce1cf"
// 	// txHash := "0x62e89e8c1238535329afe0cd3542f825cdce601888064e7678e2de5e8c689b16"
// 	txHash := "0xb3c69eec517d4c092ad5e8b7ef07471536949a8df86e3499f7855666bd204f07"
// 	ctx := context.Background()

// 	data, err := fetchTxData(ctx, txHash)
// 	if err != nil {
// 		fmt.Println("获取交易失败:", err)
// 		return
// 	}

// 	// fmt.Println("✅ 调用函数:", data.FuncName)
// 	// fmt.Println("💰 原生ETH:", data.Value)
// 	// fmt.Println("💰 代币金额:", data.TokenValue)
// 	fmt.Println("📄 是否获取到ABI:", len(data.ABIRaw) > 0)

// 	explain, err := callLLM(data)
// 	if err != nil {
// 		fmt.Println("LLM 解释失败:", err)
// 		return
// 	}
// 	printResult(explain)
// }

func fetchTxData(ctx context.Context, txHash string) (*TxOnChainData, error) {
	client, err := ethclient.DialContext(ctx, ethRPC)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	tx, isPending, err := client.TransactionByHash(ctx, common.HexToHash(txHash))
	if err != nil {
		return nil, err
	}
	if isPending {
		return nil, fmt.Errorf("交易pending")
	}

	receipt, err := client.TransactionReceipt(ctx, common.HexToHash(txHash))
	if err != nil {
		return nil, err
	}

	from, _ := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)

	// 解析函数签名
	input := tx.Data()
	funcSig, funcName := "unknown", "unknown"
	if len(input) >= 4 {
		sigHex := "0x" + common.Bytes2Hex(input[:4])
		funcSig = sigHex
		funcName = lookupFunction(sigHex)
	}

	// 解析事件列表
	var events []string
	tokenValue := "0"
	for _, log := range receipt.Logs {
		if len(log.Topics) > 0 {
			events = append(events, log.Topics[0].Hex())
			val := new(big.Int).SetBytes(log.Data)
			tokenValue = val.String()
		}
	}

	// 拉取合约ABI
	var contractABI abi.ABI
	var abiRaw string
	if tx.To() != nil {
		abiRaw, err = fetchContractABI(tx.To().Hex())
		if err == nil && abiRaw != "" {
			contractABI, _ = abi.JSON(bytes.NewReader([]byte(abiRaw)))
		}
	}

	return &TxOnChainData{
		TxHash:        txHash,
		From:          from,
		To:            tx.To(),
		Value:         tx.Value().String(),
		TokenValue:    tokenValue,
		InputData:     common.Bytes2Hex(tx.Data()),
		Status:        receipt.Status,
		GasUsed:       receipt.GasUsed,
		Logs:          receipt.Logs,
		FuncSignature: funcSig,
		FuncName:      funcName,
		Events:        events,
		ContractABI:   contractABI,
		ABIRaw:        abiRaw,
	}, nil
}

// 从Etherscan拉取合约ABI
func fetchContractABI(addr string) (string, error) {
	url := fmt.Sprintf("%s?module=contract&action=getabi&address=%s&apikey=%s", scanABIURL, addr, scanAPIKey)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Result  string `json:"result"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}
	if res.Status != "1" {
		return "", fmt.Errorf("获取ABI失败")
	}
	return res.Result, nil
}

func lookupFunction(sigHex string) string {

	// url := "https://sig.eth.samczsun.com/api/v1/signatures?function=" + sigHex
	url := fmt.Sprintf("https://api.openchain.xyz/signature-database/v1/lookup?function=%s", sigHex)
	// url := fmt.Sprintf("https://sig.eth.samczsun.com/api/v1/signatures?function=%s", sigHex)
	resp, err := http.Get(url)
	if err != nil {
		return "unknown"
	}
	defer resp.Body.Close()

	var result struct {
		Result struct {
			Function map[string][]struct {
				Name string `json:"name"`
			} `json:"function"`
		} `json:"result"`
	}
	_ = json.NewDecoder(resp.Body).Decode(&result)
	if list, ok := result.Result.Function[sigHex]; ok && len(list) > 0 {
		return list[0].Name
	}
	return "unknown"
}

func callLLM(data *TxOnChainData) (*TxExplain, error) {
	prompt := fmt.Sprintf(`
你是区块链交易解释器，只输出标准JSON。
请根据以下信息解释用户做了什么动作，涉及哪些资产和地址，哪些信息来自链上数据，
哪些是模型推断，模型不确定的地方，如果要签类似交易，用户应该检查什么，必须保证有内容输出，不能为空，不返回空数组。
输出字段：
user_action, on_chain_facts, inferences, uncertains, check_list
链上数据：
哈希：%s
发送方：%s
接收方：%s
调用函数：%s
原生ETH金额：%s
代币金额：%s
交易状态：%d
合约ABI概要：%s
`,
		data.TxHash,
		data.From.Hex(),
		addrToStr(data.To),
		data.FuncName,
		data.Value,
		data.TokenValue,
		data.Status,
		data.ABIRaw[:min(300, len(data.ABIRaw))],
	)

	body := map[string]interface{}{
		"model": llmModel,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
		"response_format": map[string]string{"type": "json_object"},
	}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", llmAPIURL, bytes.NewBuffer(bodyBytes))
	req.Header.Set("Authorization", "Bearer "+llmAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var llmResp struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			}
		}
	}
	_ = json.NewDecoder(resp.Body).Decode(&llmResp)
	if len(llmResp.Choices) == 0 {
		return nil, fmt.Errorf("llm返回空")
	}

	var exp TxExplain
	_ = json.Unmarshal([]byte(llmResp.Choices[0].Message.Content), &exp)
	return &exp, nil
}

func addrToStr(addr *common.Address) string {
	if addr == nil {
		return "contract deployment"
	}
	return addr.Hex()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const count = 10

func printResult(e *TxExplain) {
	fmt.Println("\n======= 🧾 交易解释结果 =======")
	fmt.Println("\n【用户动作】")
	fmt.Println(e.UserAction)
	fmt.Println("\n【✅ 链上事实（100%可信）】")
	for _, v := range e.OnChainFacts {
		fmt.Println("-", v)
	}
	fmt.Println("\n【🤖 模型推断（可能错误）】")
	for _, v := range e.Inferences {
		fmt.Println("-", v)
	}
	fmt.Println("\n【⚠️ 模型不确定】")
	for _, v := range e.Uncertains {
		fmt.Println("-", v)
	}
	fmt.Println("\n【🔴 签名前必须检查】")
	fmt.Println(count + 1)
	for _, v := range e.CheckList {
		fmt.Println("-", v)
	}

}
