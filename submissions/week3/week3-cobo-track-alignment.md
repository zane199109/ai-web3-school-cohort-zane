# Cobo 赛道对齐任务 — AEP 的 CAW 集成方案

> **赛道**: Cobo — Agentic Economy × Agentic Wallet
> **项目**: AEP (Agent Escrow Protocol)
> **学员**: zane

---

## AI Agent 如何在可控边界内持有钱包、管理预算、执行支付/交易/资源采购

### 1. 钱包持有架构

AEP 采用 **三 Agent 独立钱包架构**：

```
Buyer  Agent → CAW Custodial Wallet (资金源, 创建 Bounty Pact)
Seller Agent → CAW Custodial Wallet (收款账户, 接收结算)
Eval   Agent → CAW Custodial Wallet (Gas 开销, 转发裁决结果)
```

每个 Agent 持有独立钱包，通过 CAW REST API (`api-core.agenticwallet.dev.cobo.com/api/v1`) 管理。钱包创建通过 `X-API-Key` Header 认证，不支持跨钱包互操作 — Buyer 无法操作 Seller 的资金。

### 2. 预算管理

**三层分层额度控制**：

| 层级 | 实现方式 | 作用 |
|------|---------|------|
| L1: 链下策略预检 | CAW Pact Policy | 检查白名单合约地址、单笔限额、日预算上限 |
| L2: 链上强制执行 | AEPBounty.sol 状态机 | 资金锁定在合约内，无法被单方提取 |
| L3: 人类审批兜底 | CAW App (MPC 模式) | MPC 2-of-2 签名，Buyer 必须在 App 确认后才能 Release |

**预算配置方式**：在创建 Pact 时指定 `amount` 和 `valid_until`，CAW 自动做预算检查。超额度请求在链下被 Pact Policy 拦截，无需 Gas。

### 3. 支付/交易执行流程

```
POST /api/bounty (创建 Bounty)
  → CAW: POST /api/v1/wallets/{wallet_id}/pacts (锁定资金)
  → Pact Status: pending_approval → active (Custodial 自动激活)
  → Contract: AEPBounty.sol.claim() (状态机 Open → Funded)

POST /api/bounty/{id}/submit (提交交付物)
  → IPFS: 上传文件 → 获取 CID
  → Rule Engine: 字段校验 + 空数据检测
  → LLM: 相关性 + 幻觉检测 (超时 10s 降级到规则引擎)

POST /api/confirm/{jobId} (Buyer 确认)
  → BuyerApproval 检查 (必须为 true)
  → CAW: Release Pact → 放款给 Seller
  → Contract: AEPReputation.sol.updateScore() (+10/-20)
```

### 4. 风险边界记录

| 风险检查点 | 触发条件 | 动作 |
|-----------|---------|------|
| 白名单检查 | 目标合约不在白名单 | 拒绝交易, 记录日志 |
| 单笔限额 | `amount > max_per_tx` | CAW Pact Policy 拦截 |
| 日预算 | 当日累计 > `daily_budget` | CAW Pact Policy 拦截 |
| BuyerApproval | `confirm=false` | 资金永不 Release, 过期可退款 |
| LLM 超时 | `response > 10s` | 降级到规则引擎 |
| 结算失败 | Network Error / Revert | 重试 (最多 3 次, 指数退避) |

## 关键安全设计原则

1. **AI 不能自己决定花钱** — 每次资金操作都需要 Buyer 通过 CAW App 审批
2. **资金锁在合约内** — 无法被单方提走, 只有 AEPBounty.sol 状态机可触发 Release 或 Refund
3. **验收引擎无权触达资金** — Evaluator 只能输出裁决信号, 资金操作由 Relayer 通过 CAW 完成
4. **全链路可审计** — 每个步骤都有 Zap 日志 + 链上事件 + CAW Pact 记录
