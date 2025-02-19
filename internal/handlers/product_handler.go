package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"shopline/internal/services"
	"strconv"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) RegisterRoutes(r chi.Router) {

}

// GetProducts retrieves paginated products with optional filters
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters for pagination and filtering
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	categoryID, _ := strconv.Atoi(r.URL.Query().Get("category_id"))
	minPrice, _ := strconv.ParseFloat(r.URL.Query().Get("min_price"), 64)
	maxPrice, _ := strconv.ParseFloat(r.URL.Query().Get("max_price"), 64)

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	// Fetch products with filters
	products, err := h.productService.GetProducts(page, limit)
}
