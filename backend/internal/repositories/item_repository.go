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
	// Futuramente implementar a lógica de inserção no banco de dados
	return nil
}
