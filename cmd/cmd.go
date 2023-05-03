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
	&cmdModel,
	&cmdChat,
	&cmdImage,
	&cmdWhisper,
	&cmdFile,
	&cmdEdit,
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
			// }, {
			// 	Name:   "upload",
			// 	Action: file.Upload,
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
