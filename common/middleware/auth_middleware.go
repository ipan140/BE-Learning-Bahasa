package middleware

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"learning-bahasa-be/common/utils"
)

func AuthMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return utils.Unauthorized(c, "Header Authorization diperlukan")
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return utils.Unauthorized(c, "Format header Authorization tidak valid (Gunakan: Bearer <token>)")
			}

			claims, err := utils.ValidateToken(parts[1], secret)
			if err != nil {
				return utils.Unauthorized(c, err.Error())
			}

			c.Set("user_id", claims.UserID)
			c.Set("user_email", claims.Email)

			return next(c)
		}
	}
}

func GetUserID(c echo.Context) (uuid.UUID, error) {
	val := c.Get("user_id")
	if val == nil {
		return uuid.Nil, errors.New("user_id tidak ditemukan")
	}
	id, ok := val.(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("tipe user_id tidak valid")
	}
	return id, nil
}

func GetUserEmail(c echo.Context) string {
	val := c.Get("user_email")
	if val == nil {
		return ""
	}
	email, ok := val.(string)
	if !ok {
		return ""
	}
	return email
}
