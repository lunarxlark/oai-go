package chat

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	ff "github.com/ktr0731/go-fuzzyfinder"
	"github.com/lunarxlark/openai-cli/api"
	"github.com/urfave/cli/v2"
)

func CmdContinue(ctx *cli.Context) error {
	fs, err := ListLog()
	if err != nil {
		return err
	}

	i, err := ff.Find(fs, func(i int) string {
		return fs[i]
	},
		ff.WithPreviewWindow(func(i, w, h int) string {
			fb, err := os.ReadFile(filepath.Join(os.Getenv("XDG_CONFIG_HOME"), "openai-cli", "chat", fs[i]))
			if err != nil {
				log.Fatal(fmt.Errorf("failed to read log files for preview. %w", err))
			}
			log.Println(string(fb))

			conversation := new(Conversation)
			if err := json.Unmarshal(fb, &conversation); err != nil {
				log.Fatal(fmt.Errorf("failed to unmarshal json for log preview. %w", err))
			}
			var preview strings.Builder
			for _, m := range conversation.Messages {
				switch m.Role {
				case api.User:
					preview.WriteString(fmt.Sprintf("> %s\n", m.Content))
				default:
					preview.WriteString(m.Content + "\n")
				}
			}
			return fmt.Sprintf("Summary : %s\nModel : %s\nMessage : \n%s", conversation.Summary, conversation.Model, preview.String())
		}))
	if err != nil {
		return fmt.Errorf("failed to fuzzy-find for log file. %w", err)
	}

	fb, err := os.ReadFile(filepath.Join(os.Getenv("XDG_CONFIG_HOME"), "openai-cli", "chat", fs[i]))
	if err != nil {
		return err
	}

	conversation := new(Conversation)
	if err := json.Unmarshal(fb, &conversation); err != nil {
		return err
	}

	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s > ", api.User)
		if !sc.Scan() {
			conversation.Summary, err = CreateSummay(conversation.Model, conversation.Messages)
			b, err := json.MarshalIndent(conversation, "", "  ")
			if err != nil {
				return err
			}
			if err := os.WriteFile(filepath.Join(os.Getenv("XDG_CONFIG_HOME"), "openai-cli", "chat", fs[i]), b, 0644); err != nil {
				return err
			}
			break
		}
		statement := []api.Message{{
			Role:    api.User,
			Content: sc.Text(),
		}}
		conversation.Messages = append(conversation.Messages, statement...)

		res, err := api.CreateReq(conversation.Model, conversation.Messages).Request()
		if err != nil {
			return err
		}

		for _, choice := range res.Choices {
			fmt.Println(choice.Message.String())
			conversation.Messages = append(conversation.Messages, choice.Message)
		}
	}
	return nil
}

func ListLog() ([]string, error) {
	des, err := os.ReadDir(filepath.Join(os.Getenv("XDG_CONFIG_HOME"), "openai-cli", "chat"))
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
