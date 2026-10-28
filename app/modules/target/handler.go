package target

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

func (h *Handler) GetTargets(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	data, err := h.service.GetTargets(userID)
	if err != nil {
		return utils.InternalError(c, "Gagal mengambil target negara", err)
	}

	return utils.OK(c, data)
}

func (h *Handler) AddTarget(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	var req AddTargetRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "Format data tidak valid")
	}

	if err := h.service.AddTarget(userID, req.CountryID); err != nil {
		return utils.BadRequest(c, err.Error())
	}

	return utils.Success(c, http.StatusOK, "Target negara berhasil ditambahkan", nil)
}

func (h *Handler) RemoveTarget(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	countryID := c.Param("countryId")
	if err := h.service.RemoveTarget(userID, countryID); err != nil {
		return utils.BadRequest(c, err.Error())
	}

	return utils.Success(c, http.StatusOK, "Target negara berhasil dihapus", nil)
}

func (h *Handler) BatchSync(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	var req BatchSyncRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "Format data tidak valid")
	}

	if err := h.service.BatchSync(userID, req.CountryIDs); err != nil {
		return utils.InternalError(c, "Gagal sinkronisasi target", err)
	}

	return utils.Success(c, http.StatusOK, "Sinkronisasi target negara berhasil", nil)
}

func (h *Handler) SaveNotes(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	countryID := c.Param("countryId")
	var req SaveNotesRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "Format data tidak valid")
	}

	if err := h.service.SaveNotes(userID, countryID, req.Notes); err != nil {
		return utils.BadRequest(c, err.Error())
	}

	return utils.Success(c, http.StatusOK, "Catatan target negara berhasil disimpan", nil)
}
