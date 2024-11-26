package main

import (
	"Restv2/internal/config"
	"Restv2/internal/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupZapLogger(cfg.Env)

	// TODO: init database

	// TODO: init router

	// TODO: init server

	// TODO: run server
}
