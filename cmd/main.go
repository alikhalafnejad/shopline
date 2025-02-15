package cmd

import (
	"go.uber.org/zap"
	"shopline/config"
	"shopline/pkg/db"
	"shopline/pkg/logger"
	"shopline/pkg/redisdb"
)

func main() {
	// Load Settings
	settings := config.LoadSettings()

	// Initialize Logger
	logger.InitLogger(settings.Debug)
	defer logger.Sync()

	// Load application startup
	logger.Logger.Info("Starting application")

	// Initialize database connection
	database := db.InitDB()
	defer func() {
		sqlDB, _ := database.DB()
		if err := sqlDB.Close(); err != nil {
			logger.Logger.Error("Failed to close database connection", zap.Error(err))
		}
	}()
	logger.Logger.Info("Connected to database", zap.String("db_host", settings.DBHost))

	// Initialize Redis client
	redisClient := redisdb.NewRedisClient()
	defer func() {
		if err := redisClient.Client.Close(); err != nil {
			logger.Logger.Error("Failed to close redis connection", zap.Error(err))
		}
	}()
	logger.Logger.Info("Connected to Redis", zap.String("redis_addr", settings.RedisAddr))

	// Initialize repositories

}
