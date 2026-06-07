# 技术验证计划 — Week 4 关键验证点

> **学员**: zane
> **项目**: AEP (Agent Escrow Protocol)

---

## 已完成的验证

| # | 验证项 | 通过标准 | 状态 |
|---|--------|---------|------|
| 1 | Foundry 合约单元测试 | AEPBounty 状态机 8/8 测试全通过 | ✅ |
| 2 | Go 集成测试 | Demo 全流程 10/10 测试全通过 | ✅ |
| 3 | CAW Pact 创建与 Release | Custodial 模式 Pact 自动激活, Release 成功 | ✅ |
| 4 | AI 双模验收引擎 | 规则引擎硬校验 + LLM 旁轨抽样 | ✅ |

## Week 4 要验证的关键技术点

### 1. Agent Trace (调用链追踪)

| 验证项 | 方法 | 通过标准 |
|--------|------|---------|
| Trace ID 贯穿全流程 | 检查 Go 后端 Zap 日志 | 每个请求有唯一 Trace ID |
| SSE 事件流完整 | 前端监听 `/api/events` | 拓扑图 6 节点依次触发 |
| CAW 操作日志 | 检查 CAW Pact 记录 | Pact ID 与 Job ID 一一对应 |

### 2. SDK 调用 (CAW REST API)

| 验证项 | 方法 | 通过标准 |
|--------|------|---------|
| Custodial 钱包创建 | POST `/api/v1/wallets` | 返回 wallet_id, status=active |
| Pact 创建 + 策略检查 | POST `/api/v1/wallets/{id}/pacts` | 返回 pact_id, 金额正确 |
| Release 结算 | POST `/api/v1/wallets/{id}/pacts/{pid}/release` | 余额变更符合预期 |

### 3. 测试网交易

| 验证项 | 方法 | 通过标准 |
|--------|------|---------|
| Anvil 合约部署 | `forge script DeployAll` | 返回合约地址, 非 0x0 |
| 合约状态读写 | `cast call/send` | 状态机转换正确 |
| 交易哈希可查 | Base Sepolia Scan | 交易确认 > 1 个区块 |

### 4. 合约交互

| 验证项 | 方法 | 通过标准 |
|--------|------|---------|
| AEPBounty 状态机 | Foundry test | 8/8 全过 |
| AEPReputation 读写 | Foundry test | getScore, updateScore 正确 |
| onlyBounty 权限控制 | Foundry test | 非 Bounty 地址调用被 Revert |

### 5. 权限控制

| 验证项 | 方法 | 通过标准 |
|--------|------|---------|
| Evaluator 白名单 | 配置验证 | 只有白名单地址可做裁决 |
| BuyerApproval 检查 | 集成测试 | `confirm=false` 时不 Release |
| Admin Retry 权限 | 接口验证 | 携带 Admin Token 才允许操作 |

### 6. 日志记录

| 验证项 | 方法 | 通过标准 |
|--------|------|---------|
| Zap 日志全链路 | `grep -c "trace_id" logs/` | 每个请求有完整日志链 |
| 错误日志告警 | 故意触发错误 | 日志记录错误类型 + 堆栈 |
| 事件持久化 | PG `events` 表查询 | 事件记录完整, 可追溯 |

### 7. Demo 截图/录制

| 验证项 | 方法 | 通过标准 |
|--------|------|---------|
| 拓扑图 6 节点 | 浏览器截图 | SSE 连接 🟢 每个节点状态正确 |
| 分步引导流 | GIF 录制 | 5 步 Step 引导完整展示 |
| CAW App 审批 | 截图（MPC 模式） | 展示人类审批环节 |
| 链上交易哈希 | 浏览器截图 | 交易确认且可在 Etherscan 查询 |

## 验证优先级

```
P0（Demo 必需）: Agent Trace → 合约交互 → Demo 截图
P1（安全展示）: 权限控制 → SDK 调用
P2（完整性）  : 测试网交易 → 日志记录
```
