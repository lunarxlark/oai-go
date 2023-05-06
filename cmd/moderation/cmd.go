package moderation

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Command = cli.Command{
	Name:   "moderation",
	Action: cmdModeration,
	Flags: []cli.Flag{
		flagInput,
		flagModel,
	},
}

var flagInput = &cli.StringFlag{
	Name:     "input",
	Required: true,
}

var flagModel = &cli.StringFlag{
	Name:  "model",
	Value: "text-moderation-stable",
	Action: func(ctx *cli.Context, model string) error {
		if model != "text-moderation-stable" && model != "text-moderation-latest" {
			return fmt.Errorf("model must be 'text-moderation-stable' or 'text-moderation-latest'")
		}
		return nil
	},
}
