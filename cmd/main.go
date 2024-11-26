package main

import (
	"Restv2/internal/config"
	"Restv2/internal/logger"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupZapLogger(cfg.Env)

	router := mux.NewRouter()

	// TODO: init database

	// TODO: init server

	// TODO: run server
}
