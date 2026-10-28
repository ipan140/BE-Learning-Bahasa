package target

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindByUserID(userID uuid.UUID) ([]UserTarget, error)
	Add(target *UserTarget) error
	Remove(userID uuid.UUID, countryID string) error
	BatchSync(userID uuid.UUID, countryIDs []string) error
	SaveNotes(userID uuid.UUID, countryID string, notes string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindByUserID(userID uuid.UUID) ([]UserTarget, error) {
	if r.db == nil {
		return nil, errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	var targets []UserTarget
	err := r.db.Where("user_id = ?", userID).Find(&targets).Error
	return targets, err
}

func (r *repository) Add(target *UserTarget) error {
	if r.db == nil {
		return errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	if target.ID == uuid.Nil {
		target.ID = uuid.New()
	}
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "country_id"}},
		DoNothing: true,
	}).Create(target).Error
}

func (r *repository) Remove(userID uuid.UUID, countryID string) error {
	if r.db == nil {
		return errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	return r.db.Where("user_id = ? AND country_id = ?", userID, countryID).Delete(&UserTarget{}).Error
}

func (r *repository) BatchSync(userID uuid.UUID, countryIDs []string) error {
	if r.db == nil {
		return errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	if len(countryIDs) == 0 {
		return nil
	}

	targets := make([]UserTarget, len(countryIDs))
	now := time.Now()
	for i, cid := range countryIDs {
		targets[i] = UserTarget{
			ID:        uuid.New(),
			UserID:    userID,
			CountryID: cid,
			CreatedAt: now,
		}
	}

	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "country_id"}},
		DoNothing: true,
	}).Create(&targets).Error
}

func (r *repository) SaveNotes(userID uuid.UUID, countryID string, notes string) error {
	if r.db == nil {
		return errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	target := UserTarget{
		ID:        uuid.New(),
		UserID:    userID,
		CountryID: countryID,
		Notes:     notes,
		CreatedAt: time.Now(),
	}

	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "country_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"notes"}),
	}).Create(&target).Error
}
