package handlers

import (
	"github.com/go-chi/chi/v5"
	"shopline/internal/services"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) RegisterRoutes(r chi.Router) {

}
