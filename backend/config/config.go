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
		JWTSecret:            os.Getenv("JWT_SECRET"),
	}

	// Validações de segurança críticas
	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("FATAL: variaveis de ambiente incompletas ou invalidas.")
	}
	if cfg.DBPassword == "" {
		return nil, fmt.Errorf("AVISO: variaveis de ambiente incompletas ou invalidas.")
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
