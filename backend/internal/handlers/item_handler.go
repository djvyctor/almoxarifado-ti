package handlers

import (
	"encoding/json"
	"net/http"

	"almoxarifado-backend/internal/models"
	"almoxarifado-backend/internal/services"
	"almoxarifado-backend/internal/utils"
)

type ItemHandler struct {
	service *services.ItemService
}

func NewItemHandler(service *services.ItemService) *ItemHandler {
	return &ItemHandler{service: service}
}

func (h *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	var item models.Item

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Validar struct
	if err := utils.ValidateStruct(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Create(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func (h *ItemHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// Extrair ID da URL
	id, err := utils.ExtractAndValidateID(r.URL.Path, "/items/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, err := h.service.GetByID(id)
	if err != nil {
		if err.Error() == "item not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ExtractAndValidateID(r.URL.Path, "/items/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Validar struct
	if err := utils.ValidateStruct(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Update(id, &item); err != nil {
		if err.Error() == "item not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ExtractAndValidateID(r.URL.Path, "/items/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(id); err != nil {
		if err.Error() == "item not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
