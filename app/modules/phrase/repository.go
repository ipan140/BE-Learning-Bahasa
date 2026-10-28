package phrase

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindPhrasesByUserID(userID uuid.UUID) ([]UserCustomPhrase, error)
	UpsertPhrases(phrase *UserCustomPhrase) error
	DeletePhrases(userID uuid.UUID, languageID string) error

	FindCategoriesByUserID(userID uuid.UUID) ([]string, error)
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
	var list []UserCustomPhrase
	err := r.db.Where("user_id = ?", userID).Find(&list).Error
	return list, err
}

func (r *repository) UpsertPhrases(phrase *UserCustomPhrase) error {
	phrase.UpdatedAt = time.Now()
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "language_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"categories", "updated_at"}),
	}).Create(phrase).Error
}

func (r *repository) DeletePhrases(userID uuid.UUID, languageID string) error {
	return r.db.Where("user_id = ? AND language_id = ?", userID, languageID).Delete(&UserCustomPhrase{}).Error
}

func (r *repository) FindCategoriesByUserID(userID uuid.UUID) ([]string, error) {
	var cats []UserPhraseCategory
	err := r.db.Where("user_id = ?", userID).Order("created_at asc").Find(&cats).Error
	if err != nil {
		return nil, err
	}

	names := make([]string, len(cats))
	for i, c := range cats {
		names[i] = c.Name
	}
	return names, nil
}

func (r *repository) AddCategory(category *UserPhraseCategory) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "name"}},
		DoNothing: true,
	}).Create(category).Error
}

func (r *repository) DeleteCategory(userID uuid.UUID, name string) error {
	return r.db.Where("user_id = ? AND name = ?", userID, name).Delete(&UserPhraseCategory{}).Error
}
