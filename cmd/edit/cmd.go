package edit

import "github.com/urfave/cli/v2"

var Command = cli.Command{
	Name:   "edit",
	Action: cmdEdit,
	Flags: []cli.Flag{
		flagModel,
		flagInput,
		flagInstruction,
		flagN,
		flagTemperature,
		flagTopP,
	},
}

var flagModel = &cli.StringFlag{
	Name:  "model",
	Value: "text-davinci-edit-001",
}

var flagInput = &cli.StringFlag{
	Name:  "input",
	Value: "hello wrold!!",
}

var flagInstruction = &cli.StringFlag{
	Name:     "instruction",
	Required: true,
}

var flagN = &cli.IntFlag{
	Name:  "n",
	Usage: "How many completions to generate for each prompt.",
	Value: 1,
}

var flagTemperature = &cli.Float64Flag{
	Name:  "temperature",
	Value: 1,
}

var flagTopP = &cli.Float64Flag{
	Name:  "top_p",
	Value: 1,
}
