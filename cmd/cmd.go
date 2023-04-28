package cmd

import (
	"github.com/lunarxlark/openai-cli/cmd/chat"
	"github.com/lunarxlark/openai-cli/cmd/file"
	"github.com/lunarxlark/openai-cli/cmd/history"
	"github.com/lunarxlark/openai-cli/cmd/image"
	"github.com/lunarxlark/openai-cli/cmd/model"
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
		Name:    "image",
		Aliases: []string{"i"},
		Action:  image.Exec,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "prompt",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "format",
				DefaultText: "b64_json",
				Value:       "b64_json",
				Usage:       "response format. {b64_json|url}",
			},
			&cli.StringFlag{
				Name:        "size",
				DefaultText: "1024x1024",
				Value:       "1024x1024",
				Usage:       "size of image. {256x265|512x512|1024x1024}",
			},
		},
	},
	{
		Name:    "whisper",
		Aliases: []string{"w"},
	},
	{
		Name: "file",
		Subcommands: []*cli.Command{
			{
				Name:   "list",
				Action: file.List,
				// }, {
				// 	Name:   "upload",
				// 	Action: file.Upload,
			},
		},
	},
	{
		Name:    "history",
		Aliases: []string{"h"},
		Action:  history.Clean,
	},
}
