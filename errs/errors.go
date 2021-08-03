package errs

import "net/http"

type AppErr struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func NewNotFoundError(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func (e AppErr) AsMessage() *AppErr {
	return &AppErr{
		Message: e.Message,
	}
}

func NewUnexpectedError(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewValidationError(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}
