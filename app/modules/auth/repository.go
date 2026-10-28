package auth

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindByEmail(email string) (*User, error)
	FindByID(id uuid.UUID) (*User, error)
	Create(user *User) error
	UpdateProfile(id uuid.UUID, displayName, avatarURL string) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindByEmail(email string) (*User, error) {
	if r.db == nil {
		return nil, errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	var user User
	err := r.db.Where("LOWER(email) = LOWER(?)", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByID(id uuid.UUID) (*User, error) {
	if r.db == nil {
		return nil, errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	var user User
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *repository) Create(user *User) error {
	if r.db == nil {
		return errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	return r.db.Create(user).Error
}

func (r *repository) UpdateProfile(id uuid.UUID, displayName, avatarURL string) (*User, error) {
	if r.db == nil {
		return nil, errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	updates := map[string]interface{}{}
	if displayName != "" {
		updates["display_name"] = displayName
	}
	if avatarURL != "" {
		updates["avatar_url"] = avatarURL
	}

	err := r.db.Model(&User{}).Where("id = ?", id).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return r.FindByID(id)
}
