package completion

import "github.com/urfave/cli/v2"

var Command = cli.Command{
	Name:   "completion",
	Action: cmdCompletion,
	Flags: []cli.Flag{
		flagModel,
		flagPrompt,
		flagN,
		flagStop,
		flagEcho,
		flagUser,
	},
}

var flagModel = &cli.StringFlag{
	Name:  "model",
	Value: "text-davinci-003",
}

var flagPrompt = &cli.StringFlag{
	Name: "prompt",
}

var flagN = &cli.IntFlag{
	Name:  "n",
	Usage: "How many completions to generate for each prompt.",
	Value: 1,
}

var flagStop = &cli.StringFlag{
	Name: "stop",
}

var flagEcho = &cli.BoolFlag{
	Name:  "echo",
	Usage: "Echo back the prompt in addition to the completion",
	Value: false,
}

var flagUser = &cli.StringFlag{
	Name: "user",
}
