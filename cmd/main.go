package main

import (
	"github.com/oidentity/auth-server/internal/config"
	"github.com/oidentity/auth-server/internal/db"
	"github.com/oidentity/auth-server/internal/logger"
	"github.com/oidentity/auth-server/internal/server"
)

func main() {
	log := logger.GetLogger()

	config.LoadConfig()
	log.Info("Starting server...")

	db.ConnectPostgres()

	srv := server.NewServer()
	srv.Start()
}
