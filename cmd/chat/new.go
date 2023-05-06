package chat

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/lunarxlark/openai-cli/api"
	"github.com/lunarxlark/openai-cli/cmd/model"
	"github.com/lunarxlark/openai-cli/config"
	"github.com/urfave/cli/v2"
)

const (
	url = "https://api.openai.com/v1/chat/completions"
)

type Role string

const (
	User Role = "user"
)

// Request
type request struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
	TopP        float64   `json:"top_p"`
	User        Role      `json:"user"`
}

type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

// Response
type response struct {
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

func cmdNew(ctx *cli.Context) error {
	var req request
	m := ctx.String("model")
	if m == "" {
		var err error
		m, err = model.List()
		if err != nil {
			return err
		}
	}

	req = request{
		Model:       m,
		Messages:    []Message{},
		Temperature: 0,
		TopP:        0,
		User:        User,
	}

	// ログファイルの作成
	logfile, err := os.Create(filepath.Join(config.OAIConfig.Dir, "chat", fmt.Sprintf("%d.json", time.Now().Unix())))
	if err != nil {
		return err
	}
	defer logfile.Close()

	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s > ", User)
		if !sc.Scan() {
			b, err := json.MarshalIndent(req, "", "  ")
			if err != nil {
				return err
			}
			if _, err := logfile.Write(b); err != nil {
				return err
			}
			break
		}
		req.Messages = append(req.Messages, Message{
			Role:    User,
			Content: sc.Text(),
		})
		var res response
		_ = api.Request(http.MethodPost, url, req, &res)

		for _, choice := range res.Choices {
			req.Messages = append(req.Messages, choice.Message)
			fmt.Printf("%s > %s\n", choice.Message.Role, choice.Message.Content)
		}
	}
	return nil
}
