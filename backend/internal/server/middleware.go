package server

import (
	"backend/internal/auth-service"
	"net/http"
)

func AuthorizationMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Before calling the next handler

		if r.URL.Path == "/register" || r.URL.Path == "/login" || r.URL.Path == "/" {
			next.ServeHTTP(w, r) // Call the next handler
			return
		}
		jwt := r.Header.Get("Authorization")

		_, err := auth.ValidateToken(jwt)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r) // Call the next handler

	})
}
