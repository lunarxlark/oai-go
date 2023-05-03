package edit

import (
	"fmt"
	"net/http"

	"github.com/lunarxlark/openai-cli/api"
	"github.com/urfave/cli/v2"
)

const (
	url string = "https://api.openai.com/v1/edits"
)

type Request struct {
	Model       string  `json:"model"`
	Input       string  `json:"input"`
	Instruction string  `json:"instruction"`
	N           int     `json:"n"`
	Temperature float64 `json:"temperature"`
	TopP        float64 `json:"top_p"`
}

type Response struct {
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	Text  string `json:"text"`
	Index int    `json:"index"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func Exec(ctx *cli.Context) error {
	payload := Request{
		Model:       ctx.String("model"),
		Input:       ctx.String("input"),
		Instruction: ctx.String("instruction"),
		N:           ctx.Int("n"),
		Temperature: ctx.Float64("temperature"),
		TopP:        ctx.Float64("top_p"),
	}

	var res Response
	if err := api.Request(http.MethodPost, url, payload, &res); err != nil {
		return err
	}

	for _, choice := range res.Choices {
		fmt.Println(choice.Text)
	}

	return nil
}
