package progress

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitModule(router *echo.Group, db *gorm.DB, authMiddleware echo.MiddlewareFunc) {
	repo := NewRepository(db)
	svc := NewService(repo)
	h := NewHandler(svc)

	progGroup := router.Group("/progress", authMiddleware)
	progGroup.GET("", h.GetProgress)
	progGroup.POST("", h.SaveProgress)

	router.POST("/quiz-scores", h.SaveQuizScore, authMiddleware)
}
