package main

import (
	"log"

	"github.com/leprosus/wow/internal/config"
	"github.com/leprosus/wow/pkg/server"
)

func main() {
	cfg, err := config.ParseConfigFromCLI()
	if err != nil {
		log.Panicln(err)
	}

	srv := server.NewServer(cfg)

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
