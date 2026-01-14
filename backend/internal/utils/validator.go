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

// ValidateStruct valida um struct usando as tags de validação
func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		// Formatar mensagens de erro de forma amigável
		var messages []string
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, formatValidationError(err))
		}
		return fmt.Errorf("%s", strings.Join(messages, "; "))
	}
	return nil
}

// formatValidationError formata erros de validação de forma legível
func formatValidationError(err validator.FieldError) string {
	field := strings.ToLower(err.Field())

	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, err.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, err.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", field, err.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", field, err.Param())
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}
