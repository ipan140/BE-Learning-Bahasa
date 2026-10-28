package preference

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitModule(router *echo.Group, db *gorm.DB, authMiddleware echo.MiddlewareFunc) {
	repo := NewRepository(db)
	svc := NewService(repo)
	h := NewHandler(svc)

	prefGroup := router.Group("/preferences", authMiddleware)
	prefGroup.GET("", h.GetPreferences)
	prefGroup.PUT("", h.SavePreferences)
}
