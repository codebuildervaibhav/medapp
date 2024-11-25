package middleware

import (
	"net/http"
	"strings"
)

// AuthMiddleware checks if the request contains a valid token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")

		// Check if the header exists and starts with "Bearer "
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extract the token
		token := authHeader[len("Bearer "):]

		// Validate the token (dummy validation for demonstration)
		if token != "valid-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
