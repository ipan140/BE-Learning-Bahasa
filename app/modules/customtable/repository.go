package customtable

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindByUserID(userID uuid.UUID) ([]UserCustomTable, error)
	Upsert(table *UserCustomTable) error
	Delete(userID uuid.UUID, id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindByUserID(userID uuid.UUID) ([]UserCustomTable, error) {
	var tables []UserCustomTable
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&tables).Error
	return tables, err
}

func (r *repository) Upsert(table *UserCustomTable) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "voice_lang", "description", "rows", "updated_at"}),
	}).Create(table).Error
}

func (r *repository) Delete(userID uuid.UUID, id string) error {
	return r.db.Where("user_id = ? AND id = ?", userID, id).Delete(&UserCustomTable{}).Error
}
