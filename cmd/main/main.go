package main

import (
	"projects/go-auth/config"
	"projects/go-auth/internal/logger"
)

func main() {
	// Init config.
	cfg := config.InitConfig()

	// Get logger.
	logger := logger.NewLogger(cfg)

}
