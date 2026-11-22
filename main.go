package main

import (
	"log"

	"github.com/felipedestro/todolist/api/internal/auth"
	"github.com/felipedestro/todolist/api/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	auth.SetupGoogle(cfg)
}
