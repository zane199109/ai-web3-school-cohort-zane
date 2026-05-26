package internal

import (
	"bufio"
	"os"
	"strings"
)

type Config struct{ APIURL, APIKey, Model, Port string }

func LoadEnv(path string) Config {
	cfg := Config{Port: "8080"}
	file, err := os.Open(path)
	if err == nil {
		defer file.Close()
		s := bufio.NewScanner(file)
		for s.Scan() {
			line := strings.TrimSpace(s.Text())
			if line == "" || strings.HasPrefix(line, "#") || !strings.Contains(line, "=") {
				continue
			}
			parts := strings.SplitN(line, "=", 2)
			os.Setenv(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		}
	}
	cfg.APIURL = os.Getenv("HERMES_MODEL_API_URL")
	cfg.APIKey = os.Getenv("HERMES_MODEL_API_KEY")
	cfg.Model = os.Getenv("HERMES_MODEL_NAME")
	if p := os.Getenv("PORT"); p != "" {
		cfg.Port = p
	}
	return cfg
}
