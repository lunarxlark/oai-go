package image

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/png"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/lunarxlark/openai-cli/config"
	"github.com/urfave/cli/v2"
)

const (
	url = "https://api.openai.com/v1/images/generations"
)

type ImageReq struct {
	Prompt         string `json:"prompt"`          // A text description of the desired image(s). The maximum length is 1000 characters.
	N              int    `json:"n"`               // The number of images to generate. Must be between 1 and 10.
	Size           string `json:"size"`            // The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024.
	ResponseFormat string `json:"response_format"` // The format in which the generated images are returned. Must be one of url or b64_json.
	User           Role   `json:"user"`
}

type Role string

const (
	User Role = "user"
)

type ImageRes struct {
	Created int64 `json:"created"`
	Data    []struct {
		B46Json string `json:"b64_json"`
		URL     string `json:"url"`
	} `json:"data"`
}

func Exec(ctx *cli.Context) error {
	prompt := ctx.String("prompt")
	format := ctx.String("format")
	size := ctx.String("size")

	payload, err := json.Marshal(ImageReq{
		Prompt:         prompt,
		N:              1,
		Size:           size,
		ResponseFormat: format,
		User:           User,
	})
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
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

	imageRes := new(ImageRes)
	if err := json.Unmarshal(body, imageRes); err != nil {
		return err
	}

	// image decode from base64
	for i, data := range imageRes.Data {
		if format == "b64_json" {
			dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data.B46Json))
			img, err := png.Decode(dec)
			if err != nil {
				return err
			}

			if err := func() error {
				filename := fmt.Sprintf("%d_%02d.png", time.Now().Unix(), i)
				f, err := os.Create(filename)
				if err != nil {
					return err
				}
				defer f.Close()

				if err := png.Encode(f, img); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}

		if format == "url" {
			filename := fmt.Sprintf("%d_%02d.log", time.Now().Unix(), i)
			if err := func() error {
				f, err := os.Create(filename)
				if err != nil {
					return err
				}
				defer f.Close()

				if _, err := f.WriteString(data.URL); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}

		}
	}

	return nil
}
