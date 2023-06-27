package main

import (
	"MateMind/configs"
	"MateMind/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
