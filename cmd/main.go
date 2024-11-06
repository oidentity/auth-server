package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/oidentity/auth-server/internal/logger"
	"go.uber.org/zap"
)

func main() {
    log, err := logger.NewLogger()
    if err != nil {
        log.Fatal("Failed to initialize logger", zap.Error(err))
    }

    log.Info("Starting the server...")

    log.Info("Server configuration",
        zap.String("port", os.Getenv("SERVER_PORT")),
        zap.String("env", os.Getenv("APP_ENV")),
    )

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Info("Received request", zap.String("method", r.Method), zap.String("path", r.URL.Path))
        fmt.Fprintf(w, "Hello, world!")
    })

    port := os.Getenv("SERVER_PORT")

    log.Info("Listening on port", zap.String("port", port))
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal("Server failed to start", zap.Error(err))
    }
}
