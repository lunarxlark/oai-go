package chat

import (
	"os"
	"path/filepath"

	"github.com/lunarxlark/openai-cli/config"
	"github.com/urfave/cli/v2"
)

func Clear(ctx *cli.Context) error {
	os.RemoveAll(filepath.Join(config.OAIConfig.Dir, "chat", "*"))
	return nil
}
