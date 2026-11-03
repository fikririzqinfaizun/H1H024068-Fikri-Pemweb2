package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	AppPort string
	AppEnv  string
}

func LoadConfig() Config {
	_ = godotenv.Load()

	return Config{
		AppName: getEnv("APP_NAME", "Pemweb2 Fiber API"),
		AppPort: getEnv("APP_PORT", "3000"),
		AppEnv:  getEnv("APP_ENV", "local"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
