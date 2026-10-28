package phrase

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitModule(router *echo.Group, db *gorm.DB, authMiddleware echo.MiddlewareFunc) {
	repo := NewRepository(db)
	svc := NewService(repo)
	h := NewHandler(svc)

	phraseGroup := router.Group("/phrases", authMiddleware)
	phraseGroup.GET("", h.GetPhrases)
	phraseGroup.POST("", h.SavePhrases)
	phraseGroup.DELETE("/:languageId", h.DeletePhrases)

	catGroup := router.Group("/phrase-categories", authMiddleware)
	catGroup.GET("", h.GetCategories)
	catGroup.POST("", h.SaveCategory)
	catGroup.DELETE("/:name", h.DeleteCategory)
}
