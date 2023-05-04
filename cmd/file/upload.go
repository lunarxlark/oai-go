package file

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/lunarxlark/openai-cli/config"
	"github.com/urfave/cli/v2"
)

type Request struct {
	File    string `json:"file"`
	Purpose string `json:"purpose"`
}

type ResponseFileUpload struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int    `json:"created_at"`
	FileName  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

func Upload(ctx *cli.Context) error {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	f, err := os.Open(ctx.String("file"))
	if err != nil {
		return err
	}

	fileWriter, err := writer.CreateFormFile("file", ctx.String("file"))
	if err != nil {
		return err
	}
	if _, err := io.Copy(fileWriter, f); err != nil {
		return err
	}

	fieldWriter, err := writer.CreateFormField("purpose")
	if err != nil {
		return err
	}
	fieldWriter.Write([]byte(ctx.String("purpose")))

	writer.Close()

	req, err := http.NewRequest(http.MethodPost, url, &buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+config.OAIConfig.APIKey)

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var rs ResponseFileUpload
	if err := json.NewDecoder(res.Body).Decode(&rs); err != nil {
		return err
	}

	fmt.Println("success upload file")
	fmt.Printf("ID:%s\n", rs.ID)
	fmt.Printf("FileName:%s\n", rs.FileName)

	return nil
}
