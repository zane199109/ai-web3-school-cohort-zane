package internal

import "testing"

func TestRouteQuestionCompare(t *testing.T) {
	got := RouteQuestion("Smart Account 和普通钱包有什么区别？")
	if got.Type != "compare" {
		t.Fatalf("Type = %q, want compare", got.Type)
	}
	if got.TemplateName != "Concept Comparison" {
		t.Fatalf("TemplateName = %q, want Concept Comparison", got.TemplateName)
	}
	if got.Prompt == "" {
		t.Fatal("Prompt should not be empty")
	}
}

func TestRouteQuestionTechnical(t *testing.T) {
	got := RouteQuestion("Tool Use 是怎么实现的？")
	if got.Type != "technical" {
		t.Fatalf("Type = %q, want technical", got.Type)
	}
}

func TestRouteQuestionDefaultConcept(t *testing.T) {
	got := RouteQuestion("Agent 是什么意思？")
	if got.Type != "concept" {
		t.Fatalf("Type = %q, want concept", got.Type)
	}
	if len(got.RelatedConcepts) == 0 {
		t.Fatal("RelatedConcepts should not be empty")
	}
}

func TestRecommendAgent(t *testing.T) {
	got := Recommend("Agent")
	want := "Tool Use"
	if len(got) == 0 || got[0] != want {
		t.Fatalf("Recommend first = %v, want %q", got, want)
	}
}
