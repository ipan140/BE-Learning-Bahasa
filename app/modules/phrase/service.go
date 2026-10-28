package phrase

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Service interface {
	GetPhrases(userID uuid.UUID) (map[string]datatypes.JSON, error)
	SavePhrases(userID uuid.UUID, req SavePhrasesRequest) error
	DeletePhrases(userID uuid.UUID, languageID string) error

	GetCategories(userID uuid.UUID) ([]UserPhraseCategory, error)
	SaveCategory(userID uuid.UUID, req SaveCategoryRequest) error
	DeleteCategory(userID uuid.UUID, name string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetPhrases(userID uuid.UUID) (map[string]datatypes.JSON, error) {
	list, err := s.repo.FindPhrasesByUserID(userID)
	if err != nil {
		return nil, err
	}

	result := make(map[string]datatypes.JSON)
	for _, p := range list {
		result[p.LanguageID] = p.Categories
	}
	return result, nil
}

func (s *service) SavePhrases(userID uuid.UUID, req SavePhrasesRequest) error {
	langID := strings.TrimSpace(req.LanguageID)
	if langID == "" {
		return errors.New("languageId wajib diisi")
	}

	cats := req.Categories
	if len(cats) == 0 {
		cats = datatypes.JSON([]byte("[]"))
	}

	phraseID := fmt.Sprintf("phrases-%s-%s", userID.String(), langID)
	return s.repo.UpsertPhrases(&UserCustomPhrase{
		ID:         phraseID,
		UserID:     userID,
		LanguageID: langID,
		Categories: cats,
		UpdatedAt:  time.Now(),
	})
}

func (s *service) DeletePhrases(userID uuid.UUID, languageID string) error {
	langID := strings.TrimSpace(languageID)
	if langID == "" {
		return errors.New("languageId wajib diisi")
	}
	return s.repo.DeletePhrases(userID, langID)
}

func (s *service) GetCategories(userID uuid.UUID) ([]UserPhraseCategory, error) {
	return s.repo.FindCategoriesByUserID(userID)
}

func (s *service) SaveCategory(userID uuid.UUID, req SaveCategoryRequest) error {
	catName := strings.TrimSpace(req.Name)
	if catName == "" {
		return errors.New("nama kategori wajib diisi")
	}

	return s.repo.AddCategory(&UserPhraseCategory{
		ID:        uuid.New(),
		UserID:    userID,
		Name:      catName,
		Icon:      strings.TrimSpace(req.Icon),
		CreatedAt: time.Now(),
	})
}

func (s *service) DeleteCategory(userID uuid.UUID, name string) error {
	catName := strings.TrimSpace(name)
	if catName == "" {
		return errors.New("nama kategori wajib diisi")
	}
	return s.repo.DeleteCategory(userID, catName)
}
