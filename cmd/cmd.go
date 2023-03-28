package cmd

import (
	"github.com/lunarxlark/oai-go/cmd/chat"
	"github.com/lunarxlark/oai-go/cmd/model"
	"github.com/lunarxlark/oai-go/cmd/whisper"
	"github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
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
				Name:        "continue",
				Aliases:     []string{"c"},
				Description: "continue chat",
				Action:      chat.CmdContinue,
			},
		},
	},
	{
		Name:    "whisper",
		Aliases: []string{"w"},
		Action:  whisper.Exec,
	},
}
