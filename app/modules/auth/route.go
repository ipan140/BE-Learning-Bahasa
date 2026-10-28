package auth

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitModule(router *echo.Group, db *gorm.DB, jwtSecret string, authMiddleware echo.MiddlewareFunc) {
	repo := NewRepository(db)
	svc := NewService(repo, jwtSecret)
	h := NewHandler(svc)

	authGroup := router.Group("/auth")
	authGroup.POST("/register", h.Register)
	authGroup.POST("/login", h.Login)
	authGroup.POST("/google", h.GoogleLogin)

	// Protected routes
	authGroup.GET("/me", h.GetMe, authMiddleware)
	authGroup.PUT("/profile", h.UpdateProfile, authMiddleware)
}
