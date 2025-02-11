package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"shopline/pkg/constants"
	"strconv"
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

	debug, _ := strconv.ParseBool(getEnv("DEBUG", "true"))

	return &Settings{
		Debug: debug,

		// Database Settings
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "myuser"),
		DBPassword: getEnv("DB_PASSWORD", "mypassword"),
		DBName:     getEnv("DB_NAME", "mydb"),

		// JWT Settings
		JWTSecretKey: getEnv("JWT_SECRET_KEY", "my_secret_key"),
		JWTDuration:  constants.JWTDuration,

		// Redis Settings
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getIntEnv("REDIS_DB", 0),

		// Pagination Defaults
		DefaultPage:  getIntEnv("DEFAULT_PAGE", 1),
		DefaultLimit: getIntEnv("DEFAULT_LIMIT", 10),

		// Cache Settings
		CacheTTL: constants.RedisCacheDuration,
	}

}

// Helper function to get environment variables with fallbacks
func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getIntEnv(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return valueInt
}
