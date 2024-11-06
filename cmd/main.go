package main

import (
	"github.com/oidentity/auth-server/internal/config"
	"github.com/oidentity/auth-server/internal/logger"
	"github.com/oidentity/auth-server/internal/server"
)

func main() {
	log := logger.GetLogger()

	config.LoadConfig()
	log.Info("Starting server...")

	srv := server.NewServer()
	srv.Start()
}
