package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/lunarxlark/openai-cli/config"
)

const (
	chatURL = "https://api.openai.com/v1/chat/completions"
)

type Role string

const (
	User Role = "user"
)

type Request struct {
	Model       string    `json:"model" binding:"required"`
	Messages    []Message `json:"messages" binding:"required"`
	Temperature float64   `json:"temperature"`
	TopP        float64   `json:"top_p"`
	User        string    `json:"user"`
}

func CreateReq(model string, messages []Message) *Request {
	return &Request{
		Model:    model,
		Messages: messages,
	}
}

func (r *Request) Request() (*Response, error) {
	payload, err := json.Marshal(&r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, chatURL, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+config.OAIConfig.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resbody := new(Response)
	if err := json.Unmarshal(body, resbody); err != nil {
		return nil, err
	}

	return resbody, nil
}
