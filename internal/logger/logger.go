package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

func init() {
	var cfg zap.Config

	env := os.Getenv("APP_ENV")
	if env == "production" {
		cfg = zap.NewProductionConfig()
		cfg.Level.SetLevel(zap.InfoLevel)
	} else {
		cfg = zap.NewDevelopmentConfig()
		cfg.Level.SetLevel(zap.DebugLevel)
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	Logger, err = cfg.Build()
	if err != nil {
		Logger.Error("Failed to initialize logger", zap.Error(err))
	}

	defer Logger.Sync()
}

func GetLogger() *zap.Logger {
	return Logger
}
