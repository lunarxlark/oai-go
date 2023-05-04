package cmd

import (
	"fmt"

	"github.com/lunarxlark/openai-cli/cmd/chat"
	"github.com/lunarxlark/openai-cli/cmd/completion"
	"github.com/lunarxlark/openai-cli/cmd/edit"
	"github.com/lunarxlark/openai-cli/cmd/file"
	"github.com/lunarxlark/openai-cli/cmd/image"
	"github.com/lunarxlark/openai-cli/cmd/model"
	"github.com/lunarxlark/openai-cli/cmd/moderation"
	"github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
	&cmdModel,
	&cmdCompletion,
	&cmdChat,
	&cmdImage,
	&cmdWhisper,
	&cmdFile,
	&cmdEdit,
	&cmdModeration,
}

var cmdModel = cli.Command{
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
}

var cmdCompletion = cli.Command{
	Name:   "completion",
	Action: completion.Exec,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "model",
			Value: "text-davinci-003",
		},
		&cli.StringFlag{
			Name: "prompt",
		},
		FlagN,
		&cli.StringFlag{
			Name: "stop",
		},
		&cli.BoolFlag{
			Name:  "echo",
			Usage: "Echo back the prompt in addition to the completion",
			Value: false,
		},
		&cli.StringFlag{
			Name: "user",
		},
	},
}

var cmdChat = cli.Command{
	Name:    "chat",
	Aliases: []string{"c"},
	Subcommands: []*cli.Command{
		{
			Name:        "new",
			Description: "new chat",
			Action:      chat.CmdNew,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "model",
					DefaultText: "gpt-3.5-turbo-0301",
				},
			},
		},
		{
			Name:        "continue",
			Description: "continue chat",
			Action:      chat.CmdContinue,
		},
		{
			Name:    "clear",
			Aliases: []string{"h"},
			Action:  chat.Clear,
		},
	},
}

var cmdImage = cli.Command{
	Name:    "image",
	Aliases: []string{"i"},
	Action:  image.Exec,
	Flags: []cli.Flag{
		FlagPrompt,
		FlagN,
		FlagSize,
		FlagFormat,
		FlagUser,
	},
}

var cmdWhisper = cli.Command{
	Name:    "whisper",
	Aliases: []string{"w"},
}

var cmdFile = cli.Command{
	Name: "file",
	Subcommands: []*cli.Command{
		{
			Name:   "list",
			Action: file.List,
		},
		{
			Name:   "upload",
			Action: file.Upload,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "file",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "purpose",
					Required: true,
				},
			},
		},
		{
			Name:   "delete",
			Action: file.Delete,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "file_id",
					Required: true,
				},
			},
		},
	},
}

var cmdEdit = cli.Command{
	Name:   "edit",
	Action: edit.Exec,
	Flags: []cli.Flag{
		FlagModel,
		&cli.StringFlag{
			Name:  "input",
			Value: "hello wrold!!",
		},
		FlagInstruction,
		FlagN,
		FlagTemperature,
		FlagTopP,
	},
}

var cmdModeration = cli.Command{
	Name:   "moderation",
	Action: moderation.Exec,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "input",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "model",
			Value: "text-moderation-stable",
			Action: func(ctx *cli.Context, model string) error {
				if model != "text-moderation-stable" && model != "text-moderation-latest" {
					return fmt.Errorf("model must be 'text-moderation-stable' or 'text-moderation-latest'")
				}
				return nil
			},
		},
	},
}
