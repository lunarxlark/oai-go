package config

import (
	"log"
	"os"
)

type Config struct {
	Dir    string
	ApiKey string
}

var OAIConfig Config

func Load() {
	OAIConfig.Dir = "~/.config/openai-go/"
	OAIConfig.ApiKey = os.Getenv("OPENAI_API_KEY")

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
