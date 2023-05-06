package model

import (
	"github.com/urfave/cli/v2"
)

var Command = cli.Command{
	Name:    "model",
	Aliases: []string{"m"},
	Subcommands: []*cli.Command{
		{
			Name:        "list",
			Aliases:     []string{"l"},
			Description: "list model",
			Action:      cmdList,
		},
	},
}
