package edit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lunarxlark/openai-cli/config"
	"github.com/urfave/cli/v2"
)

type Request struct {
	Model       string  `json:"model"`
	Input       string  `json:"input"`
	Instruction string  `json:"instruction"`
	N           int     `json:"n"`
	Temperature float64 `json:"temperature"`
	TopP        int     `json:"top_p"`
}

const (
	url string = "https://api.openai.com/v1/edits"
)

func Exec(ctx *cli.Context) error {
	client := http.Client{}

	payload, err := json.Marshal(Request{
		Model:       "text-davinci-edit-001",
		Input:       "hello wrold!!",
		Instruction: "fix the spelling mistakes.",
		N:           1,
		Temperature: 0,
		TopP:        0,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
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
