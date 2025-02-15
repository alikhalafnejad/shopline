package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	imiddleware "shopline/internal/middleware"
	"shopline/pkg/constants"
)

// SetupRoutes sets up all routes for the application
func SetupRoutes(handlers ...Handler) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(constants.TimeoutDuration))
	r.Use(imiddleware.LoggingMiddleware)

	// Versioned API routes
	r.Route("/api", func(r chi.Router) {
		// Version 1 routes
		r.Route("v1", func(r chi.Router) {
			for _, h := range handlers {
				h.RegisterRoutes(r)
			}
		})
		// Future versions can be added here (e.g. , /v2, /v3)
	})

	return r
}
