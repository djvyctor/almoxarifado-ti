package repositories

import (
	"almoxarifado-backend/internal/models"
	"database/sql"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `
        INSERT INTO users (name, email, department)
        VALUES ($1, $2, $3)
        RETURNING id, created_at, updated_at
    `

	// Converter strings vazias para NULL
	var email, department sql.NullString
	if user.Email != "" {
		email = sql.NullString{String: user.Email, Valid: true}
	}
	if user.Department != "" {
		department = sql.NullString{String: user.Department, Valid: true}
	}

	err := r.db.QueryRow(query, user.Name, email, department).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	query := `
		SELECT id, name, email, department, created_at, updated_at
		FROM users
		ORDER BY name ASC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		var email, department sql.NullString

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&email,
			&department,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if email.Valid {
			user.Email = email.String
		}
		if department.Valid {
			user.Department = department.String
		}

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) FindByID(id string) (*models.User, error) {
	query := `
		SELECT id, name, email, department, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user models.User
	var email, department sql.NullString

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&email,
		&department,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if email.Valid {
		user.Email = email.String
	}
	if department.Valid {
		user.Department = department.String
	}

	if err == sql.ErrNoRows {
		return nil, errors.New("User not found")
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, department = $3, updated_at = NOW()
		WHERE id = $4
		RETURNING updated_at
	`

	// Converter strings vazias para NULL
	var email, department sql.NullString
	if user.Email != "" {
		email = sql.NullString{String: user.Email, Valid: true}
	}
	if user.Department != "" {
		department = sql.NullString{String: user.Department, Valid: true}
	}

	err := r.db.QueryRow(
		query,
		user.Name,
		email,
		department,
		user.ID,
	).Scan(&user.UpdatedAt)

	if err == sql.ErrNoRows {
		return errors.New("user not found")
	}

	return err
}

func (r *UserRepository) CountLinkedItems(userName string) (int, error) {
	query := `SELECT COUNT(*) FROM items WHERE assigned_to = $1`
	var count int
	err := r.db.QueryRow(query, userName).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *UserRepository) Delete(id string) error {
	// Primeiro, buscar o nome do usuário
	var userName string
	err := r.db.QueryRow(`SELECT name FROM users WHERE id = $1`, id).Scan(&userName)
	if err == sql.ErrNoRows {
		return errors.New("user not found")
	}
	if err != nil {
		return err
	}

	// Deletar todos os itens vinculados ao usuário
	_, err = r.db.Exec(`DELETE FROM items WHERE assigned_to = $1`, userName)
	if err != nil {
		return err
	}

	// Deletar o usuário
	result, err := r.db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
