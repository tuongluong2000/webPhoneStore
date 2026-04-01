package config

import (
	"os"
)

type Config struct {
	Port      string
	JWTSecret string
	DBUser    string
	DBPass    string
	DBName    string
	DBHost    string
}

func LoadConfig() Config {
	return Config{
		Port:      getEnv("PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET", "secret"),
		DBUser:    getEnv("DB_USER", "root"),
		DBPass:    getEnv("DB_PASS", ""),
		DBName:    getEnv("DB_NAME", "shop"),
		DBHost:    getEnv("DB_HOST", "localhost"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
