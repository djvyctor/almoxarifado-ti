package database

import (
	"corvi-backend/config"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func RunMigrations(db *sql.DB, dbURL string) error {
	// Ler e executar 001_create_itens_table.sql
	migration1, err := os.ReadFile("internal/database/migrations/001_create_itens_table.sql")
	if err != nil {
		return fmt.Errorf("erro ao ler migration 001: %v", err)
	}

	if _, err := db.Exec(string(migration1)); err != nil {
		return fmt.Errorf("erro ao executar migration 001: %v", err)
	}

	// Ler e executar 002_create_admins_table.sql
	migration2, err := os.ReadFile("internal/database/migrations/002_create_admins_table.sql")
	if err != nil {
		return fmt.Errorf("erro ao ler migration 002: %v", err)
	}

	if _, err := db.Exec(string(migration2)); err != nil {
		return fmt.Errorf("erro ao executar migration 002: %v", err)
	}

	// Ler e executar 003_create_users_table.sql
	migration3, err := os.ReadFile("internal/database/migrations/003_create_users_table.sql")
	if err != nil {
		return fmt.Errorf("erro ao ler migration 003: %v", err)
	}

	if _, err := db.Exec(string(migration3)); err != nil {
		return fmt.Errorf("erro ao executar migration 003: %v", err)
	}

	return nil
}
