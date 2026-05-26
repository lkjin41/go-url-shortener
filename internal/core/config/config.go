package core_config

import (
	"os"
	"strconv"
)

type RedisConfig struct {
	Host     string `json:"addr"`
	Port     string
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type Config struct {
	Redis RedisConfig `json:"redis"`
}

func Load() *Config {
	return &Config{
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
	}
}

// getEnv get the value of the environment as string variable or return the default value if not set
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt get the value of the environment as integer variable or return the default value
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
