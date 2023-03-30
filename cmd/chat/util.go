package chat

import (
	"strings"
)

type Conversation struct {
	Summary  string    `json:"summary"`
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

func CreateSummay(model string, messages []Message) (string, error) {
	conversation := append(messages, Message{
		Role:    user,
		Content: "make one title for these conversatoins with less 20 characters",
	})
	res, err := CreateReq(model, conversation).Request()
	if err != nil {
		return "", err
	}
	summary := ""
	for _, choice := range res.Choices {
		summary = strings.Trim(choice.Message.Content, "\"")
	}
	return summary, nil
}
