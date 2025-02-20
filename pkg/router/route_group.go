package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

// Handler defines the interface for registering routes.
type Handler interface {
	RegisterRoutes(r chi.Router)
}

// MiddlewareProvider defines the interface for handlers that provide their own middleware.
type MiddlewareProvider interface {
	Handler
	Middlewares() []func(http.Handler) http.Handler
}

// RouteGroup defines a group of routes with a common prefix and middleware.
type RouteGroup struct {
	Prefix      string
	Middlewares []func(http.Handler) http.Handler
	Handlers    []Handler
}
