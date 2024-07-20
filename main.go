package main

import (
	"fuzztester-cli/cmd"
	"fuzztester-cli/config"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	cmd.Execute()

	var cfg config.Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatal(err)
	}

}
