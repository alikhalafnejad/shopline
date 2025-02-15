package v1

import (
	"github.com/go-chi/chi/v5"
	"shopline/internal/handlers"
)

// AuthRoutes defines versioned routes for authentication
func AuthRoutes(handler *handlers.UserHandler) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/register", handler.Register)
		r.Post("/login", handler.Login)
	}
}
