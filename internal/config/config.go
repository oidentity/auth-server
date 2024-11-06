package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port        string
	LogLevel    string
	Environment string
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	return Config{
		Port:        getEnv("PORT", "8080"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
