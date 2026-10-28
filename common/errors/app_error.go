package errors

import "net/http"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code int, message string, detail ...string) *AppError {
	det := ""
	if len(detail) > 0 {
		det = detail[0]
	}
	return &AppError{
		Code:    code,
		Message: message,
		Detail:  det,
	}
}

func BadRequest(message string, detail ...string) *AppError {
	return New(http.StatusBadRequest, message, detail...)
}

func Unauthorized(message string, detail ...string) *AppError {
	return New(http.StatusUnauthorized, message, detail...)
}

func NotFound(message string, detail ...string) *AppError {
	return New(http.StatusNotFound, message, detail...)
}

func Internal(message string, detail ...string) *AppError {
	return New(http.StatusInternalServerError, message, detail...)
}
