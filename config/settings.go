package config

import (
	"github.com/joho/godotenv"
	"log"
	"time"
)

// Settings represents the centralized configuration for the application.
type Settings struct {
	Debug bool

	// Database Settings
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// JWT Settings
	JWTSecretKey string
	JWTDuration  time.Duration

	// Redis Settings
	RedisAddr     string
	RedisPassword string
	RedisDB       int

	// Pagination Defaults
	DefaultPage  int
	DefaultLimit int

	// Cache Settings
	CacheTTL time.Duration
}

// LoadSettings loads configuration values from environment variable or defaults
func LoadSettings() *Settings {
	// Load environment from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
