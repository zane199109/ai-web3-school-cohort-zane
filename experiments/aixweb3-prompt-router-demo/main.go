package main

import (
	"encoding/json"
	"log"
	"net/http"

	app "aixweb3-prompt-router-demo/internal"
)

type askReq struct {
	Question string `json:"question"`
}
type askResp struct {
	Question        string   `json:"question"`
	Type            string   `json:"type"`
	TemplateName    string   `json:"templateName"`
	RouteReason     string   `json:"routeReason"`
	Answer          string   `json:"answer"`
	RelatedConcepts []string `json:"relatedConcepts"`
	Error           string   `json:"error,omitempty"`
}

func main() {
	cfg := app.LoadEnv(".env")
	client := app.AIClient{Config: cfg}
	http.Handle("/", http.FileServer(http.Dir("web")))
	http.HandleFunc("/api/ask", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", 405)
			return
		}
		var in askReq
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Question == "" {
			http.Error(w, "invalid question", 400)
			return
		}
		route := app.RouteQuestion(in.Question)
		answer, err := client.Ask(route.Prompt)
		resp := askResp{Question: in.Question, Type: route.Type, TemplateName: route.TemplateName, RouteReason: route.RouteReason, Answer: answer, RelatedConcepts: route.RelatedConcepts}
		if err != nil {
			resp.Error = err.Error()
			resp.Answer = "模型接口暂不可用：请检查本地 .env 配置。"
		}
		json.NewEncoder(w).Encode(resp)
	})
	log.Println("server started at http://localhost:" + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
