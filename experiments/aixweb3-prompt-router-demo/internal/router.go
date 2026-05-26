package internal

import "strings"

type RouteResult struct {
	Type            string   `json:"type"`
	TemplateName    string   `json:"templateName"`
	RouteReason     string   `json:"routeReason"`
	RelatedConcepts []string `json:"relatedConcepts"`
	Prompt          string   `json:"-"`
}

func RouteQuestion(q string) RouteResult {
	s := strings.ToLower(q)
	if hasAny(s, []string{"区别", "对比", "vs", "不同", "差异"}) {
		return buildRoute("compare", "Concept Comparison", "问题包含对比类关键词，适合用表格辨析差异。", q)
	}
	if hasAny(s, []string{"原理", "怎么实现", "机制", "为什么", "底层"}) {
		return buildRoute("technical", "Technical Explainer", "问题关注实现机制，适合加入原理和最小例子。", q)
	}
	if hasAny(s, []string{"怎么学", "入门", "学习路线", "学习路径", "先学什么"}) {
		return buildRoute("learning_path", "Learning Path", "问题关注学习顺序，适合输出路径和练习。", q)
	}
	if hasAny(s, []string{"做一个", "实现", "demo", "workflow", "项目", "助手"}) {
		return buildRoute("practice", "Practice Task Planner", "问题关注产物实现，适合拆成最小任务。", q)
	}
	return buildRoute("concept", "Beginner Concept Explainer", "问题适合先建立概念理解。", q)
}

func hasAny(s string, words []string) bool {
	for _, w := range words {
		if strings.Contains(s, strings.ToLower(w)) {
			return true
		}
	}
	return false
}
