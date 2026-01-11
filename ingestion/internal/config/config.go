package config

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort      string
	VictoriaLogsURL string
	RedisURL        string
	BatchSize       int
	BufferSize      int
	Environment     string
}

func Load() (*Config, error) {
	cfg := &Config{
		ServerPort:      getEnv("SERVER_PORT", "8081"),
		VictoriaLogsURL: getEnv("VICTORIA_LOGS_URL", "http://localhost:9428"),
		RedisURL:        getEnv("REDIS_URL", "redis://localhost:6379"),
		BatchSize:       getEnvAsInt("BATCH_SIZE", 100),
		BufferSize:      getEnvAsInt("BUFFER_SIZE", 1000),
		Environment:     getEnv("ENVIRONMENT", "development"),
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
