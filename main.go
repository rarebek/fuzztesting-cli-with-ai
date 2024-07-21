package main

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rarebek/fuzztesting-cli-with-ai/cmd"
	"github.com/rarebek/fuzztesting-cli-with-ai/config"
)

func main() {
	cmd.Execute()

	var cfg config.Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatal(err)
	}

}
