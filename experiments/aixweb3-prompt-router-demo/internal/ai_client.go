package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type AIClient struct{ Config Config }

type chatReq struct {
	Model       string              `json:"model"`
	Messages    []map[string]string `json:"messages"`
	Temperature float64             `json:"temperature"`
}

type chatResp struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (c AIClient) Ask(prompt string) (string, error) {
	if c.Config.APIURL == "" || c.Config.APIKey == "" || c.Config.Model == "" {
		return "", errors.New("missing HERMES_MODEL_API_URL / HERMES_MODEL_API_KEY / HERMES_MODEL_NAME")
	}
	body, _ := json.Marshal(chatReq{Model: c.Config.Model, Temperature: 0.4, Messages: []map[string]string{
		{"role": "system", "content": "你是严谨的AI×Web3学习助理，回答必须结构清晰、适合初学者。"},
		{"role": "user", "content": prompt},
	}})
	req, _ := http.NewRequest("POST", c.Config.APIURL, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Config.APIKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	var out chatResp
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return "", err
	}
	if len(out.Choices) == 0 {
		return "", errors.New("empty model response")
	}
	return out.Choices[0].Message.Content, nil
}
