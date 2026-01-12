package repositories

import (
	"almoxarifado-backend/internal/models"
	"database/sql"
	"errors"
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

func (r *ItemRepository) FindAll() ([]models.Item, error) {
	query := `
		SELECT id, name, category, quantity, created_at, updated_at
		FROM items
		ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Category,
			&item.Quantity,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *ItemRepository) FindByID(id string) (*models.Item, error) {
	query := `
		SELECT id, name, category, quantity, created_at, updated_at
		FROM items
		WHERE id = $1
	`

	var item models.Item
	err := r.db.QueryRow(query, id).Scan(
		&item.ID,
		&item.Name,
		&item.Category,
		&item.Quantity,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("Item not found")
	}
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *ItemRepository) Update(item *models.Item) error {
	query := `
		UPDATE items
		SET name = $1, category = $2, quantity = $3, updated_at = NOW()
		WHERE id = $4
		RETURNING updated_at
	`

	err := r.db.QueryRow(
		query,
		item.Name,
		item.Category,
		item.Quantity,
		item.ID,
	).Scan(&item.UpdatedAt)

	if err == sql.ErrNoRows {
		return errors.New("item not found")
	}

	return err
}

func (r *ItemRepository) Delete(id string) error {
	query := `DELETE FROM items WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("item not found")
	}

	return nil
}
