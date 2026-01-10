package middleware

import (
	"almoxarifado-backend/internal/services"
	"net/http"
)

func AuthMiddleware(authService *services.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 1. Pegar header Authorization
			// 2. Validar formato "Bearer TOKEN"
			// 3. Validar token
			// 4. Adicionar adminID no context
			// 5. Chamar next.ServeHTTP()
		})
	}
}
