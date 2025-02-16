package app

import (
	"shopline/config"
	"shopline/internal/handlers"
	"shopline/internal/repositories"
	"shopline/internal/services"
	"shopline/pkg/db"
)

type App struct {
	UserHandler    *handlers.UserHandler
	ProductHandler *handlers.ProductHandler
}

func NewApp(settings *config.Settings) *App {
	// Initialize database and Redis
	database := db.InitDBWithPool(settings.DBHost, settings.DBPort, settings.DBUser, settings.DBPassword, settings.DBName)
	//redisClient := redisdb.NewRedisClientWithPool(settings.RedisAddr, settings.RedisPassword, settings.RedisDB)

	// Initialize repositories
	userRepo := repositories.NewUserRepository(database)
	productRepo := repositories.NewProductRepository(database)

	// Initialize services
	userService := services.NewUserService(userRepo)
	productService := services.NewProductService(productRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)

	return &App{
		UserHandler:    userHandler,
		ProductHandler: productHandler,
	}
}
