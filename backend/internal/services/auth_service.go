package services

import (
	"almoxarifado-backend/internal/models"
	"almoxarifado-backend/internal/repositories"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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
	admin, err := s.adminRepo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("credenciais inválidas")
	}

	// 2. Comparar senha
	err = bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("credenciais inválidas")
	}

	// 3. Gerar token
	expiresAt := time.Now().Add(time.Hour * 1).Unix()
	token, err := s.GenerateTokenWithExpiry(admin.ID, expiresAt)
	if err != nil {
		return nil, err
	}

	// 4. Retornar resposta
	return &models.LoginResponse{
		Token:   token,
		Type:    "Bearer",
		Expires: expiresAt,
	}, nil
}

func (s *AuthService) GenerateToken(adminID string) (string, error) {
	expiresAt := time.Now().Add(time.Hour * 1).Unix()
	return s.GenerateTokenWithExpiry(adminID, expiresAt)
}

func (s *AuthService) GenerateTokenWithExpiry(adminID string, expiresAt int64) (string, error) {
	// Criar claims
	claims := jwt.MapClaims{
		"admin_id": adminID,
		"exp":      expiresAt,
		"iat":      time.Now().Unix(),
	}

	// Criar token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assinar token
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) ValidateToken(tokenString string) (string, error) {
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar método de assinatura
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de assinatura inválido")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	// Validar e extrair claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		adminID, ok := claims["admin_id"].(string)
		if !ok {
			return "", errors.New("admin_id não encontrado no token")
		}
		return adminID, nil
	}

	return "", errors.New("token inválido")
}
