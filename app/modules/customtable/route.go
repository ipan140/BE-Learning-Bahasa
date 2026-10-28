package customtable

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitModule(router *echo.Group, db *gorm.DB, authMiddleware echo.MiddlewareFunc) {
	repo := NewRepository(db)
	svc := NewService(repo)
	h := NewHandler(svc)

	tableGroup := router.Group("/custom-tables", authMiddleware)
	tableGroup.GET("", h.GetTables)
	tableGroup.POST("", h.SaveTable)
	tableGroup.PUT("/:id", h.SaveTable)
	tableGroup.DELETE("/:id", h.DeleteTable)
}
