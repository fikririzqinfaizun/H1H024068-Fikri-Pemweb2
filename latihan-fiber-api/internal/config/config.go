package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	AppPort string
	AppEnv  string

	DbHost string
	DbPort string
	DbName string
	DbUser string
	DbPass string
}

func LoadConfig() Config {
	_ = godotenv.Load()

	return Config{
		AppName: getEnv("APP_NAME", "Pemweb2 Fiber API"),
		AppPort: getEnv("APP_PORT", "3000"),
		AppEnv:  getEnv("APP_ENV", "local"),

		DbHost: getEnv("DB_HOST", "127.0.0.1"),
		DbPort: getEnv("DB_PORT", "3306"),
		DbName: getEnv("DB_NAME", "pemweb2_fiber"),
		DbUser: getEnv("DB_USER", "root"),
		DbPass: getEnv("DB_PASS", ""),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
