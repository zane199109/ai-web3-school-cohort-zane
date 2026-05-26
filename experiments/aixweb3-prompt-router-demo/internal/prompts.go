package internal

import (
	"fmt"
	"strings"
)

func buildRoute(kind, name, reason, q string) RouteResult {
	r := RouteResult{Type: kind, TemplateName: name, RouteReason: reason, RelatedConcepts: Recommend(q)}
	r.Prompt = promptFor(kind, q)
	return r
}

func promptFor(kind, q string) string {
	base := "你是 AI×Web3 概念学习助理。面向普通初学者，用中文回答，避免术语堆砌，不编造事实。用户问题：" + q
	switch kind {
	case "compare":
		return base + "\n请输出：1.核心区别；2.对比表；3.适用场景；4.新手选择建议；5.常见误区；6.推荐继续学习的2-3个概念。"
	case "technical":
		return base + "\n请输出：1.一句话结论；2.简化原理；3.最小流程；4.最小实现例子；5.风险与限制；6.推荐继续学习的2-3个概念。"
	case "learning_path":
		return base + "\n请输出：1.学习顺序；2.每一步目标；3.一个小练习；4.常见卡点；5.推荐继续学习的2-3个概念。"
	case "practice":
		return base + "\n请输出：1.最小可行版本；2.输入输出；3.实现步骤；4.人工确认点；5.验证方式；6.风险与限制。"
	default:
		return base + "\n请输出：1.一句话解释；2.普通人类比；3.具体例子；4.常见误区或边界；5.和AI×Web3的关系；6.推荐继续学习的2-3个概念。"
	}
}

func Recommend(q string) []string {
	text := strings.ToLower(fmt.Sprintf("%s", q))
	if hasAny(text, []string{"agent", "智能体"}) {
		return []string{"Tool Use", "Workflow", "Human-in-the-loop"}
	}
	if hasAny(text, []string{"钱包", "account", "账户"}) {
		return []string{"Smart Account", "Session Key", "Guardrails"}
	}
	if hasAny(text, []string{"prompt", "提示词"}) {
		return []string{"LLM", "Context Window", "Prompt Engineering"}
	}
	return []string{"LLM", "Agent", "Tool Use"}
}
