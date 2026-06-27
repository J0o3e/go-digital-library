package config

import (
	"os"

	"github.com/joho/godotenv"

	"go-digital-library/internal/database"
)

func LoadDatabaseConfig() database.Config {
	_ = godotenv.Load()

	return database.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", ""),
		Name:     getEnv("DB_NAME", "go_digital_library"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}