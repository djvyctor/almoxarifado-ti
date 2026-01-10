package repositories

import (
	"database/sql"

	"almoxarifado-backend/internal/models"
)

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Create(item *models.Item) error {
	query := `
        INSERT INTO items (name, category, quantity)
        VALUES ($1, $2, $3)
        RETURNING id, created_at, updated_at
    `

	err := r.db.QueryRow(query, item.Name, item.Category, item.Quantity).
		Scan(&item.ID, &item.CreatedAt, &item.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}
