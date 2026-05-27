# Week 1 Proof-of-Work Pack

> AI × Web3 School 第1周学习成果汇总
>
> 学员：zane199109 | GitHub: [@zane199109](https://github.com/zane199109)
> 学习仓库：https://github.com/zane199109/ai-web3-school-cohort-zane

---

## 一、汇总表格

| # | 任务 | 完成情况 | 链接 |
|---|------|---------|------|
| 1 | AI 概念卡片 | ✅ 涵盖 LLM、Prompt、Context Window、Workflow、Agent、Tool Use、AI Coding、Guardrails、Tracing、Human-in-the-loop | [`submissions/AI Fundamentals Concept.md`](AI%20Fundamentals%20Concept.md) |
| 2 | Web3 概念卡片 | ✅ 涵盖 Address、Wallet、Private Key、Seed Phrase、Signature、Transaction、Gas、Smart Contract、Testnet、Block Explorer | [`submissions/Web3 Fundamentals Concept.md`](Web3%20Fundamentals%20Concept.md) |
| 3 | EOA / 智能账户 / 多签对比 | ✅ 三类账户权限边界与适用场景完整比较 | [`submissions/eoa-smart-account-multisig-comparison.md`](eoa-smart-account-multisig-comparison.md) |
| 4 | 行业信息流搭建 | ✅ 5 大类目 18 个信息源，含内容笔记与课程关联分析 | [`submissions/week1-ai-web3-industry-information-flow.md`](week1-ai-web3-industry-information-flow.md) |
| 5 | Learning Agent 工作流配置 | ✅ Hermes Agent 学习辅助系统，含安全边界与人工确认规则 | [`submissions/learning-agent-workflow/README.md`](learning-agent-workflow/README.md) |
| 6 | AI×Web3 Prompt Router Demo | ✅ Go 语言实现的最小 Prompt 路由学习助手，含前端页面 | [`submissions/aixweb3-prompt-router-demo/README.md`](aixweb3-prompt-router-demo/README.md) — 代码：`experiments/aixweb3-prompt-router-demo/` |
| 7 | 受限 Web3 AI 助手 Workflow | ✅ AI 辅助解析 + 人工最终确权的分层链上操作安全设计 | [`submissions/a Restricted Web3 Assistant.md`](a%20Restricted%20Web3%20Assistant.md) |
| 8 | AI×Web3 协议分析（ERC-8004 & ERC-8183） | ✅ 双协议联动拆解：Agent 信任层 + 商业执行层 | [`submissions/Analysis AIXWeb3 projects.md`](Analysis%20AIXWeb3%20projects.md) |
| 9 | Sepolia 测试网基础转账 | ✅ 0.01 ETH 转账成功，区块浏览器全参数验证 | [`submissions/sepolia-basic-transaction.md`](sepolia-basic-transaction.md) — [Etherscan](https://sepolia.etherscan.io/tx/0xe6b44ff57e527cb7d03973b7d31c621653089056c88699abbc1e166b0c674bc2) |
| 10 | Sepolia 合约部署与交互 | ✅ SimpleStorage 合约部署 + setNumber(1000) 写入 + 读取验证 | [`submissions/deploy-operation-smartcontract/deploy-operation-smartcontrat.md`](deploy-operation-smartcontract/deploy-operation-smartcontrat.md) — [Etherscan](https://sepolia.etherscan.io/address/0x358E0d48d079DE87940E243242D2F60886a215E3) |
| 11 | AI×Web3 最小链上支付工作流 | ✅ 8 步闭环流程图：用户发起 → AI 解析 → 人工核验 → 钱包签名 → 链上确认 → AI 校验 | [`submissions/AIXWEB3-workflow/Draw a minimal AIXWEB3 workflow.md`](AIXWEB3-workflow/Draw%20a%20minimal%20AIXWEB3%20workflow.md) |
| 12 | 问题与人工修正记录 | ✅ 2 条：Prompt Router 方向修正、Learning Agent commit/push 安全边界确认 | 见下方第六节 |

---

## 二、AI 学习记录 / 概念卡片

### 2.1 AI 核心概念概念卡片（10 张）

围绕 10 个 AI 基础概念，结合课程学习、Agent 使用经验和实际开发场景，逐一改写释义、用例与常见误区：

LLM、Prompt、Context Window、Workflow、Agent、Tool Use、AI Coding、Guardrails、Tracing、Human-in-the-loop

📄 **提交文件**：`submissions/AI Fundamentals Concept.md`

### 2.2 Web3 核心概念概念卡片（10 张）

围绕 10 个 Web3 基础概念，结合课程实验和实操经验逐一改写释义、用例、安全提醒与常见误区：

Address、Wallet、Private Key、Seed Phrase、Signature、Transaction、Gas、Smart Contract、Testnet、Block Explorer

📄 **提交文件**：`submissions/Web3 Fundamentals Concept.md`

### 2.3 行业信息流搭建

建立了覆盖 5 大核心类目共 18 个信息源的标准化行业观察体系，并精选 5 篇高质量内容笔记（含 Cobo Agent 钱包安全、ERC-8004 Agent 标准、Vitalik 隐私方案、GoPlus 安全风控、Phala TEE 可信执行）。

📄 **提交文件**：`submissions/week1-ai-web3-industry-information-flow.md`

### 2.4 EOA / 智能账户 / 多签账户权限比较

深度比较三类账户在控制权、交易发起、批准机制、恢复、限额、自动化、Gas 支付和安全风险上的差异，明确 AI Agent 场景下账户选择的工程原则。

📄 **提交文件**：`submissions/eoa-smart-account-multisig-comparison.md`

---

## 三、Learning Agent / AI 工具实践记录

### 3.1 Learning Agent 工作流配置

在 WSL 环境中配置 Hermes Agent 作为个人学习辅助系统：

- **主 Agent**：Hermes Agent（CLI 模式运行在 WSL）
- **辅助能力**：GitHub 仓库维护、Markdown 学习记录生成、每日打卡草稿、Handbook feedback 记录
- **安全边界**：Agent 辅助规划/解释/整理/校验；commit/push/WCB 提交需人工确认；钱包操作完全隔离

📄 **提交文件**：`submissions/learning-agent-workflow/README.md`
📄 **补充证明**：`submissions/learning-agent-workflow/evidence.md`

### 3.2 AI×Web3 Prompt Router Demo

用 Go 语言实现的一个最小可运行 Prompt Router 学习助手：

- 用户输入学习问题 → 关键词启发式判断问题类型 → 匹配 Prompt 模板 → 调用 Hermes 模型接口 → 展示路由原因、AI 回答与推荐后续概念
- 覆盖概念解释、流程说明、对比分析、代码示例等多类型问题路由

📄 **提交文档**：`submissions/aixweb3-prompt-router-demo/README.md`
💻 **代码目录**：`experiments/aixweb3-prompt-router-demo/`

### 3.3 受限 Web3 AI 助手 Workflow 设计

设计了一套「AI 智能辅助 + 人工最终确权」的分层 Workflow，AI 全权负责解析/规划/风控/草稿生成/结果校验，人工全权负责签名/授权/转账/合约写入等高危操作。

📄 **提交文件**：`submissions/a Restricted Web3 Assistant.md`

### 3.4 AI×Web3 协议分析

拆解 ERC-8004（Agent 信任层：身份+信誉+校验三层注册表）与 ERC-8183（Agent 商业层：任务+托管+结算商业执行链路）双协议联动逻辑，内含 Mermaid 流程图。

📄 **提交文件**：`submissions/Analysis AIXWeb3 projects.md`

---

## 四、测试网交易记录 & 链上验证

### 4.1 Sepolia 测试网基础转账

在 Sepolia 测试网完成 0.01 ETH 转账交易，并通过 Etherscan 完整验证交易状态、Gas 消耗、区块高度等参数。

| 项目 | 内容 |
|------|------|
| 交易哈希 | `0xe6b44ff57e527cb7d03973b7d31c621653089056c88699abbc1e166b0c674bc2` |
| 区块浏览器链接 | https://sepolia.etherscan.io/tx/0xe6b44ff57e527cb7d03973b7d31c621653089056c88699abbc1e166b0c674bc2 |
| 交易状态 | Success（区块 10912519） |
| 发送方地址 | `0xcc6142f3f79Dd1d42FC0446C5B7218C5F520021E` |

📄 **提交文件**：`submissions/sepolia-basic-transaction.md`

### 4.2 Sepolia 测试网合约部署与交互

部署 `SimpleStorage.sol` 最小状态存储合约，完成从部署 → 写入(setNumber=1000) → 读取 → 区块浏览器验证的完整链路。

| 项目 | 内容 |
|------|------|
| 合约地址 | `0x358E0d48d079DE87940E243242D2F60886a215E3` |
| Etherscan 链接 | https://sepolia.etherscan.io/address/0x358E0d48d079DE87940E243242D2F60886a215E3 |
| 写入结果 | `setNumber(1000)` 成功上链 |
| 读取验证 | `number` 返回 `1000` |

📄 **提交文件**：`submissions/deploy-operation-smartcontract/deploy-operation-smartcontrat.md`
🖼️ **截图证据**：`submissions/deploy-operation-smartcontract/`（deploy.png, write.png, read.png, verify.png, etherscan.png, environment.png）

---

## 五、AI × Web3 最小交叉实验 / 流程图

### 5.1 最小 AI×Web3 链上支付工作流

设计 8 步链上支付工作流，涵盖用户发起 → AI 解析 → 工具查询 → 人工核验 → 钱包签名 → 链上确认 → AI 校验 → 报告输出的完整闭环。核心原则：**AI 辅助运算，人工掌控核心**。

![链上支付工作流](AIXWEB3-workflow/agent链上支付.jpg)

📄 **提交文件**：`submissions/AIXWEB3-workflow/Draw a minimal AIXWEB3 workflow.md`

### 5.2 ERC-8004 + ERC-8183 联动闭环图

包含三层 Mermaid 图：
1. ERC-8004 三层信任架构流转图（身份→校验→信誉）
2. ERC-8183 任务状态机流转图（创建→托管→交付→验收→结算）
3. 8004+8183 端到端联动时序图（事前信誉筛选→事中托管风控→事后信用迭代）

📄 **提交文件**：`submissions/Analysis AIXWeb3 projects.md`（内嵌 Mermaid 图）

---

## 六、问题与人工修正记录

### 问题 1：第一版 Prompt Router Demo 设计不符合需求

**卡点描述**：Agent 生成的 Prompt Router Demo 第一版与我构思的产品方向不一致，直接生成不是我想要的效果。

**人工修正过程**：
1. 要求 Agent 先向我说明设计方案，经我确认后才开始编码
2. 沟通后重新定位为「学习辅助」场景，而非纯技术演示
3. Agent 输出设计方案 → 我确认 → 执行代码 → 本地运行验证

📄 **相关文件**：`submissions/aixweb3-prompt-router-demo/README.md`（第六节"AI 辅助与人工验证说明"）

### 问题 2：Learning Agent 工作流的安全边界确认

**卡点描述**：Agent 尝试自动执行 Git 提交，但根据安全策略，Agent 不应自动 commit/push。

**人工修正过程**：
1. 明确要求 Agent 在 commit/push 前展示 `git status` 并等待人工确认
2. 在 Learning Agent 配置中固化此安全边界规则
3. 最终形成"Agent 创建本地文件 → 展示状态 → 人工确认后提交"的协作流程

📄 **相关文件**：`submissions/learning-agent-workflow/README.md`（第5节"一次人工复核、修正或拒绝 Agent 建议的记录"）

---

## 七、提交验证

- [x] 汇总入口为公开 GitHub 仓库 README
- [x] 包含 AI 概念卡片（10 张）
- [x] 包含 Web3 概念卡片（10 张）
- [x] 包含 Learning Agent / AI 工具实践记录
- [x] 包含测试网交易记录（tx hash + 区块浏览器链接 + 合约地址）
- [x] 包含 AI × Web3 流程图 / 工作流设计
- [x] 包含问题与人工修正记录（2 条）
- [x] 不包含私钥、助记词、API Key、token、.env 等敏感信息
