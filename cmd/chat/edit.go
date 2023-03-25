package chat

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/lunarxlark/openai-go/cmd/model"
	"github.com/urfave/cli/v2"
)

func CmdContinue(ctx *cli.Context) error {
	model, err := model.List()
	if err != nil {
		return err
	}

	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s      > ", user)
		sc.Scan()
		content := sc.Text()

		reqbody := ReqBody{
			Model: model,
			Messages: []Message{
				{Role: "user", Content: content},
			},
		}

		payload, err := json.Marshal(&reqbody)
		if err != nil {
			return err
		}

		req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
		req.Header.Set("Content-Type", "application/json")

		client := new(http.Client)
		res, err := client.Do(req)
		if err != nil {
			return err
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		resbody := new(ResBody)
		if err := json.Unmarshal(body, resbody); err != nil {
			return err
		}
		for _, choice := range resbody.Choices {
			fmt.Printf("%s > %s\n", choice.Message.Role, strings.TrimSpace(choice.Message.Content))
		}
	}
}
