package auth

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

// Register godoc
// @Summary Daftar Akun Baru
// @Description Mendaftarkan pengguna baru dengan email dan password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Payload Pendaftaran"
// @Success 201 {object} utils.Response{data=AuthResponse}
// @Failure 400 {object} utils.Response
// @Router /auth/register [post]
func (h *Handler) Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "Format data tidak valid")
	}

	res, err := h.service.Register(req)
	if err != nil {
		return utils.BadRequest(c, err.Error())
	}

	return utils.Success(c, http.StatusCreated, "Akun berhasil dibuat", res)
}

// Login godoc
// @Summary Masuk ke Akun
// @Description Otentikasi pengguna menggunakan email dan password, mengembalikan token JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Payload Login"
// @Success 200 {object} utils.Response{data=AuthResponse}
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /auth/login [post]
func (h *Handler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "Format data tidak valid")
	}

	res, err := h.service.Login(req)
	if err != nil {
		return utils.Unauthorized(c, err.Error())
	}

	return utils.OK(c, res)
}

// GoogleLogin godoc
// @Summary Masuk dengan Google OAuth
// @Description Login atau registrasi otomatis menggunakan akun Google
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body GoogleLoginRequest true "Payload Google Auth"
// @Success 200 {object} utils.Response{data=AuthResponse}
// @Failure 400 {object} utils.Response
// @Router /auth/google [post]
func (h *Handler) GoogleLogin(c echo.Context) error {
	var req GoogleLoginRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "Format data Google tidak valid")
	}

	res, err := h.service.GoogleLogin(req)
	if err != nil {
		return utils.BadRequest(c, err.Error())
	}

	return utils.OK(c, res)
}

// GetMe godoc
// @Summary Ambil Data Pengguna Saat Ini
// @Description Mendapatkan informasi profil pengguna yang sedang login
// @Tags Auth
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.Response{data=User}
// @Failure 401 {object} utils.Response
// @Router /auth/me [get]
func (h *Handler) GetMe(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	user, err := h.service.GetProfile(userID)
	if err != nil {
		return utils.NotFound(c, err.Error())
	}

	return utils.OK(c, user)
}

// UpdateProfile godoc
// @Summary Perbarui Profil Pengguna
// @Description Mengubah nama tampilan dan foto profil avatar pengguna
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body UpdateProfileRequest true "Payload Update Profil"
// @Success 200 {object} utils.Response{data=User}
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /auth/profile [put]
func (h *Handler) UpdateProfile(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "Unauthorized")
	}

	var req UpdateProfileRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "Format data tidak valid")
	}

	user, err := h.service.UpdateProfile(userID, req)
	if err != nil {
		return utils.InternalError(c, "Gagal memperbarui profil", err)
	}

	return utils.Success(c, http.StatusOK, "Profil berhasil diperbarui", user)
}
