package target

import (
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
	var targets []UserTarget
	err := r.db.Where("user_id = ?", userID).Find(&targets).Error
	return targets, err
}

func (r *repository) Add(target *UserTarget) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "country_id"}},
		DoNothing: true,
	}).Create(target).Error
}

func (r *repository) Remove(userID uuid.UUID, countryID string) error {
	return r.db.Where("user_id = ? AND country_id = ?", userID, countryID).Delete(&UserTarget{}).Error
}

func (r *repository) BatchSync(userID uuid.UUID, countryIDs []string) error {
	if len(countryIDs) == 0 {
		return nil
	}

	targets := make([]UserTarget, len(countryIDs))
	now := time.Now()
	for i, cid := range countryIDs {
		targets[i] = UserTarget{
			UserID:    userID,
			CountryID: cid,
			CreatedAt: now,
			UpdatedAt: now,
		}
	}

	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "country_id"}},
		DoNothing: true,
	}).Create(&targets).Error
}

func (r *repository) SaveNotes(userID uuid.UUID, countryID string, notes string) error {
	target := UserTarget{
		UserID:    userID,
		CountryID: countryID,
		Notes:     notes,
		UpdatedAt: time.Now(),
	}

	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "country_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"notes", "updated_at"}),
	}).Create(&target).Error
}
