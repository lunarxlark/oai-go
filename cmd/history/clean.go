package history

import (
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func Clean(ctx *cli.Context) error {
	os.RemoveAll(filepath.Join(os.Getenv("XDG_CONFIG_HOME"), "oai", "chat", "*"))
	return nil
}
