package phrase

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"learning-bahasa-be/common/middleware"
	"learning-bahasa-be/common/utils"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetPhrases(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	data, err := h.service.GetPhrases(userID)
	if err != nil {
		return utils.InternalError(c, "Gagal mengambil frasa kustom", err)
	}

	return utils.OK(c, data)
}

func (h *Handler) SavePhrases(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	var req SavePhrasesRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "Format data tidak valid")
	}

	if err := h.service.SavePhrases(userID, req); err != nil {
		return utils.BadRequest(c, err.Error())
	}

	return utils.Success(c, http.StatusOK, "Frasa kustom berhasil disimpan", nil)
}

func (h *Handler) DeletePhrases(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	langID := c.Param("languageId")
	if err := h.service.DeletePhrases(userID, langID); err != nil {
		return utils.BadRequest(c, err.Error())
	}

	return utils.Success(c, http.StatusOK, "Frasa kustom berhasil dihapus", nil)
}

func (h *Handler) GetCategories(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	cats, err := h.service.GetCategories(userID)
	if err != nil {
		return utils.InternalError(c, "Gagal mengambil kategori frasa", err)
	}

	return utils.OK(c, cats)
}

func (h *Handler) SaveCategory(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	var req SaveCategoryRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "Format data tidak valid")
	}

	if err := h.service.SaveCategory(userID, req.Name); err != nil {
		return utils.BadRequest(c, err.Error())
	}

	return utils.Success(c, http.StatusOK, "Kategori frasa berhasil ditambahkan", nil)
}

func (h *Handler) DeleteCategory(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	name := c.Param("name")
	if err := h.service.DeleteCategory(userID, name); err != nil {
		return utils.BadRequest(c, err.Error())
	}

	return utils.Success(c, http.StatusOK, "Kategori frasa berhasil dihapus", nil)
}
