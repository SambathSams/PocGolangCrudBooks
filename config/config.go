package config

import (
	"go-crud-backend/logger"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
}

const fallbackPort = "8080"

func LoadConfig() *Config {
	// 1. Determine which environment we are in (default to development)
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// 2. Load the specific .env file based on APP_ENV
	// e.g., loads .env.production or .env.development
	err := godotenv.Load(".env." + env)
	if err != nil {
		logger.Warn(".env.%s file not found, using system env", env)
	}

	return &Config{
		Port: getEnv("PORT", fallbackPort),
		Env:  env,
	}
}

// Helper to provide default values
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
