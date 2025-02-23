package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"shopline/internal/services"
	"shopline/pkg/response"
	"strconv"
)

type CommentHandler struct {
	commentService *services.CommentService
}

func NewCommentHandler(commentService *services.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

// RegisterRoutes registers comment-related routes.
func (h *CommentHandler) RegisterRoutes(r chi.Router) {
	r.Get("/products/{id}/comments", h.GetPublishedCommentsByProduct)
}

// Middlewares returns middleware specific to this handler (if needed).
func (h *CommentHandler) Middlewares() []func(http.Handler) http.Handler {
	return nil // No additional middleware for this handler
}

// GetPublishedCommentsByProduct retrieves paginated published comments for a specific product
func (h *CommentHandler) GetPublishedCommentsByProduct(w http.ResponseWriter, r *http.Request) {
	productIDStr := chi.URLParam(r, "id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		response.RespondError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	// Parse Paginate Parameters
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	// Fetch published comments with pagination
	comments, err := h.commentService.GetPublishedCommentsByProduct(uint(productID), page, limit)
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "Failed to fetch published comments")
		return
	}

	response.RespondSuccess(w, http.StatusOK, "Published comments retrieved successfully", comments)
}
