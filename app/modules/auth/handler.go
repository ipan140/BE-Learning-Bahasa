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
