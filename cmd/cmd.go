package cmd

import (
	"github.com/lunarxlark/openai-go/cmd/chat"
	"github.com/lunarxlark/openai-go/cmd/model"
	"github.com/lunarxlark/openai-go/cmd/whisper"
	"github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
	{
		Name:    "chat",
		Aliases: []string{"c"},
		Subcommands: []*cli.Command{
			{
				Name:        "new",
				Aliases:     []string{"n"},
				Description: "new chat",
				Action:      chat.CmdNew,
			}, {
				Name:        "edit",
				Aliases:     []string{"e"},
				Description: "edit chat",
				Action:      chat.CmdEdit,
			},
		},
	},
	{
		Name:    "whisper",
		Aliases: []string{"w"},
		Action:  whisper.Exec,
	},
	{
		Name:    "model",
		Aliases: []string{"m"},
		Subcommands: []*cli.Command{
			{
				Name:        "list",
				Aliases:     []string{"l"},
				Description: "list model",
				Action:      model.CmdList,
			},
		},
	},
}
