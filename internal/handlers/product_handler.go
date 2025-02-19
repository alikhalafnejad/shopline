package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"shopline/internal/models"
	"shopline/internal/services"
	"shopline/pkg/response"
	"shopline/pkg/validation"
	"strconv"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) RegisterRoutes(r chi.Router) {
	r.Route("/products", func(r chi.Router) {
		r.Get("/", h.GetProducts)
		r.Get("/{id}", h.GetProduct)
		r.Post("/", h.CreateProduct)
		r.Put("/{id}", h.UpdateProduct)
		r.Delete("/{id}", h.DeleteProduct)
		r.Get("/promoted", h.GetPromotedProducts)
	})
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
	products, err := h.productService.GetProducts(page, limit, categoryID, minPrice, maxPrice)
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "Failed to fetch products")
		return
	}

	response.RespondSuccess(w, http.StatusOK, "Products retrieved successfully", products)
}

// GetProduct retrieves a single product by ID.
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.RespondError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "Product not found")
		return
	}

	response.RespondSuccess(w, http.StatusOK, "Product retrieved successfully", product)
}

// CreateProduct creates a new product.
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		response.RespondError(w, http.StatusBadRequest, "Invalid product data")
		return
	}

	// Validate the product data
	if err := validation.ValidateStruct(product); err != nil {
		response.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := h.productService.CreateProduct(&product)
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "Failed to create product")
		return
	}

	response.RespondSuccess(w, http.StatusCreated, "Product created successfully", product)
}

// UpdateProduct updates an existing product.
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.RespondError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		response.RespondError(w, http.StatusBadRequest, "Invalid product data")
		return
	}

	err = h.productService.UpdateProduct(uint(id), updates)
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "Failed to update product")
		return
	}

	response.RespondSuccess(w, http.StatusOK, "Product updated successfully", nil)
}

// DeleteProduct delete a product
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.RespondError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	err = h.productService.DeleteProduct(uint(id))
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	response.RespondSuccess(w, http.StatusOK, "Product deleted successfully", nil)
}

// GetPromotedProducts retrieves a list of promoted products.
func (h *ProductHandler) GetPromotedProducts(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters for pagination
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	products, err := h.productService.GetPromotedProducts(page, limit)
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "Failed to fetch promoted products")
		return
	}

	response.RespondSuccess(w, http.StatusOK, "Promoted products retrieved successfully", products)
}
