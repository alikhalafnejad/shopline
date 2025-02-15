package handlers

import (
	"encoding/json"
	"net/http"
	"shopline/internal/dto"
	"shopline/internal/models"
	"shopline/internal/services"
	"shopline/pkg/response"
	"shopline/pkg/validation"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate the request
	if err := validation.ValidateStruct(req); err != nil {
		response.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.userService.CreateUser(&user)
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "Failed to register user")
		return
	}

	response.RespondSuccess(w, http.StatusCreated, "User registered successfully", user)
}

// Login handles user authentication
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate the request
	if err := validation.ValidateStruct(req); err != nil {
		response.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.userService.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		response.RespondError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}
	response.RespondSuccess(w, http.StatusOK, "Login successful", dto.LoginResponse{Token: token})
}

// GetUser retrieves a user by ID
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		response.RespondError(w, http.StatusNotFound, "User not found")
		return
	}
	response.RespondSuccess(w, http.StatusOK, "User retrieved successfully", user)
}
