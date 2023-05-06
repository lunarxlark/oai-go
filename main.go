package main

import (
	"log"
	"os"

	"github.com/lunarxlark/openai-cli/cmd/chat"
	"github.com/lunarxlark/openai-cli/cmd/completion"
	"github.com/lunarxlark/openai-cli/cmd/edit"
	"github.com/lunarxlark/openai-cli/cmd/file"
	"github.com/lunarxlark/openai-cli/cmd/image"
	"github.com/lunarxlark/openai-cli/cmd/model"
	"github.com/lunarxlark/openai-cli/cmd/moderation"
	"github.com/lunarxlark/openai-cli/config"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "openai-cli"
	app.Usage = "play openai"
	app.Version = "0.0.1"
	app.Commands = commands
	app.Before = before

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func before(ctx *cli.Context) error {
	config.Load()
	return nil
}

var commands = []*cli.Command{
	&model.Command,
	&completion.Command,
	&chat.Command,
	&image.Command,
	&file.Command,
	&edit.Command,
	&moderation.Command,
}
