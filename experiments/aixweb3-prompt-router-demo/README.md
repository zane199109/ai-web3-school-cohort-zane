# AI×Web3 Prompt Router Demo

一个使用 Go 后端 + Hermes 模型接口的最小联网学习 Demo。用户只需要输入简单问题，系统会自动识别问题类型，选择对应 Prompt 模板，再调用 AI 生成更适合初学者理解的答案。

## 安全说明

- `.env` 只用于本地运行，不得提交到 GitHub。
- 仓库只保留 `.env.example`。
- 不在代码、README、截图、提交文档中暴露真实 API Key。

## 本地运行

```bash
cp .env.example .env
# 编辑 .env，填入本地 Hermes 模型接口配置
go run .
```

浏览器打开：

```text
http://localhost:8080
```
