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

// GetPreferences godoc
// @Summary Ambil preferensi pengguna
// @Description Mengambil preferensi pengguna yang sedang login (audio speed, audio enabled, theme)
// @Tags preferences
// @Accept json
// @Produce json
// @Success 200 {object} UserPreference
// @Failure 401 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /preferences [get]
// @Security BearerAuth
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

// SavePreferences godoc
// @Summary Simpan preferensi pengguna
// @Description Menyimpan/memperbarui preferensi audio speed, audio enabled, dan tema
// @Tags preferences
// @Accept json
// @Produce json
// @Param request body SavePreferenceRequest true "Preference payload"
// @Success 200 {object} UserPreference
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /preferences [put]
// @Security BearerAuth
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
