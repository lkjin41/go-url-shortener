package core_errors

import "net/http"

type AppError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewNotFoundError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusNotFound,
		Message:    message,
		Err:        err,
	}
}

func NewInternalServerError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
		Err:        err,
	}
}

func NewDecodeError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Err:        err,
	}
}
