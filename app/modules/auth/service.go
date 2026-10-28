package auth

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"learning-bahasa-be/common/utils"
)

type Service interface {
	Register(req RegisterRequest) (*AuthResponse, error)
	Login(req LoginRequest) (*AuthResponse, error)
	GoogleLogin(req GoogleLoginRequest) (*AuthResponse, error)
	GetProfile(userID uuid.UUID) (*User, error)
	UpdateProfile(userID uuid.UUID, req UpdateProfileRequest) (*User, error)
}

type service struct {
	repo      Repository
	jwtSecret string
}

func NewService(repo Repository, jwtSecret string) Service {
	return &service{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *service) Register(req RegisterRequest) (*AuthResponse, error) {
	email := strings.TrimSpace(strings.ToLower(req.Email))
	if email == "" || req.Password == "" {
		return nil, errors.New("email dan password wajib diisi")
	}

	if len(req.Password) < 6 {
		return nil, errors.New("password minimal 6 karakter")
	}

	existing, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("email ini sudah terdaftar. Silakan login atau gunakan email lain.")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("gagal mengenkripsi kata sandi")
	}

	displayName := strings.TrimSpace(req.DisplayName)
	if displayName == "" {
		displayName = strings.Split(email, "@")[0]
	}

	newUser := &User{
		Email:        email,
		PasswordHash: hashedPassword,
		DisplayName:  displayName,
	}

	if err := s.repo.Create(newUser); err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(newUser.ID, newUser.Email, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  *newUser,
	}, nil
}

func (s *service) Login(req LoginRequest) (*AuthResponse, error) {
	email := strings.TrimSpace(strings.ToLower(req.Email))
	if email == "" || req.Password == "" {
		return nil, errors.New("email dan password wajib diisi")
	}

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("email atau password salah")
	}

	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("email atau password salah")
	}

	token, err := utils.GenerateToken(user.ID, user.Email, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *service) GoogleLogin(req GoogleLoginRequest) (*AuthResponse, error) {
	email := strings.TrimSpace(strings.ToLower(req.Email))
	if email == "" {
		return nil, errors.New("email Google wajib ada")
	}

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		displayName := req.DisplayName
		if displayName == "" {
			displayName = strings.Split(email, "@")[0]
		}
		user = &User{
			Email:       email,
			DisplayName: displayName,
			AvatarURL:   req.AvatarURL,
		}
		if err := s.repo.Create(user); err != nil {
			return nil, err
		}
	} else {
		if req.AvatarURL != "" && user.AvatarURL != req.AvatarURL {
			user, _ = s.repo.UpdateProfile(user.ID, user.DisplayName, req.AvatarURL)
		}
	}

	token, err := utils.GenerateToken(user.ID, user.Email, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *service) GetProfile(userID uuid.UUID) (*User, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("pengguna tidak ditemukan")
	}
	return user, nil
}

func (s *service) UpdateProfile(userID uuid.UUID, req UpdateProfileRequest) (*User, error) {
	return s.repo.UpdateProfile(userID, req.DisplayName, req.AvatarURL)
}
