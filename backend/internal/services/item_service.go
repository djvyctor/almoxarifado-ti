package services

import (
	"errors"

	"almoxarifado-backend/internal/models"
	"almoxarifado-backend/internal/repositories"
)

type ItemService struct {
	repo *repositories.ItemRepository
}

func NewItemService(repo *repositories.ItemRepository) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) Create(item *models.Item) error {
	if item.Name == "" {
		return errors.New("name is required")
	}

	if item.Category == "" {
		return errors.New("category is required")
	}

	if item.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}

	return s.repo.Create(item)
}
