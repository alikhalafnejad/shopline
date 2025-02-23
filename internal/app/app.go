package app

import (
	"shopline/config"
	"shopline/internal/handlers"
	"shopline/internal/repositories"
	"shopline/internal/services"
	"shopline/pkg/db"
	"shopline/pkg/redisdb"
)

type App struct {
	UserHandler    *handlers.UserHandler
	ProductHandler *handlers.ProductHandler
	AdminHandler   *handlers.AdminHandler
	CommentHandler *handlers.CommentHandler
}

func NewApp(settings *config.Settings) *App {
	// Initialize database and Redis
	database := db.InitDBWithPool(settings.DBHost, settings.DBPort, settings.DBUser, settings.DBPassword, settings.DBName)
	redisClient := redisdb.NewRedisClientWithPool(settings.RedisAddr, settings.RedisPassword, settings.RedisDB)

	// Initialize repositories
	userRepo := repositories.NewUserRepository(database)
	productRepo := repositories.NewProductRepository(database)
	adminRepo := repositories.NewAdminRepository(database)
	commentRepo := repositories.NewCommentRepository(database)

	// Initialize services
	userService := services.NewUserService(userRepo)
	productService := services.NewProductService(productRepo, redisClient)
	adminService := services.NewAdminService(adminRepo)
	commentService := services.NewCommentService(commentRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)
	adminHandler := handlers.NewAdminHandler(adminService)
	commentHandler := handlers.NewCommentHandler(commentService)

	return &App{
		UserHandler:    userHandler,
		ProductHandler: productHandler,
		AdminHandler:   adminHandler,
		CommentHandler: commentHandler,
	}
}
