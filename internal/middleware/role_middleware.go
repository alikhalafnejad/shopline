package middleware

import (
	"net/http"
	"shopline/pkg/response"
)

// RequireRole ensures the user has the specified role.
func RequireRole(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Retrieve the user's role from the context
			userRole := r.Context().Value("roleName").(string)

			// Check if the user has the required role
			if userRole != requiredRole {
				response.RespondError(w, http.StatusForbidden, "You do not have the permission to access this resource")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
