package core_errors

import "net/http"

// AppError defines a custom app error
type AppError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

// Error return the error message of the AppError
func (e *AppError) Error() string {
	return e.Message
}

// Unwrap returns the underlying error of the AppError
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewNotFoundError creates a new AppError with a 404 Not Found status code
func NewNotFoundError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusNotFound,
		Message:    message,
		Err:        err,
	}
}

// NewInternalServerError creates a new AppError with a 500 Internal Server Error status code
func NewInternalServerError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
		Err:        err,
	}
}

// NewDecodeError creates a new AppError with a 400 Bad Request status code
func NewDecodeError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Err:        err,
	}
}
