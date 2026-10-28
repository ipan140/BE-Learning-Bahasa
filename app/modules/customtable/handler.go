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

// GetTables godoc
// @Summary Ambil Seluruh Tabel Aksara Kustom
// @Description Mengambil semua tabel aksara buatan pengguna beserta baris/grid huruf
// @Tags Custom Tables
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.Response{data=[]UserCustomTable}
// @Failure 401 {object} utils.Response
// @Router /custom-tables [get]
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

// SaveTable godoc
// @Summary Buat atau Perbarui Tabel Aksara Kustom
// @Description Menyimpan tabel aksara baru atau memperbarui susunan kotak huruf (Mode Edit aktif)
// @Tags Custom Tables
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string false "ID Tabel Kustom (opsional jika update)"
// @Param request body UpsertTableRequest true "Payload Tabel Aksara"
// @Success 200 {object} utils.Response{data=UserCustomTable}
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /custom-tables [post]
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

// DeleteTable godoc
// @Summary Hapus Tabel Aksara Kustom
// @Description Menghapus tabel aksara buatan pengguna berdasarkan ID
// @Tags Custom Tables
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID Tabel Kustom (contoh: custom-173000000)"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /custom-tables/{id} [delete]
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
