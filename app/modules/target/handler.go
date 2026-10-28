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

// GetTargets godoc
// @Summary Ambil Daftar Negara Target & Catatan
// @Description Mengambil semua negara berbintang dan catatan motivasi milik user
// @Tags Targets
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.Response{data=TargetsResponse}
// @Failure 401 {object} utils.Response
// @Router /targets [get]
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

// AddTarget godoc
// @Summary Tandai Negara Target Baru (Bintang)
// @Description Menambahkan satu kode negara ke dalam daftar target prioritas belajar
// @Tags Targets
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body AddTargetRequest true "Payload Tambah Target"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /targets [post]
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

// RemoveTarget godoc
// @Summary Hapus Bintang dari Negara Target
// @Description Menghapus satu negara dari daftar target belajar pengguna
// @Tags Targets
// @Produce json
// @Security BearerAuth
// @Param countryId path string true "ID/Kode Negara (contoh: JP)"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /targets/{countryId} [delete]
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

// BatchSync godoc
// @Summary Sinkronisasi Massal Target Negara (Batch)
// @Description Menyimpan daftar negara berbintang dari memori browser lokal saat pertama kali login
// @Tags Targets
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body BatchSyncRequest true "Payload Daftar Kode Negara"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /targets/batch [post]
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

// SaveNotes godoc
// @Summary Simpan Catatan Motivasi / Rencana Negara
// @Description Menyimpan catatan impian, rencana kunjungan, atau target hafalan pada negara tertentu
// @Tags Targets
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param countryId path string true "Kode Negara (contoh: JP)"
// @Param request body SaveNotesRequest true "Payload Catatan"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /targets/{countryId}/notes [put]
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
