package data

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ValidationErrorResponse struct {
	Error       string       `json:"error"`
	FieldErrors []FieldError `json:"field_errors"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewValidationErrorResponse(err error) *ValidationErrorResponse {
	if err != nil {
		var fieldErrors []FieldError
		for _, fErr := range err.(validator.ValidationErrors) {
			fieldErrors = append(fieldErrors, FieldError{
				Field:   fErr.Field(),
				Message: TranslateFieldError(fErr),
			})
		}
		return &ValidationErrorResponse{Error: "Invalid input field(s)", FieldErrors: fieldErrors}
	}
	return nil
}

func TranslateFieldError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%v field is required", err.Field())
	case "email":
		return fmt.Sprintf("%v field must be an valid email", err.Field())
	case "eqfield":
		return fmt.Sprintf("%v field must be equal with %v", err.Field(), err.Param())
	default:
		message := fmt.Sprintf("Invalid input with tag '%v'", err.Tag())
		if err.Param() != "" {
			message += fmt.Sprintf(" and param '%v'", err.Param())
		}
		return message
	}
}
