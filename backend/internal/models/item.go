package models

import "time"

type Item struct {
	ID          string    `json:"id"`
	Name        string    `json:"name" validate:"required,min=2,max=100"`
	Category    string    `json:"category" validate:"required,min=2,max=50"`
	Patrimony   string    `json:"patrimony" validate:"omitempty,max=50"`
	AssignedTo  string    `json:"assigned_to" validate:"omitempty,max=100"`
	Description string    `json:"description" validate:"omitempty,max=500"`
	Value       *float64  `json:"value" validate:"omitempty,min=0"`
	Quantity    int       `json:"quantity" validate:"gte=0"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
