package services

import (
	"almoxarifado-backend/internal/models"
	"almoxarifado-backend/internal/repositories"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	adminRepo *repositories.AdminRepository
	jwtSecret []byte
}

func NewAuthService(adminRepo *repositories.AdminRepository, jwtSecret string) *AuthService {
	return &AuthService{
		adminRepo: adminRepo,
		jwtSecret: []byte(jwtSecret),
	}
}

func (s *AuthService) Login(email, password string) (*models.LoginResponse, error) {
	// Buscar admin
	admin, err := s.adminRepo.FindByEmail(email)

	// Sempre verificar senha mesmo que o admin não exista
	var passwordHash string
	if err != nil {
		// Usando HASH fake
		passwordHash = "$2a$10$XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	} else {
		passwordHash = admin.PasswordHash
	}

	// Verificar senha
	bytesErr := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))

	// Verifica se houve erro na busca ou na senha
	if err != nil || bytesErr != nil {
		return nil, errors.New("email ou senha inválidos")
	}

	// Gerar token
	expiresAT := time.Now().Add(time.Hour * 1).Unix()
	token, err := s.GenerateTokenWithExpiry(admin.ID, expiresAT)
	if err != nil {
		return nil, err
	}

	// Retornar resposta
	return &models.LoginResponse{
		Token:   token,
		Type:    "Bearer",
		Expires: expiresAT,
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
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) ValidateToken(tokenString string) (string, error) {
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar método de assinatura
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("método de assinatura inválido")
		}
		return s.jwtSecret, nil
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
