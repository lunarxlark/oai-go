package config

import (
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Dir    string
	APIKey string
}

var OAIConfig Config

func Load() {
	OAIConfig.Dir = filepath.Join(os.Getenv("XDG_DATA_HOME"), "openai-cli")
	OAIConfig.APIKey = os.Getenv("OPENAI_API_KEY")

	dir, err := os.Open(OAIConfig.Dir)
	if err != nil {
		log.Fatal(err)
	}

	fs, err := os.Stat(dir.Name())
	if err != nil {
		log.Fatal(err)
	}

	if !fs.IsDir() {
		log.Fatal("there is ")
	}
}
