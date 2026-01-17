package utils

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

// Vai extrair o ID da URL e validar se é um UUID válido
func ExtractAndValidateID(path, prefix string) (string, error) {
	// Remove o prefixo da URL
	id := strings.TrimPrefix(path, prefix)

	// Remove barras extras
	id = strings.Trim(id, "/")

	// Verifica se está vazio
	if id == "" {
		return "", errors.New("id é obrigatório")
	}

	// Valida se é um UUID válido
	_, err := uuid.Parse(id)
	if err != nil {
		return "", errors.New("id inválido")
	}

	return id, nil
}

// vai remover o prefixo tipo: /items/, remove barras extras, valida o UUID e retorna erro generico
