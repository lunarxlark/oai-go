package chat

import (
	"strings"

	"github.com/lunarxlark/openai-cli/api"
)

type Conversation struct {
	Summary  string        `json:"summary"`
	Model    string        `json:"model"`
	Messages []api.Message `json:"messages"`
}

func CreateSummay(model string, messages []api.Message) (string, error) {
	conversation := append(messages, api.Message{
		Role:    api.User,
		Content: "make one title for these conversatoins with less 20 characters",
	})
	res, err := api.CreateReq(model, conversation).Request()
	if err != nil {
		return "", err
	}
	summary := ""
	for _, choice := range res.Choices {
		summary = strings.Trim(choice.Message.Content, "\"")
	}
	return summary, nil
}
