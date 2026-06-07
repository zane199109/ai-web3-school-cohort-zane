# Sponsor SDK / API Integration Plan — CAW 集成方案

> **学员**: zane
> **项目**: AEP (Agent Escrow Protocol)

---

## 接入什么 — Cobo Agentic Wallet (CAW)

| 接入项 | 用途 | 状态 |
|--------|------|------|
| CAW Custodial Wallet | Buyer/Seller/Evaluator 三 Agent 资金托管 | ✅ 已实现 |
| CAW Pact | 资金锁定、策略预检、Release 结算 | ✅ 已实现 |
| CAW REST API | 钱包创建、地址生成、转账、查询 | ✅ 已实现 |

## 怎么接

### 架构集成

```
Go Backend → HTTP Client → api-core.agenticwallet.dev.cobo.com/api/v1
                ↓
           X-API-Key: <api_key>
                ↓
           JSON Request/Response
```

### 关键端点

| 操作 | 端点 | 参数 |
|------|------|------|
| 创建钱包 | `POST /api/v1/wallets` | `wallet_type: Custodial` |
| 创建 Pact | `POST /api/v1/wallets/{wallet_id}/pacts` | `amount, valid_until, approvers` |
| Release | `POST /api/v1/wallets/{wallet_id}/pacts/{pact_id}/release` | `amount, receiver_address` |
| 查询余额 | `GET /api/v1/wallets/{wallet_id}` | — |
| 地址转账 | `POST /api/v1/transfers` | `from, to, amount` |

### 认证方式

- Header: `X-API-Key: <api_key>`（简单认证，无 HMAC 签名）
- 相关 CAW 配置在 `conf/config.yaml` 中管理

## Week 4 是否能做完

| 功能 | 状态 | Week 4 可交付 |
|------|------|--------------|
| Custodial Wallet 全链路 | ✅ 已完成 | ✅ 可演示 |
| MPC TSS Node | ⚠️ 环境依赖 | 🟡 截图展示, Custodial 保底 |
| Pact Policy 策略配置 | ✅ 已完成 | ✅ 可演示 |
| Release 结算 + 重试 | ✅ 已完成 | ✅ 可演示 |
| 多钱包并发管理 | ✅ 已完成 | ✅ 可演示 |

## 如果接不通的 Fallback

| 失败场景 | Fallback |
|---------|---------|
| CAW API 不可达 | 本地 SimpleCAW Mock（SimpleCAW.sol + Mock CAW Client） |
| Custodial 钱包创建失败 | 使用预创建钱包, 跳过创建步骤 |
| Pact Release 失败 | Relayer 重试机制（最多 3 次）+ 管理员手动 Retry API |
| MPC TSS 网络断连 | Custodial 模式全线保底 |
| API Key 配额耗尽 | 本地 Mock 模式跑通后再申请新 Key |
