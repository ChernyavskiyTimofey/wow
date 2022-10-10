package main

import (
	"log"

	"github.com/leprosus/wow/internal/config"
	"github.com/leprosus/wow/pkg/client"
)

func main() {
	cfg, err := config.ParseConfigFromCLI()
	if err != nil {
		log.Panicln(err)
	}

	cln := client.NewClient(cfg)
	cln.Run()
}
