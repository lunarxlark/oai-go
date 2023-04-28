package file

import (
	"fmt"
	"io"
	"net/http"

	"github.com/lunarxlark/openai-cli/config"
	"github.com/urfave/cli/v2"
)

func List(ctx *cli.Context) error {
	client := http.Client{}

	url := "https://api.openai.com/v1/files"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+config.OAIConfig.APIKey)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println((string(body)))

	return nil
}
