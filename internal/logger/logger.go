package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewLogger() (*zap.Logger, error) {
	var cfg zap.Config

	if os.Getenv("APP_ENV") == "production" {
		cfg = zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	} else {
		cfg = zap.NewDevelopmentConfig()
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	defer logger.Sync()
	return logger, nil
}
