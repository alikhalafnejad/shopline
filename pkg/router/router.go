package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

// SetupRoutes sets up all routes with configurable middleware groups.
func SetupRoutes(globalMiddlewares []func(http.Handler) http.Handler, routeGroups ...RouteGroup) *chi.Mux {
	r := chi.NewRouter()

	// Apply global middleware
	r.Use(append([]func(http.Handler) http.Handler{
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
		middleware.Timeout(60 * time.Second),
	}, globalMiddlewares...)...)

	// Register route groups
	for _, group := range routeGroups {
		r.Group(func(r chi.Router) {
			// Apply group-specific middleware
			r.Use(group.Middlewares...)

			// Apply path prefix if specified
			if group.Prefix != "" {
				r.Route(group.Prefix, func(r chi.Router) {
					registerHandlers(r, group.Handlers)
				})
			} else {
				registerHandlers(r, group.Handlers)
			}
		})
	}

	return r
}

// Helper function to register handlers with their specific middleware.
func registerHandlers(r chi.Router, handlers []Handler) {
	for _, handler := range handlers {
		if mp, ok := handler.(MiddlewareProvider); ok {
			// If the handler provides its own middleware, wrap it in a group
			r.Group(func(r chi.Router) {
				r.Use(mp.Middlewares()...)
				mp.RegisterRoutes(r)
			})
		} else {
			// Otherwise, directly register the handler's routes
			handler.RegisterRoutes(r)
		}
	}
}
