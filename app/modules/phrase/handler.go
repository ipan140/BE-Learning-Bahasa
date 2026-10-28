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

// GetPhrases godoc
// @Summary Ambil Seluruh Frasa Kustom
// @Description Mengambil kumpulan frasa harian buatan pengguna yang dikelompokkan per bahasa
// @Tags Survival Phrases
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.Response{data=map[string]interface{}}
// @Failure 401 {object} utils.Response
// @Router /phrases [get]
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

// SavePhrases godoc
// @Summary Simpan Kumpulan Frasa Kustom
// @Description Menyimpan atau memperbarui kumpulan frasa dan kategori untuk satu bahasa
// @Tags Survival Phrases
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body SavePhrasesRequest true "Payload Frasa Bahasa"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /phrases [post]
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

// DeletePhrases godoc
// @Summary Hapus Frasa Kustom per Bahasa
// @Description Menghapus seluruh frasa kustom pada bahasa tertentu
// @Tags Survival Phrases
// @Produce json
// @Security BearerAuth
// @Param languageId path string true "ID Bahasa (contoh: japanese)"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /phrases/{languageId} [delete]
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

// GetCategories godoc
// @Summary Ambil Daftar Kategori Frasa
// @Description Mengambil semua kategori frasa buatan pengguna
// @Tags Survival Phrases
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.Response{data=[]UserPhraseCategory}
// @Failure 401 {object} utils.Response
// @Router /phrase-categories [get]
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

// SaveCategory godoc
// @Summary Tambah Kategori Frasa Baru
// @Description Membuat satu nama kategori frasa baru beserta icon pilihannya
// @Tags Survival Phrases
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body SaveCategoryRequest true "Payload Kategori"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /phrase-categories [post]
func (h *Handler) SaveCategory(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	var req SaveCategoryRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "Format data tidak valid")
	}

	if err := h.service.SaveCategory(userID, req); err != nil {
		return utils.BadRequest(c, err.Error())
	}

	return utils.Success(c, http.StatusOK, "Kategori frasa berhasil ditambahkan", nil)
}

// DeleteCategory godoc
// @Summary Hapus Kategori Frasa
// @Description Menghapus kategori frasa kustom berdasarkan nama
// @Tags Survival Phrases
// @Produce json
// @Security BearerAuth
// @Param name path string true "Nama Kategori (contoh: Belanja)"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /phrase-categories/{name} [delete]
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
