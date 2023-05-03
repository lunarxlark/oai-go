package chat

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	ff "github.com/ktr0731/go-fuzzyfinder"
	"github.com/lunarxlark/openai-cli/api"
	"github.com/lunarxlark/openai-cli/config"
	"github.com/urfave/cli/v2"
)

func CmdContinue(ctx *cli.Context) error {
	fs, err := ListLog()
	if err != nil {
		return err
	}

	i, err := ff.Find(fs, func(i int) string {
		return fs[i]
	}, ff.WithPreviewWindow(func(i, w, h int) string {
		fb, err := os.ReadFile(filepath.Join(config.OAIConfig.Dir, "chat", fs[i]))
		if err != nil {
			log.Fatal(fmt.Errorf("failed to read log files for preview. %w", err))
		}
		log.Println(string(fb))

		req := new(request)
		if err := json.Unmarshal(fb, &req); err != nil {
			log.Fatal(fmt.Errorf("failed to unmarshal json for log preview. %w", err))
		}
		var preview strings.Builder
		for _, m := range req.Messages {
			switch m.Role {
			case User:
				preview.WriteString(fmt.Sprintf("> %s\n", m.Content))
			default:
				preview.WriteString(m.Content + "\n")
			}
		}
		return fmt.Sprintf("Model : %s\nMessage : \n%s", req.Model, preview.String())
	}))
	if err != nil {
		return fmt.Errorf("failed to fuzzy-find for log file. %w", err)
	}

	fb, err := os.ReadFile(filepath.Join(config.OAIConfig.Dir, "chat", fs[i]))
	if err != nil {
		return err
	}

	req := new(request)
	if err := json.Unmarshal(fb, &req); err != nil {
		return err
	}

	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s > ", User)
		if !sc.Scan() {
			// conversation.Summary, err = CreateSummay(conversation.Model, conversation.Messages)
			b, err := json.MarshalIndent(req, "", "  ")
			if err != nil {
				return err
			}
			if err := os.WriteFile(filepath.Join(config.OAIConfig.Dir, "chat", fs[i]), b, 0644); err != nil {
				return err
			}
			break
		}
		req.Messages = append(req.Messages, Message{
			Role:    User,
			Content: sc.Text(),
		})

		var res response
		if err := api.Request(http.MethodPost, url, req, &res); err != nil {
			return err
		}

		for _, choice := range res.Choices {
			req.Messages = append(req.Messages, choice.Message)
			fmt.Printf("%s > %s\n", choice.Message.Role, choice.Message.Content)
		}
	}
	return nil
}

func ListLog() ([]string, error) {
	des, err := os.ReadDir(filepath.Join(config.OAIConfig.Dir, "chat"))
	if err != nil {
		return nil, err
	}

	fnames := []string{}
	for _, de := range des {
		if !de.IsDir() {
			fnames = append(fnames, de.Name())
		}
	}
	return fnames, nil
}
