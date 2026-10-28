package progress

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindByUserID(userID uuid.UUID) ([]UserStudyProgress, error)
	Upsert(progress *UserStudyProgress) error
	SaveScore(userID uuid.UUID, tableID string, score, streak int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindByUserID(userID uuid.UUID) ([]UserStudyProgress, error) {
	if r.db == nil {
		return nil, errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	var list []UserStudyProgress
	err := r.db.Where("user_id = ?", userID).Find(&list).Error
	return list, err
}

func (r *repository) Upsert(progress *UserStudyProgress) error {
	if r.db == nil {
		return errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	if progress.ID == uuid.Nil {
		progress.ID = uuid.New()
	}
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "table_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"memorized_chars", "score", "streak_count", "last_practiced_at", "metadata"}),
	}).Create(progress).Error
}

func (r *repository) SaveScore(userID uuid.UUID, tableID string, score, streak int) error {
	if r.db == nil {
		return errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	progress := UserStudyProgress{
		ID:              uuid.New(),
		UserID:          userID,
		TableID:         tableID,
		Score:           score,
		StreakCount:     streak,
		LastPracticedAt: time.Now(),
	}

	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "table_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"score", "streak_count", "last_practiced_at"}),
	}).Create(&progress).Error
}
