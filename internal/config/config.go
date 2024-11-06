package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port        string
	LogLevel    string
	Environment string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
}

func LoadConfig() Config {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	return Config{
		Port:        viper.GetString("PORT"),
		LogLevel:    viper.GetString("LOG_LEVEL"),
		Environment: viper.GetString("APP_ENV"),
		DBHost:      viper.GetString("DB_HOST"),
		DBPort:      viper.GetString("DB_PORT"),
		DBUser:      viper.GetString("DB_USER"),
		DBPassword:  viper.GetString("DB_PASSWORD"),
		DBName:      viper.GetString("DB_NAME"),
	}
}
