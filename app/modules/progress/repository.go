package progress

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindByUserID(userID uuid.UUID) ([]UserStudyProgress, error)
	Upsert(progress *UserStudyProgress) error
	SaveScore(userID uuid.UUID, tableID string, score int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindByUserID(userID uuid.UUID) ([]UserStudyProgress, error) {
	var list []UserStudyProgress
	err := r.db.Where("user_id = ?", userID).Find(&list).Error
	return list, err
}

func (r *repository) Upsert(progress *UserStudyProgress) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "table_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"memorized_chars", "score", "last_practiced_at"}),
	}).Create(progress).Error
}

func (r *repository) SaveScore(userID uuid.UUID, tableID string, score int) error {
	progress := UserStudyProgress{
		UserID:          userID,
		TableID:         tableID,
		Score:           score,
		LastPracticedAt: time.Now(),
	}

	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "table_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"score", "last_practiced_at"}),
	}).Create(&progress).Error
}
