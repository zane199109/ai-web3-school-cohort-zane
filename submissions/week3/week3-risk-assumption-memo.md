# Risk / Assumption Memo — AEP (Agent Escrow Protocol)

> **学员**: zane
> **项目依赖前提、最大失败风险与 Fallback Plan**

---

## 项目成立依赖的前提

| 前提 | 依赖等级 | 验证方式 |
|------|---------|---------|
| LLM API (DeepSeek) 可用且响应 < 10s | 🔴 关键 | Go 集成测试已验证 |
| CAW REST API 可用（Custodial 模式） | 🔴 关键 | 真实 Pact 创建 + Release 已验证 |
| PostgreSQL + Redis 可运行（Docker） | 🟡 重要 | docker compose up 已验证 |
| Anvil 本地链可启动 | 🟡 重要 | Foundry 测试已验证 |

## 最可能失败在哪里

| 失败模式 | 概率 | 影响 | 详细分析 |
|---------|------|------|---------|
| DeepSeek LLM API Key 失效或欠费 | 中 | 高 | 使用 Hermes 内部代理 Key 不可用于外部调用。需 `sk-` 开头正式 Key。**缓解**：LLM 超时 10s 自动降级到规则引擎，Demo 不受影响 |
| CAW TSS MPC 网络延迟 | 中 | 中 | TSS Node 需要 WebSocket 连接 CAW Relay，WSL 环境网络不稳定。**缓解**：已切换到 Custodial 模式，秒级创建 |
| Docker 镜像拉取超时 | 低 | 高 | WSL 网络有时不通。**缓解**：使用 Windows 代理 172.21.80.1:7897 |
| React 构建包体积过大 | 低 | 低 | MUI 导致 bundle 较大 (~470KB)。**缓解**：使用 dist 静态文件 |

## Fallback Plan

| 失败场景 | Fallback |
|---------|---------|
| DeepSeek 不可用 | 纯规则引擎（字段校验 + 空数据检测）足以跑通 Happy Path |
| CAW API 不可用 | 本地 SimpleCAW Mock（已保留在实验代码中） |
| PostgreSQL/Redis 不通 | SQLite 代替 PG，内存 Map 代替 Redis（需修改 store 层） |
| 前端构建超时 | 使用 dist 中已构建的静态 HTML |
| 合约部署失败 | Anvil 本地链测试通过即可，Demo 本地演示 |

## 应急决策树

```
LLM 验收失败?
├─ 超时 (>10s) → 降级到规则引擎 → 继续流程
├─ 返回错误 → 重试 1 次 → 降级到规则引擎
└─ 返回乱码 → 标记 Uncertain → 人工审核

CAW 结算失败?
├─ 网络错误 → 重试 (最多 3 次, 指数退避)
├─ 余额不足 → 报错提示充 Gas
└─ Pact 状态异常 → 管理员手动 Retry (Admin API)

合约交易失败?
├─ Gas 不够 → Relayer 补 Gas 重试
├─ Revert → 回滚 DB 状态, 记录错误日志
└─ Nonce 冲突 → Nonce 管理 + 重试
```
