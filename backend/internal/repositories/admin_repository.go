package repositories

import (
	"almoxarifado-backend/internal/models"
	"database/sql"
)

type AdminRepository struct {
	DB *sql.DB
}

func NewAdminRepository(db *sql.DB) *AdminRepository {
	return &AdminRepository{DB: db}
}

func (r *AdminRepository) FindByUsername(username string) (*models.Admin, error) {
	admin := &models.Admin{}

	err := r.DB.QueryRow(
		"SELECT id, username, password_hash, created_at FROM admins WHERE username = $1",
		username,
	).Scan(&admin.ID, &admin.Username, &admin.PasswordHash, &admin.CreatedAt)
	if err != nil {
		return nil, err
	}
	return admin, nil
}
