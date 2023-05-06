package file

import (
	"net/http"

	"github.com/k0kubun/pp"
	"github.com/lunarxlark/openai-cli/api"
	"github.com/urfave/cli/v2"
)

const (
	url string = "https://api.openai.com/v1/files"
)

type Response struct {
	Data   []Data `json:"data"`
	Object string `json:"object"`
}

type Data struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int    `json:"created_at"`
	FileName  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

func cmdList(ctx *cli.Context) error {
	if ctx.String("file_id") != "" {
		return Get(ctx)
	} else {
		return List(ctx)
	}
}

func List(ctx *cli.Context) error {
	var res Response
	if err := api.Request(http.MethodGet, url, nil, &res); err != nil {
		return err
	}

	// TODO:fix output
	pp.Println(res)
	return nil
}

func Get(ctx *cli.Context) error {
	var res Data
	if err := api.Request(http.MethodGet, url+"/"+ctx.String("file_id"), nil, &res); err != nil {
		return err
	}

	// TODO:fix output
	pp.Println(res)
	return nil
}
