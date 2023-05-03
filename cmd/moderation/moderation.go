package moderation

import (
	"net/http"

	"github.com/k0kubun/pp"
	"github.com/lunarxlark/openai-cli/api"
	"github.com/urfave/cli/v2"
)

const (
	url = "https://api.openai.com/v1/moderations"
)

type Request struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

type Response struct {
	ID      string   `json:"id"`
	Model   string   `json:"model"`
	Results []Result `json:"results"`
}

type Result struct {
	Categories struct {
		Hate            bool `json:"hate"`
		HateThreatening bool `json:"hate/threatening"`
		SelfHarm        bool `json:"self-harm"`
		Sexual          bool `json:"sexual"`
		SexualMinors    bool `json:"sexual/minors"`
		Violence        bool `json:"violence"`
		ViolenceGraphic bool `json:"violence/graphic"`
	} `json:"category"`
	CategoryScores struct {
		Hate            float64 `json:"hate"`
		HateThreatening float64 `json:"hate/threatening"`
		SelfHarm        float64 `json:"self-harm"`
		Sexual          float64 `json:"sexual"`
		SexualMinors    float64 `json:"sexual/minors"`
		Violence        float64 `json:"violence"`
		ViolenceGraphic float64 `json:"violence/graphic"`
	} `json:"category_scores"`
	Flagged bool `json:"flagged"`
}

func Exec(ctx *cli.Context) error {
	req := Request{
		Input: ctx.String("input"),
		Model: ctx.String("model"),
	}

	var res Response
	if err := api.Request(http.MethodPost, url, req, &res); err != nil {
		return err
	}
	pp.Println(res)
	return nil
}
