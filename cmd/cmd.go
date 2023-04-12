package cmd

import (
	"github.com/lunarxlark/openai-cli/cmd/chat"
	"github.com/lunarxlark/openai-cli/cmd/history"
	"github.com/lunarxlark/openai-cli/cmd/model"
	"github.com/lunarxlark/openai-cli/cmd/whisper"
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
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "model",
						DefaultText: "gpt-3.5-turbo-0301",
					},
				},
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
	{
		Name:    "history",
		Aliases: []string{"h"},
		Action:  history.Clean,
	},
}
