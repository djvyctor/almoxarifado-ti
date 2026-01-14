package models

import "time"

type Item struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" validate:"required,min=2,max=100"`
	Category  string    `json:"category" validate:"required,min=2,max=50"`
	Quantity  int       `json:"quantity" validate:"gte=0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
