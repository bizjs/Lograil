package config

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort  string
	DatabaseURL string
	RedisURL    string
	JWTSecret   string
	Environment string
}

func Load() (*Config, error) {
	cfg := &Config{
		ServerPort:  getEnv("SERVER_PORT", "9012"),
		DatabaseURL: getEnv("DATABASE_URL", "file:lograil.db?cache=shared&_fk=1"),
		RedisURL:    getEnv("REDIS_URL", "redis://localhost:6379"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
