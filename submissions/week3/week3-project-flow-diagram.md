# 项目流程图 — AEP 最小闭环

> **项目**: AEP (Agent Escrow Protocol)
> **赛道**: Cobo — Agentic Economy × Agentic Wallet

---

## 架构总览

```
┌─────────────┐     ┌──────────────┐     ┌──────────────┐
│  Buyer      │     │  Seller      │     │  Evaluator   │
│  (CAW钱包)  │     │  (CAW钱包)   │     │  (CAW钱包)   │
└──────┬──────┘     └──────┬───────┘     └──────┬───────┘
       │                    │                     │
       ▼                    ▼                     ▼
┌──────────────────────────────────────────────────────┐
│                AEP Go Backend (Chi + Zap)              │
│  ┌────────┐  ┌──────────┐  ┌────────┐  ┌───────────┐ │
│  │ Handler│→ │ Provider │→ │ Engine │→ │ Relayer   │ │
│  │ (API)  │  │ Layer    │  │ (验收)  │  │ (结算)    │ │
│  └────────┘  └──────────┘  └────────┘  └───────────┘ │
└──────────────────────────────────────────────────────┘
       │           │           │           │
       ▼           ▼           ▼           ▼
┌───────────┐ ┌────────┐ ┌──────────┐ ┌───────────┐
│ CAW Pact │ │REST API│ │  IPFS    │ │ DeepSeek  │
│ 资金锁定 │ │CAW结算 │ │ 文件存储  │ │ LLM验收   │
└───────────┘ └────────┘ └──────────┘ └───────────┘
       │
       ▼
┌────────────────────────────────────────────┐
│          链上合约层 (Anvil / Base Sepolia)  │
│  ┌────────────────┐  ┌──────────────────┐  │
│  │  AEPBounty.sol  │  │ AEPReputation.sol│  │
│  │  ERC-8183 状态机│  │ ERC-8004 极简版   │  │
│  └────────────────┘  └──────────────────┘  │
└────────────────────────────────────────────┘
```

## 用户输入 → AI Agent 处理 → Web3 机制 → 输出

### 流程分解

```
1. 用户输入 (Buyer)
   Buyer 通过 MUI 前端填写：
   - 交付物要求 (自然语言描述)
   - 赏金金额 (USDC)
   - 最低声誉要求 (min_reputation)
       │
       ▼
2. AI Agent 处理 (Go Backend)
   a. API Handler 解析请求
   b. Provider Layer: CAW Pact 创建, 资金锁定
   c. Engine Layer: 新任务入队, 等待 Seller 抢单
       │
       ▼
3. Web3 机制 (链上)
   a. AEPBounty.sol: Open → Funded (资金锁定)
   b. Seller 检查自身声誉 → 如果 >= min_reputation → 抢单
   c. Redis + PG 双锁防并发
       │
       ▼
4. AI Agent 处理 (Seller 提交)
   a. Seller 提交交付物 (人工或自动)
   b. IPFS Pinata 上传 → 获取 CID
       │
       ▼
5. 验证 (AI 双模验收引擎)
   a. Rule Engine (硬校验): 
      - 非空字段
      - 必填结构完整性
      - 一票否决权
   b. LLM 旁轨 (DeepSeek):
      - 相关性检测
      - 幻觉/空数据检测
      - 超时 10s 降级到规则引擎
       │
       ▼
6. Web3 机制 (结算)
   a. 验收通过 → CAW Release 放款给 Seller
   b. 验收失败 → CAW Slash 罚没 + 退款给 Buyer
   c. AEPReputation.sol: +10 (成功) / -20 (欺诈)
       │
       ▼
7. 输出结果
   - Buyer: 收到交付物 (质量合格)
   - Seller: 收到报酬 + 声誉 +1
   - 链上: 交易哈希 + 事件日志
   - 前端: 拓扑图实时更新
```

## 验证材料

| 结果 | 验证方式 | 状态 |
|------|---------|------|
| 合约测试 8/8 | `forge test` | ✅ |
| 集成测试 10/10 | `go test -run TestDemo_` | ✅ |
| CAW Pact 记录 | CAW Dashboard | ✅ |
| 链上交易哈希 | Anvil / Base Sepolia Scan | ✅ |
| Demo GIF | 浏览器录制 | 📝 Week 4 |
| 截图集 | 浏览器截图 | 📝 Week 4 |
