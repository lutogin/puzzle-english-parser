package main

import (
	"go-pe-parser/src/config"
	"log"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	start(cfg)
}
