package errs

import "net/http"

type AppError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

func NewHttpNotFoundError(message string) *AppError {
	return &AppError{
		http.StatusNotFound,
		message,
	}
}

func NewInternalServerError(message string) *AppError {
	return &AppError{
		http.StatusInternalServerError,
		message,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		http.StatusUnprocessableEntity,
		message,
	}
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}
