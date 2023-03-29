package chat

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	url = "https://api.openai.com/v1/chat/completions"
)

const (
	user string = "user"
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

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
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

type Response struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Usage   Usage    `json:"usage"`
	Choices []Choice `json:"choices"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Choice struct {
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
	Index        int     `json:"index"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (m Message) String() string {
	return m.Role + " > " + strings.TrimSpace(m.Content)
}
