package main

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rarebek/fuzzer/cmd"
	"github.com/rarebek/fuzzer/config"
)

func main() {
	cmd.Execute()

	var cfg config.Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatal(err)
	}

}
