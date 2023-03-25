package main

import (
	"log"
	"os"

	"github.com/lunarxlark/openai-go/cmd"
	"github.com/lunarxlark/openai-go/config"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "oai"
	app.Usage = "play open ai"
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
