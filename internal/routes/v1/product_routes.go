package v1

import (
	"github.com/go-chi/chi/v5"
	"shopline/internal/handlers"
)

// ProductRoutes registers product-related routes.
func ProductRoutes(h *handlers.ProductHandler) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/products", h.GetProducts)                  // Get paginated list of products
		r.Get("/products/{id}", h.GetProduct)              // Get a single product by ID
		r.Post("/products", h.CreateProduct)               // Add a new product
		r.Put("/products/{id}", h.UpdateProduct)           // Update an existing product
		r.Delete("/products/{id}", h.DeleteProduct)        // Delete a product
		r.Get("/products/promoted", h.GetPromotedProducts) // Get promoted products
	}
}
