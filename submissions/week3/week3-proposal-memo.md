# Proposal Memo — AEP (Agent Escrow Protocol)

> **学员**: zane
> **日期**: 2026-06-07
> **赛道**: Cobo — Agentic Economy × Agentic Wallet

---

## 目标用户

- AI Agent 开发者（需链上数据/算力的 Agent）
- 去中心化数据服务提供方
- DAO 采购方（需要先验货后付款的链上采购）
- Agent-to-Agent 商业结算基础设施需求方

## 真实场景

一个交易员 Agent 购买链上数据分析 Agent 的"巨鲸监控报告"——买方 Agent 锁定 0.1 USDC 到托管合约，卖方提交报告后由 AEP 验收：格式正确且内容相关则放款，返回空数据或废话则罚没退款。Buyer 通过 CAW App 审批确认放款。

## 最小功能

| # | 功能 | 状态 |
|---|------|------|
| 1 | ERC-8183 状态机：Open → Funded → Submitted → Complete/Reject | ✅ |
| 2 | AEP 双模验收引擎：Rule（格式/空数据硬校验）+ LLM（相关性/幻觉检测） | ✅ |
| 3 | CAW Pact 资金锁定与 Release 结算 | ✅ |
| 4 | AEPReputation.sol 声誉合约（ERC-8004 极简版，+10/-20） | ✅ |
| 5 | MUI 前端演示面板（5 步引导流） | ✅ |
| 6 | 全链路集成测试（Foundry 8/8 + Go 10/10） | ✅ |

## 验证方式

- Anvil 本地链 + Go 后端 + 三场景 Demo
- 场景 1：合法交付 → APPROVE → 链上支付 → 声誉 +1
- 场景 2：恶意提交 → REJECT → 罚没 → 声誉 -2
- 场景 3：超额请求 → 链下 CAW 策略拦截（0 Gas）
- 全流程通过 Foundry 测试 + Go 集成测试自动验证

## 风险边界

| 风险 | 缓解 |
|------|------|
| LLM 幻觉放行次品 | 规则层硬过滤 + LLM 仅抽样，高额人工 |
| TSS 网络延迟 | Custodial 保底 + 重试 |
| Prompt Injection | Evaluator 白名单 + 链下预检 |
| 合约漏洞 | Foundry 全覆盖 + 紧急退款机制 |

## 可能赛道

- AI × Web3 Hackathon（Cobo 赛道）
- Agent-to-Agent 商业结算基础设施
- 去中心化数据市场 / Agent 经济中间件
