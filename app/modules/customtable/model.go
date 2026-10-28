package customtable

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type UserCustomTable struct {
	ID          string         `gorm:"primaryKey;size:100" json:"id"`
	UserID      uuid.UUID      `gorm:"type:uuid;index;not null" json:"userId"`
	Name        string         `gorm:"size:150;not null" json:"name"`
	VoiceLang   string         `gorm:"size:20" json:"voiceLang"`
	Description string         `gorm:"type:text" json:"description"`
	Rows        datatypes.JSON `gorm:"type:jsonb" json:"rows"` // Susunan grid aksara per baris
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

type UpsertTableRequest struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	VoiceLang   string         `json:"voiceLang"`
	Description string         `json:"description"`
	Rows        datatypes.JSON `json:"rows"`
}
