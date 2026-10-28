package phrase

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UserCustomPhrase struct {
	ID         string         `gorm:"primaryKey;type:text" json:"id"` // phrases-{userId}-{languageId}
	UserID     uuid.UUID      `gorm:"type:uuid;index;not null" json:"userId"`
	LanguageID string         `gorm:"type:text;not null" json:"languageId"`
	Categories datatypes.JSON `gorm:"type:jsonb" json:"categories"`
	UpdatedAt  time.Time      `json:"updatedAt"`
}

type UserPhraseCategory struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index:idx_user_catname,unique" json:"userId"`
	Name      string    `gorm:"type:text;not null;index:idx_user_catname,unique" json:"name"`
	Icon      string    `gorm:"type:text" json:"icon"`
	CreatedAt time.Time `json:"createdAt"`
}

func (c *UserPhraseCategory) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

type SavePhrasesRequest struct {
	LanguageID string         `json:"languageId"`
	Categories datatypes.JSON `json:"categories"`
}

type SaveCategoryRequest struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}
