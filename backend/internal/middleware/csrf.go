package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strings"
)

const (
	csrfCookieName = "csrf-token"
	csrfHeaderName = "X-CSRF-Token"
)

// Gera um token CSRF aleatório
func GenerateCSRFToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(token), nil
}

// Define o cookie CSRF na resposta
func SetCSRFCookie(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:     csrfCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   3600, // 1 hora
	}
	http.SetCookie(w, cookie)
}

// Valida o token CSRF
func CSRFMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Permite métodos seguros sem validação CSRF
		if r.Method == http.MethodGet || r.Method == http.MethodHead || r.Method == http.MethodOptions {
			next.ServeHTTP(w, r)
			return
		}

		// Pegar token do cookie
		cookie, err := r.Cookie(csrfCookieName)
		if err != nil || cookie.Value == "" {
			http.Error(w, "CSRF token ausente no cookie", http.StatusForbidden)
			return
		}

		// Pegar token do header
		headerToken := r.Header.Get(csrfHeaderName)
		if headerToken == "" {
			http.Error(w, "CSRF token ausente no header", http.StatusForbidden)
			return
		}

		// Valida se os tokens são iguais
		if !strings.EqualFold(cookie.Value, headerToken) {
			http.Error(w, "CSRF token inválido", http.StatusForbidden)
			return
		}

		// Token válido
		next.ServeHTTP(w, r)
	})
}
