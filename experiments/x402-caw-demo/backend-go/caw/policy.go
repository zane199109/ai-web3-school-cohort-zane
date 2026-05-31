package caw

import (
	"fmt"
	"math/big"
)

// Policy 模拟 Cobo CAW 中的 Pact/Policy 规则
type Policy struct {
	MaxBudgetPerDay  *big.Int
	AllowedRecipient string
}

// AuditLog 模拟 Cobo CAW 的可审计记录
type AuditLog struct {
	Action    string
	Reason    string
	Amount    *big.Int
	Recipient string
}

// CheckPolicy 模拟 Agent 发起交易前，CAW 的链下策略评估
func CheckPolicy(reqAmount *big.Int, reqRecipient string, policy Policy) (bool, AuditLog) {
	audit := AuditLog{Amount: reqAmount, Recipient: reqRecipient}

	// 规则 1：收款方白名单检查
	if reqRecipient != policy.AllowedRecipient {
		audit.Action = "REJECT"
		audit.Reason = "Recipient not in CAW whitelist"
		return false, audit
	}

	// 规则 2：每日预算上限检查
	if reqAmount.Cmp(policy.MaxBudgetPerDay) > 0 {
		audit.Action = "REJECT"
		audit.Reason = "Exceeds CAW daily budget limit"
		return false, audit
	}

	// 审批通过
	audit.Action = "APPROVE"
	audit.Reason = "Within policy bounds"
	return true, audit
}

// PrintAudit 打印审计日志
func PrintAudit(audit AuditLog) {
	fmt.Printf("🛡️ CAW Audit: Action=%s, Reason=%s, Amount=%s, Recipient=%s\n",
		audit.Action, audit.Reason, audit.Amount.String(), audit.Recipient)
}
