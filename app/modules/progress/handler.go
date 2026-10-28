package progress

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

// GetProgress godoc
// @Summary Ambil Progres Huruf yang Dihafal
// @Description Mengambil riwayat huruf yang sudah ditandai hafal (bergaris hijau) per tabel
// @Tags Progress & Quiz
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.Response{data=map[string][]string}
// @Failure 401 {object} utils.Response
// @Router /progress [get]
func (h *Handler) GetProgress(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	data, err := h.service.GetProgress(userID)
	if err != nil {
		return utils.InternalError(c, "Gagal mengambil progres belajar", err)
	}

	return utils.OK(c, data)
}

// SaveProgress godoc
// @Summary Simpan Huruf yang Baru Dihafal
// @Description Menyimpan daftar huruf yang berhasil dihafal pengguna pada suatu tabel
// @Tags Progress & Quiz
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body SaveProgressRequest true "Payload Progres Belajar"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /progress [post]
func (h *Handler) SaveProgress(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	var req SaveProgressRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "Format data tidak valid")
	}

	if err := h.service.SaveProgress(userID, req); err != nil {
		return utils.BadRequest(c, err.Error())
	}

	return utils.Success(c, http.StatusOK, "Progres belajar berhasil disimpan", nil)
}

// SaveQuizScore godoc
// @Summary Simpan Skor Kuis Hafalan & Listening
// @Description Mencatat perolehan skor dan streak kemenangan kuis latihan aksara
// @Tags Progress & Quiz
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body SaveQuizScoreRequest true "Payload Skor Kuis"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /quiz-scores [post]
func (h *Handler) SaveQuizScore(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	var req SaveQuizScoreRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "Format data tidak valid")
	}

	if err := h.service.SaveQuizScore(userID, req); err != nil {
		return utils.BadRequest(c, err.Error())
	}

	return utils.Success(c, http.StatusOK, "Skor kuis berhasil dicatat", nil)
}
