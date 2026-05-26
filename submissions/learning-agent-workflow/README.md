# AI × Web3 School Learning Agent 工作流配置与运行记录

## 1. 选择的 Agent / AI 工具

本次选择的学习辅助工作流为：

- 主 Agent：Hermes Agent（运行在 WSL / CLI 环境中）
- 辅助能力：
  - GitHub 学习仓库维护
  - Markdown 学习记录生成
  - 每日打卡草稿整理
  - Handbook feedback 记录
  - 会议提醒与学习日程辅助
  - 人工确认后的小任务执行
- 使用原则：Agent 辅助规划、解释、整理、生成草稿和检查结果；不替代本人学习，不自动完成高风险或需人工判断的操作。

本地学习仓库路径：

```text
/home/administrator/ai-web3-school-cohort-zane
```

公开 GitHub 仓库：

```text
https://github.com/zane199109/ai-web3-school-cohort-zane
```

## 2. 让 Agent 帮我完成的学习任务

本次配置的目标是让 Agent 成为 AI × Web3 School 的学习辅助系统，主要任务包括：

1. 阅读并遵循 AI × Web3 School Learning Agent 启动 Prompt。
2. 结合 Handbook，初始化个人学习计划和学习仓库结构。
3. 按每日学习内容生成中文学习日志与打卡草稿。
4. 在学习仓库中维护 submissions、daily、experiments、handbook-feedback 等目录。
5. 将学习过程中发现的问题整理为 Handbook feedback 草稿。
6. 在我明确确认后，辅助创建文件、修改文档、检查 Git 状态。
7. 在我明确确认后，才允许执行 commit、push、WCB 提交等带外部影响的动作。

Agent 的定位不是代写作业，而是帮助我：

- 拆解学习目标；
- 解释 AI / Web3 概念；
- 生成练习和复盘问题；
- 管理公开学习记录；
- 提醒我人工复核关键结果；
- 保持提交内容可审计、可追踪、无敏感信息。

## 3. 关键 Prompt / 配置说明

本次使用的关键启动 Prompt 如下：

```text
请作为我的 AI × Web3 School Learning Agent，先阅读启动 Prompt：https://aiweb3.school/learning-agent.zh.txt ，并结合 Handbook：https://aiweb3.school/zh/handbook/ ，帮我初始化个人学习计划、GitHub 学习仓库、每日打卡草稿和 Handbook feedback 流程。

执行过程中，请特别注意：

Agent 可以帮你规划、解释、整理、生成草稿和检查结果。
涉及创建 repo、写文件、commit、push、WCB 提交等动作，必须先让我人工确认。
涉及钱包签名、转账、授权、合约写入、API Key、token、私钥和助记词的操作，不能由 Agent 自动执行或接触敏感信息。
```

补充配置原则：

- GitHub 仓库为公开仓库，不写入 API Key、token、私钥、助记词、`.env` 文件或任何敏感信息。
- 文件写入可以由 Agent 辅助完成，但涉及 commit / push 前必须展示 `git status` 并等待人工确认。
- WCB 提交、链上操作、钱包签名、授权、转账、合约写入必须由本人在界面中手动完成。
- Agent 可以生成提交草稿，但不能绕过人工复核。
- 每日学习日志采用中文结构化格式，包含学习总览、核心收获、实操踩坑和明日计划。

## 4. 一次成功输出记录

本次 Agent 成功在学习仓库中生成了本任务的提交材料目录，并创建了配置记录 Markdown 文件。

生成路径：

```text
/home/administrator/ai-web3-school-cohort-zane/submissions/learning-agent-workflow/README.md
```

成功输出内容包括：

- 已选择的 Agent / AI 工具；
- Agent 辅助学习任务范围；
- 关键 Prompt 与配置说明；
- 成功运行记录；
- 人工复核与修正记录；
- 安全边界说明；
- 后续可继续执行的学习工作流。

示例成功输出片段：

```text
Agent 的定位不是代写作业，而是帮助我拆解学习目标、解释 AI / Web3 概念、生成练习和复盘问题、管理公开学习记录，并提醒我人工复核关键结果。
```

## 5. 一次人工复核、修正或拒绝 Agent 建议的记录

本次工作流中保留了明确的人工确认边界：

- Agent 可以创建本地 Markdown 文件；
- Agent 可以检查仓库目录和 Git 状态；
- Agent 不会自动 commit / push；
- Agent 不会自动提交 WCB 作业；
- Agent 不会接触或处理钱包私钥、助记词、API Key、token、`.env` 等敏感信息。

一次人工复核记录如下：

```text
人工复核点：在生成本地提交材料后，先检查文件内容和 git status。
处理结果：确认只生成学习记录材料；暂不允许 Agent 自动 commit / push，由本人后续检查后再决定是否提交。
```

这体现了本学习工作流的边界：Agent 负责辅助生成和检查，最终提交与外部动作由本人确认。

## 6. 安全边界与风险控制

为避免 Agent 替代学习或触碰高风险操作，本工作流设置以下限制：

1. 不自动签名钱包交易。
2. 不自动转账或授权。
3. 不自动执行合约写入操作。
4. 不读取、保存或输出私钥、助记词、API Key、token、`.env` 文件。
5. 不自动提交 WCB 作业。
6. 不在未确认情况下创建远程仓库、commit 或 push。
7. 不把未复核的 Agent 输出直接作为最终学习成果提交。

## 7. 后续使用方式

后续我可以继续用该 Agent 工作流完成以下学习辅助任务：

- 每天生成学习计划与打卡草稿；
- 将课程问题整理到 `handbook-feedback/`；
- 为具体概念生成练习题；
- 辅助维护 `daily/` 学习日志；
- 辅助整理 `submissions/` 作业材料；
- 在提交前检查 Markdown 是否覆盖作业要求；
- 在人工确认后再执行 Git 提交相关动作。

## 8. 本次提交检查清单

- [x] 说明选择的 Agent / AI 工具
- [x] 说明 Agent 辅助完成的学习任务
- [x] 提供关键 Prompt / 配置说明
- [x] 提供一次成功输出记录
- [x] 提供一次人工复核、修正或拒绝 Agent 建议的记录
- [x] 不包含 API Key、token、私钥、助记词、`.env` 等敏感信息
- [x] 明确 Agent 只辅助学习，不替代学习
