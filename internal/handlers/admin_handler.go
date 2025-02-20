package handlers

import (
	"github.com/go-chi/chi/v5"
	"shopline/internal/services"
)

type AdminHandler struct {
	adminService *services.AdminService
}

func NewAdminHandler(adminService *services.AdminService) *AdminHandler {
	return &AdminHandler{adminService: adminService}
}

// RegisterRoutes register user related-routes
func (h *AdminHandler) RegisterRoutes(r chi.Router) {

}
