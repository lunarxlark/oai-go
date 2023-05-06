package file

import (
	"fmt"
	"net/http"

	"github.com/lunarxlark/openai-cli/api"
	"github.com/urfave/cli/v2"
)

type ResponseFileDelete struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

func cmdDelete(ctx *cli.Context) error {

	url := fmt.Sprintf("%s/%s", url, ctx.String("file_id"))

	var res ResponseFileDelete
	if err := api.Request(http.MethodDelete, url, nil, &res); err != nil {
		return err
	}

	if res.Deleted {
		fmt.Println("success file delete")
		fmt.Printf("ID:%s", res.ID)
	}

	return nil
}
