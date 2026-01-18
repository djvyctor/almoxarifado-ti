package handlers

import (
	"almoxarifado-backend/internal/middleware"
	"almoxarifado-backend/internal/models"
	"almoxarifado-backend/internal/services"
	"almoxarifado-backend/internal/utils"
	"encoding/json"
	"log"
	"net/http"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Decodificar JSON do body
	var loginReq models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "dados inválidos"})
		return
	}

	// Validar struct
	if err := utils.ValidateStruct(&loginReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Chamar authService.Login()
	response, err := h.authService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Gerar e enviar token CSRF
	csrfToken, err := middleware.GenerateCSRFToken()
	if err != nil {
		log.Println("Erro ao gerar CSRF token:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "erro interno do servidor"})
		return
	}
	middleware.SetCSRFCookie(w, csrfToken)

	// Retornar token JWT
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
