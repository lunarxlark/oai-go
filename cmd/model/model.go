package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"text/tabwriter"

	ff "github.com/ktr0731/go-fuzzyfinder"
	"github.com/lunarxlark/openai-cli/config"
	"github.com/lunarxlark/openai-cli/models/model"
	"github.com/urfave/cli/v2"
)

const (
	url = "https://api.openai.com/v1/models"
)

func cmdList(ctx *cli.Context) error {
	modelID, err := List()
	if err != nil {
		return err
	}

	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet, url+"/"+modelID, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+config.OAIConfig.APIKey)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return err
}

func List() (string, error) {
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Bearer "+config.OAIConfig.APIKey)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var models model.ListAPIResponse
	if err := json.Unmarshal(body, &models); err != nil {
		return "", err
	}

	lines := []string{model.Header()}
	for _, data := range models.Data {
		lines = append(lines, data.String())
	}

	formatLines := format(lines)
	i, err := ff.Find(models.Data, func(i int) string {
		return formatLines[1+i]
	}, ff.WithHeader(formatLines[0]))
	if err != nil {
		return "", err
	}

	return models.Data[i].ID, nil
}

func format(lines []string) []string {
	var buf bytes.Buffer
	writer := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	for _, l := range lines {
		fmt.Fprintf(writer, "%s\n", l)
	}
	writer.Flush()
	return strings.Split(buf.String(), "\n")
}
