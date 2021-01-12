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
			message := fmt.Sprintf("Invalid input with tag '%v'", fErr.Tag())
			if fErr.Param() != "" {
				message += fmt.Sprintf(" and param '%v'", fErr.Param())
			}
			fieldErrors = append(fieldErrors, FieldError{
				Field:   fErr.Field(),
				Message: message,
			})
		}
		return &ValidationErrorResponse{Error: "Invalid input field(s)", FieldErrors: fieldErrors}
	}
	return nil
}
