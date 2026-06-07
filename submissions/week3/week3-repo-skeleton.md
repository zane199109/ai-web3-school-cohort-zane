# Repo Skeleton — AEP (Agent Escrow Protocol)

> **GitHub**: https://github.com/zane/AEP-Hackathon
> **提交时间**: 2026-06-07

---

## 项目 Repo 结构

```
AEP-Hackathon/
├── backend-go/             # Go 后端
│   ├── api/               # HTTP 路由 + SSE
│   ├── cmd/               # main.go + demo_test.go
│   ├── config/            # 配置加载 (Viper)
│   ├── store/             # PostgreSQL + Redis 持久化
│   ├── providers/         # CAW / IPFS / LLM 外部服务封装
│   ├── engine/            # 规则引擎 + LLM 验收
│   └── relayer/           # CAW 结算重试
├── contract-foundry/       # Solidity 合约
│   ├── src/               # AEPBounty.sol + AEPReputation.sol
│   └── script/            # 部署脚本
├── frontend-web/           # React + MUI 前端
│   └── src/               # App.jsx, components
├── conf/                   # 配置文件
├── docs/                   # PRD, Design, Demo
├── scripts/                # DB schema, CLI 工具
├── verification_reports/   # 验证清单
├── README.md               # 本文件
├── CONTEXT.yaml            # 项目状态快照
├── Makefile                # Build 快捷命令
└── docker-compose.yml      # PostgreSQL + Redis
```

## README 包含内容

| 要求 | 状态 | 说明 |
|------|------|------|
| Problem | ✅ | AI Agent 间缺乏可信中立验收机制 |
| Track | ✅ | Cobo — Agentic Economy × Agentic Wallet |
| MVP Flow | ✅ | 5-step 引导流（发榜→抢单→验收→审批→结算） |
| Tech Stack | ✅ | Go + Solidity + React + CAW + DeepSeek |
| Risks | ✅ | 6 条关键红线（BuyerApproval/Redis Lock/LLM 降级等） |
| Validation Plan | ✅ | Foundry 8/8 + Go 10/10 + 三场景 Demo |

## 快速启动

```bash
# 1. 启动基础设施
docker compose up -d

# 2. 启动后端
cd backend-go && go run ./cmd/main.go -config ../conf/config.yaml

# 3. 启动前端
cd frontend-web && npm run dev

# 4. 运行集成测试
cd backend-go && go test ./cmd/ -run TestDemo_ -v -count=1
```
