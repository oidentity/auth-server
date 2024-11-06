package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oidentity/auth-server/internal/config"
	"github.com/oidentity/auth-server/internal/logger"
	"go.uber.org/zap"
)

type Server struct {
	Router *gin.Engine
}

func NewServer() *Server {
	config := config.LoadConfig()
	gin.SetMode(gin.DebugMode)
	router := gin.New()

	// Middleware for logging
	if config.LogLevel == "debug" {
		router.Use(gin.Logger())
	}
	router.Use(gin.Recovery())

	return &Server{
		Router: router,
	}
}

func (s *Server) Start() {
	log := logger.GetLogger()
	// Define routes
	s.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	srv := &http.Server{
		Addr:    ":" + config.LoadConfig().Port,
		Handler: s.Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: %s\n", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", zap.Error(err))
	}

	log.Info("Server exiting")
}
