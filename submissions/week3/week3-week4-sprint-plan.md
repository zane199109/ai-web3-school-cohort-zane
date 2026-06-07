# Week 4 Sprint Plan — AEP (Agent Escrow Protocol)

> **学员**: zane | **单人参赛**
> **赛道**: Cobo — Agentic Economy × Agentic Wallet
> **项目**: AEP (Agent Escrow Protocol)

---

## 概述

AEP 项目已完成 Day 0–6 全链路开发（合约、后端、前端、集成测试）。Week 4 重点是 **Demo 打磨、风险场景覆盖、提交材料准备**。

| 阶段 | 目标 | 可交付 |
|------|------|--------|
| **真实实现** | CAW Custodial Wallet 全链路、AI 双轨验收、AEPReputation 声誉合约 | 可运行 Demo |
| **Mock/Fallback** | MPC TSS 模式（Custodial 保底）、FluxA 声誉聚合（推迟至 Hackathon 后） | 系统降级方案 |

---

## 每日计划

### Day 1 (Mon 6/8) — 稳定性加固

| 事项 | 类型 | 状态 |
|------|------|------|
| Docker Compose 一键启动脚本完善 | 真实实现 | 📝 |
| 所有外部依赖健康检查 (PostgreSQL/Redis/Backend) | 真实实现 | 📝 |
| SSE 重连 + 前端加载状态优化 | 真实实现 | 📝 |
| 错误边界测试（网络断线→恢复→重试） | 真实实现 | 📝 |

### Day 2 (Tue 6/9) — 风险场景覆盖

| 事项 | 类型 | 状态 |
|------|------|------|
| 超时场景：LLM 超时（>10s）→ 降级到规则引擎 | 真实实现 | 📝 |
| 并发场景：两个 Provider 同时抢单 → 双锁防重 | 真实实现 | 📝 |
| 失败场景：验收不通过 → 正确罚没 + 退款 + 声誉扣分 | 真实实现 | 📝 |
| 边界场景：大额交易 → 触发熔断拦截 | 真实实现 | 📝 |

### Day 3 (Wed 6/10) — Demo 打磨

| 事项 | 类型 | 状态 |
|------|------|------|
| Demo 演示脚本（8 分钟版，评审向） | 真实实现 | ✅ 已有初版 |
| Screenshot / GIF 录制 | 真实实现 | 📝 |
| 拓扑图美化 (ReactFlow 6 节点) | 真实实现 | 📝 |
| CAW App 审批环节演示（MPC 模式截图） | Mock | 📝 |

### Day 4 (Thu 6/11) — 提交材料

| 事项 | 类型 | 状态 |
|------|------|------|
| README 完善（含 Demo GIF、架构图、验证清单） | 真实实现 | ✅ 已有初版 |
| 提交页面截图 / 交易哈希收集 | 真实实现 | 📝 |
| 赛道对齐说明（Cobo） | 真实实现 | ✅ 已提交 |
| Sponsor/Mentor 问答案例 | 真实实现 | ✅ 已提交 |

### Day 5 (Fri 6/12) — 总验收

| 事项 | 类型 | 状态 |
|------|------|------|
| Foundry 测试全过 | 真实实现 | ✅ 8/8 |
| Go 集成测试全过 | 真实实现 | ✅ 10/10 |
| 评审 Demo 预演 | 真实实现 | 📝 |
| 所有提交材料最终检查 | 真实实现 | 📝 |

---

## Fallback 计划

| 功能 | 真实路径 | Fallback |
|------|---------|----------|
| CAW Wallet | Custodial（真实 API 调用） | MPC TSS（环境不通时降级） |
| 链上合约 | AEPBounty.sol on Anvil | Base Sepolia 部署（时间不够则本地演示） |
| 验收引擎 | DeepSeek LLM 真实调用 | 纯规则引擎（API Key 问题） |
| 前端 | React + MUI | 静态 HTML（紧急降级） |
