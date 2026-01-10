package handlers

import (
	"almoxarifado-backend/internal/services"
	"net/http"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// 1. Decodificar JSON do body
	// 2. Chamar authService.Login()
	// 3. Retornar token ou erro
}
