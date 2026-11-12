package config

import (
	"os"
	"strconv"

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

	JwtSecret   string
	JwtTtlMenit int
}

func LoadConfig() Config {
	_ = godotenv.Load()

	ttl, err := strconv.Atoi(getEnv("JWT_TTL_MENIT", "120"))
	if err != nil || ttl <= 0 {
		ttl = 120
	}

	return Config{
		AppName: getEnv("APP_NAME", "Pemweb2 Fiber API"),
		AppPort: getEnv("APP_PORT", "3000"),
		AppEnv:  getEnv("APP_ENV", "local"),

		DbHost: getEnv("DB_HOST", "127.0.0.1"),
		DbPort: getEnv("DB_PORT", "3306"),
		DbName: getEnv("DB_NAME", "pemweb2_fiber"),
		DbUser: getEnv("DB_USER", "root"),
		DbPass: getEnv("DB_PASS", ""),

		JwtSecret:   getEnv("JWT_SECRET", "rahasia-bawaan"),
		JwtTtlMenit: ttl,
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
