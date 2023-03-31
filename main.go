package main

import (
	"log"
	"os"

	"github.com/lunarxlark/openai-cli/cmd"
	"github.com/lunarxlark/openai-cli/config"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "openai-cli"
	app.Usage = "play openai"
	app.Version = "0.0.1"
	app.Commands = cmd.Commands
	app.Before = before

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var API_KEY string

func before(ctx *cli.Context) error {
	config.Load()
	return nil
}
