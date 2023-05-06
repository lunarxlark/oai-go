package image

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Command = cli.Command{
	Name:   "image",
	Action: cmdCreate,
	Flags: []cli.Flag{
		flagPrompt,
		flagN,
		flagSize,
		flagFormat,
		flagUser,
	},
}

var flagPrompt = &cli.StringFlag{
	Name:     "prompt",
	Required: true,
	Action: func(ctx *cli.Context, prompt string) error {
		if len(ctx.String("prompt")) > 1000 {
			return fmt.Errorf("prompt must be less than 1000 characters")
		}
		return nil
	},
}

var flagN = &cli.IntFlag{
	Name:  "n",
	Usage: "How many completions to generate for each prompt.",
	Value: 1,
}

var flagSize = &cli.StringFlag{
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

var flagFormat = &cli.StringFlag{
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

var flagUser = &cli.StringFlag{
	Name: "user",
}
