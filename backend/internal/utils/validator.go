package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct valida um struct usando as tags
func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		// Formatar mensagens de erro
		var messages []string
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, formatValidationError(err))
		}
		return fmt.Errorf("%s", strings.Join(messages, "; "))
	}
	return nil
}

// formata erros de validação de forma legível
func formatValidationError(err validator.FieldError) string {
	field := strings.ToLower(err.Field())

	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s é obrigatório", field)
	case "email":
		return fmt.Sprintf("%s deve ser um e-mail válido", field)
	case "min":
		return fmt.Sprintf("%s deve ter no mínimo %s caracteres", field, err.Param())
	case "max":
		return fmt.Sprintf("%s deve ter no máximo %s caracteres", field, err.Param())
	case "gte":
		return fmt.Sprintf("%s deve ser maior ou igual a %s", field, err.Param())
	case "lte":
		return fmt.Sprintf("%s deve ser menor ou igual a %s", field, err.Param())
	default:
		return fmt.Sprintf("%s é inválido", field)
	}
}
