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
	return s.repo.Create(item)
}

func (s *ItemService) GetAll() ([]models.Item, error) {
	return s.repo.FindAll()
}

func (s *ItemService) GetByID(id string) (*models.Item, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.repo.FindByID(id)
}

func (s *ItemService) Update(id string, item *models.Item) error {
	if id == "" {
		return errors.New("id is required")
	}

	// Verificar se o item existe
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	item.ID = id
	return s.repo.Update(item)
}

func (s *ItemService) Delete(id string) error {
	if id == "" {
		return errors.New("id is required")
	}

	return s.repo.Delete(id)
}
