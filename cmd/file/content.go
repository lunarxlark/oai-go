package file

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/lunarxlark/openai-cli/config"
	"github.com/urfave/cli/v2"
)

func Content(ctx *cli.Context) error {

	url := fmt.Sprintf("%s/%s/content", url, ctx.String("file_id"))

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.OAIConfig.APIKey)

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)

	return nil
}
