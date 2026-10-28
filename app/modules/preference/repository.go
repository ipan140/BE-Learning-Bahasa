package preference

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindByUserID(userID uuid.UUID) (*UserPreference, error)
	Upsert(pref *UserPreference) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindByUserID(userID uuid.UUID) (*UserPreference, error) {
	if r.db == nil {
		return nil, errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	var pref UserPreference
	err := r.db.Where("user_id = ?", userID).First(&pref).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &pref, nil
}

func (r *repository) Upsert(pref *UserPreference) error {
	if r.db == nil {
		return errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	pref.UpdatedAt = time.Now()
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"audio_speed", "audio_enabled", "theme", "updated_at"}),
	}).Create(pref).Error
}
