package preference

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

func (h *Handler) GetPreferences(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	data, err := h.service.GetPreferences(userID)
	if err != nil {
		return utils.InternalError(c, "Gagal mengambil preferensi pengguna", err)
	}

	return utils.OK(c, data)
}

func (h *Handler) SavePreferences(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	var req SavePreferenceRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "Format data tidak valid")
	}

	pref, err := h.service.SavePreferences(userID, req)
	if err != nil {
		return utils.InternalError(c, "Gagal menyimpan preferensi", err)
	}

	return utils.Success(c, http.StatusOK, "Preferensi pengguna berhasil disimpan", pref)
}
