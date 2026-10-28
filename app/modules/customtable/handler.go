package customtable

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

func (h *Handler) GetTables(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	tables, err := h.service.GetTables(userID)
	if err != nil {
		return utils.InternalError(c, "Gagal mengambil tabel kustom", err)
	}

	return utils.OK(c, tables)
}

func (h *Handler) SaveTable(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	var req UpsertTableRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "Format data tidak valid")
	}

	if idParam := c.Param("id"); idParam != "" && req.ID == "" {
		req.ID = idParam
	}

	table, err := h.service.SaveTable(userID, req)
	if err != nil {
		return utils.BadRequest(c, err.Error())
	}

	return utils.Success(c, http.StatusOK, "Tabel kustom berhasil disimpan", table)
}

func (h *Handler) DeleteTable(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	id := c.Param("id")
	if err := h.service.DeleteTable(userID, id); err != nil {
		return utils.BadRequest(c, err.Error())
	}

	return utils.Success(c, http.StatusOK, "Tabel kustom berhasil dihapus", nil)
}
