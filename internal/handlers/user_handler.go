package handlers

import (
	"encoding/json"
	"net/http"
	"shopline/internal/models"
	"shopline/internal/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := decodeJSONBody(r, &user); err != nil {
		response.Res
	}
}

func decodeJSONBody(r *http.Request, dst interface{}) error {
	return json.NewDecoder(r.Body).Decode(dst)
}
