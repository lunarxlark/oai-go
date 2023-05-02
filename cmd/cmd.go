package cmd

import (
	"github.com/lunarxlark/openai-cli/cmd/chat"
	"github.com/lunarxlark/openai-cli/cmd/edit"
	"github.com/lunarxlark/openai-cli/cmd/file"
	"github.com/lunarxlark/openai-cli/cmd/image"
	"github.com/lunarxlark/openai-cli/cmd/model"
	"github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
	cmdModel,
	cmdChat,
	cmdImage,
	cmdWhisper,
	cmdFile,
	cmdEdit,
}

var cmdModel = &cli.Command{
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

var cmdChat = &cli.Command{
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
		},
		{
			Name:        "continue",
			Aliases:     []string{"c"},
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

var cmdImage = &cli.Command{
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
}

var cmdWhisper = &cli.Command{
	Name:    "whisper",
	Aliases: []string{"w"},
}

var cmdFile = &cli.Command{
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
}

var cmdEdit = &cli.Command{
	Name:   "edit",
	Action: edit.Exec,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "model",
			Value: "text-davinci-edit-001",
		},
		&cli.StringFlag{
			Name:  "input",
			Value: "hello wrold!!",
		},
		&cli.StringFlag{
			Name:     "instruction",
			Required: true,
		},
		&cli.IntFlag{
			Name:  "n",
			Value: 1,
		},
		&cli.Float64Flag{
			Name:  "temperature",
			Value: 1,
		},
		&cli.Float64Flag{
			Name:  "top_p",
			Value: 1,
		},
	},
}
