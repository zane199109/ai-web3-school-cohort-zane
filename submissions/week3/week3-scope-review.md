# Scope Review — AEP 不做什么

> 明确 Week 4 砍掉、延后或 Mock 的功能清单，避免范围膨胀。

---

## 砍掉的功能（Week 4 不做）

| # | 功能 | 原因 | 替代方案 |
|---|------|------|---------|
| 1 | **FluxA 声誉聚合** | 需要 FluxA 合约部署和跨协议对接，AEPReputation 已独立运行 | AEPReputation.sol 极简版（+10/-20） |
| 2 | **MPC TSS 多链部署** | CAW Custodial 模式已可跑通全链路，TSS 环境依赖复杂 | Custodial Wallet 保底，MPC 截图展示 |
| 3 | **多 Evaluator 争议仲裁** | 多 Oracle 投票仲裁增加复杂度，Demo 单 Evaluator 即可展示核心流程 | 单 Evaluator + 人类审批兜底 |
| 4 | **Staking 质押经济模型** | 声誉代币的经济模型设计超出 MVP 范围 | 声誉仅用作门槛准入，不做经济博弈 |

## 延后的功能（Week 4 之后）

| # | 功能 | 原计划 | 新时间 |
|---|------|--------|--------|
| 1 | Base Sepolia 正式部署 | Week 4 | Week 4 可选，本地 Anvil 保底 |
| 2 | 多 Provider 竞价机制 | Week 4 | Hackathon 后 |
| 3 | 通知系统（邮件/Discord） | Week 4 | Hackathon 后 |
| 4 | Dashboard 分析看板 | Week 4 | Hackathon 后 |

## Mock 的功能

| # | 功能 | Mock 方案 |
|---|------|----------|
| 1 | CAW MPC TSS Node | Custodial Wallet（真实 API，非 Mock） |
| 2 | 多链部署 | 单链（Anvil）演示 |
| 3 | 生产级错误处理 | 日志 + 重试（最多 3 次） |

## 核心聚焦

Week 4 只做三件事：
1. **Demo 打磨** — 演示脚本、GIF 录制、拓扑图美化
2. **场景覆盖** — 超时降级、并发抢单、失败回滚
3. **提交材料** — README、截图、验证清单
