package routes

import "github.com/go-chi/chi/v5"

// Handler defines the interface for route registration.
type Handler interface {
	RegisterRoutes(r chi.Router)
}
