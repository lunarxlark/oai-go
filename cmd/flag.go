package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var FlagPrompt = &cli.StringFlag{
	Name:     "prompt",
	Required: true,
	Action: func(ctx *cli.Context, prompt string) error {
		if len(ctx.String("prompt")) > 1000 {
			return fmt.Errorf("prompt must be less than 1000 characters")
		}
		return nil
	},
}

var FlagN = &cli.IntFlag{
	Name:  "n",
	Usage: "How many completions to generate for each prompt.",
	Value: 1,
}

var FlagTemperature = &cli.Float64Flag{
	Name:  "temperature",
	Value: 1,
}

var FlagTopP = &cli.Float64Flag{
	Name:  "top_p",
	Value: 1,
}

var FlagUser = &cli.StringFlag{
	Name: "user",
}

var FlagFormat = &cli.StringFlag{
	Name:  "format",
	Value: "b64_json",
	Usage: "response format. {b64_json|url}",
	Action: func(ctx *cli.Context, format string) error {
		if format != "url" && format != "b64_json" {
			return fmt.Errorf("format must be url or b64_json")
		}
		return nil
	},
}

var FlagSize = &cli.StringFlag{
	Name:  "size",
	Value: "1024x1024",
	Usage: "size of image. {256x265|512x512|1024x1024}",
	Action: func(ctx *cli.Context, size string) error {
		if size != "256x256" && size != "512x512" && size != "1024x1024" {
			return fmt.Errorf("size must be 256x256, 512x512 or 1024x1024")
		}
		return nil
	},
}

var FlagInstruction = &cli.StringFlag{
	Name:     "instruction",
	Required: true,
}

var FlagModel = &cli.StringFlag{
	Name:  "model",
	Value: "text-davinci-edit-001",
}
