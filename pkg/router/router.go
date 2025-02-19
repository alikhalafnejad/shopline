package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

// Handler Enhanced Handler interface with middleware capability
type Handler interface {
	RegisterRoutes(r chi.Router)
}

// MiddlewareProvider interface for handlers that require specific middleware
type MiddlewareProvider interface {
	Handler
	Middlewares() []func(http.Handler) http.Handler
}

// RouteGroup defines a group of routes with common middleware and prefix
type RouteGroup struct {
	Prefix      string
	Middlewares []func(http.Handler) http.Handler
	Handlers    []Handler
}

// SetupRoutes sets up all routes with configurable middleware groups
func SetupRoutes(globalMiddlewares []func(http.Handler) http.Handler, routeGroups ...RouteGroup) *chi.Mux {
	r := chi.NewRouter()

	// Apply global middleware
	r.Use(globalMiddlewares...)

	// Configure route groups
	for _, group := range routeGroups {
		r.Group(func(r chi.Router) {
			// Apply group-specific middlewares
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

// Helper function to register handlers with their specific middlewares
func registerHandlers(r chi.Router, handlers []Handler) {
	for _, h := range handlers {
		// Check if handler provides specific middlewares
		if mp, ok := h.(MiddlewareProvider); ok {
			r.Group(func(r chi.Router) {
				r.Use(mp.Middlewares()...)
				mp.RegisterRoutes(r)
			})
		} else {
			h.RegisterRoutes(r)
		}
	}
}

//// SetupRoutes sets up all routes for the application
//func SetupRoutes(handlers ...Handler) *chi.Mux {
//	r := chi.NewRouter()
//
//	// Middleware
//	r.Use(middleware.RequestID)
//	r.Use(middleware.RealIP)
//	r.Use(middleware.Recoverer)
//	r.Use(middleware.Timeout(constants.TimeoutDuration))
//	r.Use(imiddleware.LoggingMiddleware)
//
//	// Versioned API routes
//	r.Route("/api", func(r chi.Router) {
//		// Version 1 routes
//		r.Route("/v1", func(r chi.Router) {
//			for _, h := range handlers {
//				h.RegisterRoutes(r)
//			}
//		})
//		// Future versions can be added here (e.g. , /v2, /v3)
//	})
//
//	return r
//}
