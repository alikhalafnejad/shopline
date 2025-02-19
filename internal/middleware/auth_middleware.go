package middleware

import (
	"context"
	"net/http"
	"shopline/pkg/auth"
	"shopline/pkg/response"
	"strings"
)

// AuthMiddleware ensures the user is authenticated
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.RespondError(w, http.StatusUnauthorized, "Missing Authorization header")
			return
		}

		// Validate the format of the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			response.RespondError(w, http.StatusUnauthorized, "Invalid Authorization header format")
			return
		}

		// Validate the JWT token and extract user information
		userID, roleName, err := auth.ValidateJWT(tokenString)
		if err != nil {
			response.RespondError(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		// Attach user information to the request context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "userID", userID)
		ctx = context.WithValue(ctx, "roleName", roleName)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
