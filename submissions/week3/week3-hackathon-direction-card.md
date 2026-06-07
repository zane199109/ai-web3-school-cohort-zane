# Hackathon Direction Card — AEP (Agent Escrow Protocol)

> **学员**: zane
> **项目**: AEP (Agent Escrow Protocol)
> **赛道**: Cobo — Agentic Economy × Agentic Wallet
> **日期**: 2026-06-07

---

## 参赛赛道

**Cobo｜Agentic Economy × Cobo Agentic Wallet**

项目使用 Cobo Agentic Wallet (CAW) 作为资金托管和结算载体，实现 AI Agent 间的可信商业交易。

## 项目名

**AEP (Agent Escrow Protocol)** — 链上担保交易的 AI 验收引擎

## 目标用户

- AI Agent 开发者（需要链上数据或算力的 Agent）
- 去中心化数据服务提供方
- DAO 采购方（需先验货后付款的链上采购）
- Agent-to-Agent 商业结算需求方

## 要解决的问题

**AI Agent 之间的商业交易缺乏可信的中立验收机制。**

- 买方 Agent 付款后，卖方可能交付空数据或次品
- AI 缺乏链上身份和声誉，换个地址就能洗白作恶
- 纯链上合约无法理解语义化的交付内容（JSON、报告、分析结果）
- 需要"AI 负责验收语义，Web3 负责强制执行"的双层架构

## 最小功能 (MVP)

| # | 功能 | 状态 |
|---|------|------|
| 1 | ERC-8183 状态机：Open → Funded → Submitted → Complete/Reject | ✅ 已实现 |
| 2 | AEP 双模验收引擎：Rule（一票否决）+ LLM（抽样旁轨） | ✅ 已实现 |
| 3 | CAW Pact 资金锁定与 Release 结算 | ✅ 已实现 |
| 4 | AEPReputation.sol 声誉合约（ERC-8004 极简版） | ✅ 已实现 |
| 5 | MUI 前端演示面板（分步引导流） | ✅ 已实现 |
| 6 | 全链路集成测试（Go + httptest） | ✅ 已实现 |

## 技术路径

```
前端 (React + MUI) → Go 后端 (Chi + Zap) → CAW REST API (资金托管)
                                              → PostgreSQL/Redis (任务状态)
                                              → Contract (AEPBounty.sol + AEPReputation.sol)
                                              → IPFS (Pinata) + DeepSeek LLM (验收)
```

## 主要风险

| 风险 | 缓解措施 |
|------|---------|
| LLM 验收幻觉放行次品 | 规则层硬过滤 + LLM 仅做抽样旁轨 |
| CAW TSS 网络延迟 | Custodial 模式保底 + 重试机制 |
| Prompt Injection | Evaluator 地址白名单 + 链下策略预检 |
| 单点开发压力 | 单人参赛，模块明确分工 |

## 验证方式

1. Foundry 合约测试（8/8 全通过）
2. Go 集成测试（10/10 全通过）
3. Demo 三场景：Happy Path / 欺诈拒绝 / 超额度链下拦截
4. 链上交易哈希可查（Base Sepolia）

---

## 进度

- Week 2: 方向选择 + 项目 Proposal
- Week 3: 全链路开发 + 集成测试 + 前端 + 声誉合约
- Week 4: Demo 打磨 + 提交材料准备
