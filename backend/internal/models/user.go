package models

import "time"

type User struct {
	ID         string    `json:"id"`
	Name       string    `json:"name" validate:"required,min=2,max=100"`
	Email      string    `json:"email" validate:"omitempty,email,max=100"`
	Department string    `json:"department" validate:"omitempty,max=100"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
