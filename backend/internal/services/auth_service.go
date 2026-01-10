package services

import (
	"almoxarifado-backend/internal/models"
	"almoxarifado-backend/internal/repositories"
)

var jwtSecret = []byte("your-secret-key-change-in-production") // Mover para ENV depois

type AuthService struct {
	adminRepo *repositories.AdminRepository
}

func NewAuthService(adminRepo *repositories.AdminRepository) *AuthService {
	return &AuthService{adminRepo: adminRepo}
}

func (s *AuthService) Login(username, password string) (*models.LoginResponse, error) {
	// 1. Buscar admin
	// 2. Comparar senha
	// 3. Gerar token
	// 4. Retornar resposta
}

func (s *AuthService) GenerateToken(adminID string) (string, error) {
	// Criar claims com expiração de 24h
	// Assinar token
}

func (s *AuthService) ValidateToken(tokenString string) (string, error) {
	// Parse e valida token
	// Retorna adminID se válido
}
