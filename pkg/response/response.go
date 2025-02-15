package response

import (
	"encoding/json"
	"net/http"
)

// Response represents the standardized structure of the api responses.
type Response struct {
	Status  string      `json:"status"`         // "success" or "error"
	Message string      `json:"message"`        // Description of the response
	Data    interface{} `json:"data,omitempty"` // Response data (optional)
	Code    int         `json:"code,omitempty"` // Http status code (optional)
}

// RespondSuccess sends a successful JSON response with the given message and data.
func RespondSuccess(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	response := Response{
		Status:  "success",
		Message: message,
		Data:    data,
		Code:    statusCode,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// RespondError sends an error JSON response with the given message.
func RespondError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	response := Response{
		Status:  "error",
		Message: message,
		Code:    statusCode,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
