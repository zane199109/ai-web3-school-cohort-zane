# 深度研究包 — 围绕主方向的 3 个核心标准/协议/SDK

> **主方向**: Payment / Commerce / Settlement — Agent 自主商业的可信结算
> **学员**: zane

---

## 1. ERC-8183 — Agentic Commerce Protocol

### 它解决什么问题

ERC-8183 定义了 AI Agent 间商业交易的标准状态机：**Open → Funded → Submitted → Complete / Reject**。每个状态转换有明确的触发条件和可验证数据（如 jobIntent、proofURI、evaluationResult）。

它为 Agent 经济提供了一个**中立、可编程的结算原语**——任何 Agent 只要实现了这个接口，就可以直接与其他 Agent 进行标准化的商业交互，无需自定义合约。

### 边界

- 标准化了"怎么做"（状态机），但**不定义"怎么验证"**— 验收逻辑留给开发者实现
- 假设 Evaluator 是可信的，没有内置争议仲裁机制
- 定义的是单 Job 生命周期，没有声誉/身份/权限层面

### 还缺什么

- 缺少 LLM 验收的标准接口（怎么传评估参数、怎么处理不确定结果）
- 缺少跨协议声誉互通（ERC-8183 的 Job 记录和 ERC-8004 的声誉需要手动同步）
- 缺少紧急熔断/退款机制（合约漏洞时资金可能锁死）

### AEP 的改进

AEP 以 ERC-8183 为状态机基线，增加了：
- **AEPReputation.sol**：独立声誉层（ERC-8004 极简版）
- **双模验收引擎**：Rule + LLM，不依赖单点 Evaluator
- **紧急退款**：`claimRefund` 不受 Hook 阻止

---

## 2. CAW (Cobo Agentic Wallet) + Pact 策略引擎

### 它解决什么问题

CAW 为 AI Agent 提供了**可控边界内的链上资金管理能力**。核心组件：

- **MPC 2-of-2 签名**：Agent 持 1 个分片，Cobo 云持 1 个分片——Agent 提议交易，人类在 CAW App 上审批（第二签名），双重保障
- **Pact 策略引擎**：链下预检白名单合约地址、单笔限额、日预算上限，不符合条件在链下拦截（0 Gas）
- **Recipe 知识胶囊**：预定义安全模板，Agent 只需填入参数即可发起合规交易

### 边界

- Custodial 模式无 MPC 签名，适用于快速验证场景，但缺少"人类审批"的安全环节
- 策略引擎在链下运行，理论上可被绕过（如果攻击者直接操作 CAW API）
- 不支持批量交易原子性（多个 Pact 需要分别审批）

### 还缺什么

- Go 语言 SDK 缺失（目前只有 Python/TypeScript）
- MPC TSS Node 网络稳定性对开发环境依赖大（WSL 环境下表现不佳）
- 本地调试体验不够友好（TSS 初始化需要 16+ 位密码）

### AEP 的使用

AEP 使用 CAW 作为资金托管层：Buyer 创建 Pact 锁定资金 → 验收通过后 Release → 失败则自动退还。当前使用 Custodial 模式跑通全链路，后续可切到 MPC 模式展示完整安全叙事。

---

## 3. ERC-8004 — Agent Identity & Asset (极简版)

### 它解决什么问题

ERC-8004 定义了 AI Agent 的链上身份三层注册表：
- **Identity（身份）**：Agent 的唯一链上标识
- **Reputation（声誉）**：可验证的链上行为记录（Pass/Fail 次数、权重积分）
- **Verification（验证）**：外部校验机制（EAS Attestation、ZK 证明等）

它为 Agent 提供了"不能换个地址就洗白"的链上声誉锚定。

### 边界

- 仅定义了**读写接口**（getScore, updateScore），没有内置计算逻辑
- 声誉权重的设计规则（+10/-20 是否合理）完全留给应用层
- 没有跨协议声誉聚合（AEP 的声誉记录和 FluxA 的声誉记录不互通）

### 还缺什么

- 防刷机制（Sybil 攻击）需要应用层实现
- 声誉可移植性（如何让不同协议间的声誉互相认可）
- 分数的时间衰减（100 天前的 1 次成功和 1 天前的 1 次失败，哪个更可信？）

### AEP 的实现

AEP 实现了 ERC-8004 极简版——`AEPReputation.sol`：
- `onlyBounty` 权限：只有 AEPBounty 合约可以写入声誉
- `+10` 成功 / `-20` 欺诈（不对称权重，惩罚比奖励更重）
- `getScore(address)` → `eth_call` 查询（0 Gas，Go 后端用作抢单准入检查）
