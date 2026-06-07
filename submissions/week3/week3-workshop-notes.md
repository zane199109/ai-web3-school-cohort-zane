# Workshop 笔记 — Cobo Agentic Wallet + 黑客松赛道实战

> **学员**: zane
> **项目**: AEP (Agent Escrow Protocol)

---

## Workshop 1: Cobo Agentic Wallet — 产品经理分享 (5/26)

### Sponsor 解决什么问题

Cobo Agentic Wallet 解决的是 AI Agent **"如何安全地持有和支配链上资金"** 的问题。传统的私钥管理模式不适合 Agent——Agent 可能被 Prompt Injection 欺骗签名、可能超出预算、可能调用了恶意合约。

CAW 的核心理念：**Agent 可以提议交易，但人类必须批准**。

### 提供什么工具

| 工具 | 用途 | AEP 使用方式 |
|------|------|-------------|
| MPC 2-of-2 签名 | 双人审批签名 | Agents 提议, Buyer 在 CAW App 审批 |
| Pact 策略引擎 | 链下权限预检 | Buyer 创建 Pact 锁定资金时检查预算 |
| Recipe 知识胶囊 | 安全模板 | 预定义 Bounty 发放模板 |
| REST API | 程序化操作 | Go 后端直接调用 |
| TSS Node | MPC 密钥管理 | 本地签名守护进程 |

### 适合哪个赛道

**Cobo — Agentic Economy × Agentic Wallet**。AEP 的整个资金托管—裁决—结算流程完全依赖 CAW 的 Pact 机制和 MPC 签名架构。

### 可以做什么 Demo

- Buyer 创建带 Pact 的 Bounty（展示 CAW 审批环节）
- Seller 提交交付物后 AEP 验收
- Buyer 在 CAW App 上确认放款
- 链上声誉更新

---

## Workshop 2: 黑客松赛道实战 (6/3)

### Sponsor 解决什么问题

Cobo 团队介绍了 CAW 的三种接入方式：

| 方式 | 适用场景 | 复杂度 |
|------|---------|--------|
| **Skills** | 结合 Agent 框架使用，最简单 | 低 |
| **CLI** | 快速实验和手动操作 | 中 |
| **REST API** | 最底层接口，灵活性和自主性最强 | 高 |

### 关键收获

1. **AEP 选择 REST API 是正确的** — 因为 Go 后端需要完全控制资金操作流程
2. **Policy 可以通过 Agent 自动编写** — Intent 识别后 Agent 生成策略 → 用户在 App 审批
3. **Skills CLI 适合快速原型验证** — 后续 Hackathon 演示可准备 Skills CLI 作为备选方案

### 对 AEP 的启示

- Custodial 模式已跑通全链路，MPC 模式作为安全叙事展示
- CAW 的 Rate Limit 和并发策略需要进一步了解（已列入 Sponsor/Mentor 问题清单）
- Demo 中需要重点展示 CAW App 审批环节以增强可信叙事
