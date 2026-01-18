package repositories

import (
	"corvi-backend/internal/models"
	"database/sql"
)

type AdminRepository struct {
	DB *sql.DB
}

func NewAdminRepository(db *sql.DB) *AdminRepository {
	return &AdminRepository{DB: db}
}

func (r *AdminRepository) FindByEmail(email string) (*models.Admin, error) {
	admin := &models.Admin{}

	err := r.DB.QueryRow(
		"SELECT id, email, password_hash, created_at FROM admins WHERE email = $1",
		email,
	).Scan(&admin.ID, &admin.Email, &admin.PasswordHash, &admin.CreatedAt)
	if err != nil {
		return nil, err
	}
	return admin, nil
}
