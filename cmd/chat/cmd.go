package chat

import "github.com/urfave/cli/v2"

var Command = cli.Command{
	Name: "chat",
	Subcommands: []*cli.Command{
		{
			Name:   "new",
			Action: cmdNew,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "model",
					DefaultText: "gpt-3.5-turbo-0301",
				},
			},
		},
		{
			Name:   "continue",
			Action: cmdContinue,
		},
		{
			Name:   "clear",
			Action: cmdClear,
		},
	},
}
