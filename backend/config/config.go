package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort              string
	DBHost               string
	DBPort               string
	DBUser               string
	DBPassword           string
	DBName               string
	AdminDefaultEmail    string
	AdminDefaultPassword string
	JWTSecret            string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using system environment variables")
	}

	cfg := &Config{
		AppPort:              getEnv("APP_PORT", ""),
		DBHost:               getEnv("DB_HOST", ""),
		DBPort:               getEnv("DB_PORT", ""),
		DBUser:               getEnv("DB_USER", ""),
		DBPassword:           getEnv("DB_PASSWORD", ""),
		DBName:               getEnv("DB_NAME", ""),
		AdminDefaultEmail:    getEnv("ADMIN_DEFAULT_EMAIL", ""),
		AdminDefaultPassword: getEnv("ADMIN_DEFAULT_PASSWORD", ""),
		JWTSecret:            getEnv("JWT_SECRET", ""),
	}

	// Vou implementar uma validação simples, se faltar algo, o app não inicia
	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("FATAL: A variavel de ambiente JWT_SECRET é obrigatória.")
	}
	if cfg.DBPassword == "" {
		return nil, fmt.Errorf("AVISO: Rodando sem senha de banco de dados meu brodi!!!")
	}
	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func (c *Config) DatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
}
