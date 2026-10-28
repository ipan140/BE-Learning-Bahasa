package phrase

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type UserCustomPhrase struct {
	ID         string         `gorm:"primaryKey;size:150" json:"id"`
	UserID     uuid.UUID      `gorm:"type:uuid;index;not null" json:"userId"`
	LanguageID string         `gorm:"size:50;not null" json:"languageId"`
	Categories datatypes.JSON `gorm:"type:jsonb" json:"categories"`
	UpdatedAt  time.Time      `json:"updatedAt"`
}

type UserPhraseCategory struct {
	UserID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"userId"`
	Name      string    `gorm:"primaryKey;size:100" json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type SavePhrasesRequest struct {
	LanguageID string         `json:"languageId"`
	Categories datatypes.JSON `json:"categories"`
}

type SaveCategoryRequest struct {
	Name string `json:"name"`
}
