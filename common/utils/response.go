package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Success(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func OK(c echo.Context, data interface{}) error {
	return Success(c, http.StatusOK, "Berhasil", data)
}

func Error(c echo.Context, statusCode int, message string, errDetail string) error {
	return c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Error:   errDetail,
	})
}

func BadRequest(c echo.Context, message string) error {
	return Error(c, http.StatusBadRequest, message, "")
}

func Unauthorized(c echo.Context, message string) error {
	return Error(c, http.StatusUnauthorized, message, "")
}

func NotFound(c echo.Context, message string) error {
	return Error(c, http.StatusNotFound, message, "")
}

func InternalError(c echo.Context, message string, err error) error {
	detail := ""
	if err != nil {
		detail = err.Error()
	}
	return Error(c, http.StatusInternalServerError, message, detail)
}
