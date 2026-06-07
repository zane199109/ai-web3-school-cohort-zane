# Sponsor / Mentor 问题清单

> 整理自 Week 3 开发过程中遇到的问题，用于向 Sponsor/Mentor 请教

---

## 1. Cobo

| 项目 | 内容 |
|------|------|
| **问题** | CAW Custodial 钱包创建后，如何确认钱包地址的所属权？Custodial 模式下是否有类似 MPC 的签名验证机制？ |
| **背景/卡点** | AEP 项目中 Buyer/Seller/Evaluator 使用独立 Custodial Wallet。目前通过 `wallet_id` 管理，但不确定 Custodial 钱包的链上地址签名是否能被合约或第三方验证为"该钱包拥有者发起的操作"。 |
| **希望回答** | Custodial 钱包的地址签名验证流程是什么？是否支持 EIP-1271 合约级签名验证？ |

## 2. Cobo

| 项目 | 内容 |
|------|------|
| **问题** | CAW Pact 的 `approve_pending_pacts` 在 Custodial 模式下是否会自动通过？如何模拟 MPC 模式的人类审批环节来增强 Demo 的可信叙事？ |
| **背景/卡点** | Demo 流程中需要展示"人类 CAW App 审批"这个关键安全环节，但 Custodial 模式下 Pact 自动激活，缺少这个交互节点。 |
| **希望回答** | Custodial 模式是否可以通过 API 模拟审批延迟？或者在 MPC 模式下是否可以在 Demo 环境真实展示 App 推送审批流程？ |

## 3. Cobo

| 项目 | 内容 |
|------|------|
| **问题** | Go 语言通过 REST API 调用 CAW 时，对于大量并发请求的 Rate Limit 和并发安全是否有官方建议？ |
| **背景/卡点** | AEP 中可能有多个 Agent 同时发起交易（如 Evaluator 裁决后批次结算）。需要了解 CAW API 的最佳并发策略。 |
| **希望回答** | CAW API 是否有 Rate Limit？Go 客户端的最佳并发数和重试策略建议？ |
