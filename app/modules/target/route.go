package target

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitModule(router *echo.Group, db *gorm.DB, authMiddleware echo.MiddlewareFunc) {
	repo := NewRepository(db)
	svc := NewService(repo)
	h := NewHandler(svc)

	targetGroup := router.Group("/targets", authMiddleware)
	targetGroup.GET("", h.GetTargets)
	targetGroup.POST("", h.AddTarget)
	targetGroup.DELETE("/:countryId", h.RemoveTarget)
	targetGroup.POST("/batch", h.BatchSync)
	targetGroup.PUT("/:countryId/notes", h.SaveNotes)
}
