package cmd

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"shopline/config"
	"shopline/internal/app"
	"shopline/pkg/logger"
	"shopline/pkg/router"
	"syscall"
	"time"
)

func main() {
	// Load Settings
	settings := config.LoadSettings()

	// Initialize Logger
	logger.InitLogger(settings.Debug)
	defer logger.Sync()

	// Load application startup
	logger.Logger.Info("Starting application")

	// Initialize the application
	newApp := app.NewApp(settings)

	// Register routes using chi
	r := router.SetupRoutes(
		newApp.UserHandler,
		newApp.ProductHandler,
	)

	// Create HTTP sever
	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start the server in a goroutine
	go func() {
		logger.Logger.Info("Server started", zap.String("address", server.Addr))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Logger.Fatal("Server failed", zap.Error(err))
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Logger.Error("Server forced to shutdown", zap.Error(err))
	}

	logger.Logger.Info("Server exiting")
}
