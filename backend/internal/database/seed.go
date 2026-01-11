package database

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// SEED ADMIN PADRÃO
func SeedAdminUser(db *sql.DB, defaultPassword string) error {
	// Verificar se já existe um admin
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM admins WHERE username = $1", "admin").Scan(&count)
	if err != nil {
		return err
	}

	// se já existe, não faz nada
	if count > 0 {
		log.Println("admin user already exists")
		return nil
	}

	// Criar hash da senha
	if defaultPassword == "" {
		defaultPassword = "admin123"
		log.Println("usando senha padrão - defina ADMIN_DEFAULT_PASSWORD no .env")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Inserir admin padrão
	_, err = db.Exec(`
		INSERT INTO admins (id, username, password_hash)
		VALUES ($1, $2, $3)
	`, uuid.New().String(), "admin", string(passwordHash))

	if err != nil {
		return err
	}

	log.Println("default admin user created (username: admin, password: " + defaultPassword + ")")
	return nil
}
