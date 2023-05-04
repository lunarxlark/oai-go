package completion

import (
	"fmt"
	"net/http"

	"github.com/lunarxlark/openai-cli/api"
	"github.com/urfave/cli/v2"
)

const (
	url string = "https://api.openai.com/v1/completions"
)

type Request struct {
	Model           string         `json:"model"`
	Prompt          string         `json:"prompt"`
	Suffix          *string        `json:"suffix"`
	MaxTokens       int            `json:"max_tokens"`
	Temperature     int            `json:"temperature"`
	TopP            int            `json:"top_p"`
	N               int            `json:"n"`
	Stream          bool           `json:"stream"`
	Logprobs        *int           `json:"logprobs"`
	Echo            bool           `json:"echo"`
	Stop            string         `json:"stop,omitempty"`
	PresencePenalty float64        `json:"presence_penalty"`
	BestOf          int            `json:"best_of"`
	LogitBias       map[string]int `json:"logit_bias,omitempty"`
	User            string         `json:"user"`
}

type Response struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type Choice struct {
	Text         string `json:"text"`
	Index        int    `json:"index"`
	Logprobs     *int   `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}

func Exec(ctx *cli.Context) error {
	req := Request{
		Model:       ctx.String("model"),
		Prompt:      ctx.String("prompt"),
		MaxTokens:   16,
		Temperature: 0,
		TopP:        1,
		N:           ctx.Int("n"),
		Stream:      false,
		Logprobs:    nil,
		Stop:        ctx.String("stop"),
		Echo:        ctx.Bool("echo"),
		BestOf:      1,
		User:        ctx.String("user"),
	}

	var res Response
	if err := api.Request(http.MethodPost, url, req, &res); err != nil {
		return err
	}

	for _, choice := range res.Choices {
		fmt.Println(choice.Text)
	}
	return nil
}
