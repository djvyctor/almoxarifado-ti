package middleware

import (
	"context"
	"corvi-backend/internal/services"
	"encoding/json"
	"net/http"
	"strings"
)

func AuthMiddleware(authService *services.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 1. Pegar header Authorization
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "token não fornecido"})
				return
			}

			// 2. Validar formato "Bearer TOKEN"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "formato de token inválido"})
				return
			}

			tokenString := parts[1]

			// 3. Validar token
			adminID, err := authService.ValidateToken(tokenString)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "token inválido"})
				return
			}

			// 4. Adicionar adminID no context
			ctx := context.WithValue(r.Context(), "admin_id", adminID)

			// 5. Chamar next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
