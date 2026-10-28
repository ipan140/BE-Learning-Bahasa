package phrase

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindPhrasesByUserID(userID uuid.UUID) ([]UserCustomPhrase, error)
	UpsertPhrases(phrase *UserCustomPhrase) error
	DeletePhrases(userID uuid.UUID, languageID string) error

	FindCategoriesByUserID(userID uuid.UUID) ([]UserPhraseCategory, error)
	AddCategory(category *UserPhraseCategory) error
	DeleteCategory(userID uuid.UUID, name string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindPhrasesByUserID(userID uuid.UUID) ([]UserCustomPhrase, error) {
	if r.db == nil {
		return nil, errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	var list []UserCustomPhrase
	err := r.db.Where("user_id = ?", userID).Find(&list).Error
	return list, err
}

func (r *repository) UpsertPhrases(phrase *UserCustomPhrase) error {
	if r.db == nil {
		return errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	phrase.UpdatedAt = time.Now()
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"categories", "updated_at"}),
	}).Create(phrase).Error
}

func (r *repository) DeletePhrases(userID uuid.UUID, languageID string) error {
	if r.db == nil {
		return errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	return r.db.Where("user_id = ? AND language_id = ?", userID, languageID).Delete(&UserCustomPhrase{}).Error
}

func (r *repository) FindCategoriesByUserID(userID uuid.UUID) ([]UserPhraseCategory, error) {
	if r.db == nil {
		return nil, errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	var cats []UserPhraseCategory
	err := r.db.Where("user_id = ?", userID).Order("created_at asc").Find(&cats).Error
	return cats, err
}

func (r *repository) AddCategory(category *UserPhraseCategory) error {
	if r.db == nil {
		return errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	if category.ID == uuid.Nil {
		category.ID = uuid.New()
	}
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "name"}},
		DoUpdates: clause.AssignmentColumns([]string{"icon"}),
	}).Create(category).Error
}

func (r *repository) DeleteCategory(userID uuid.UUID, name string) error {
	if r.db == nil {
		return errors.New("database belum terhubung. Periksa konfigurasi DATABASE_URL di file .env")
	}
	return r.db.Where("user_id = ? AND name = ?", userID, name).Delete(&UserPhraseCategory{}).Error
}
