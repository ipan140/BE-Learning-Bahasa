package customtable

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type UserCustomTable struct {
	ID          string         `gorm:"primaryKey;type:text" json:"id"`
	UserID      uuid.UUID      `gorm:"type:uuid;index;not null" json:"userId"`
	Name        string         `gorm:"type:text;not null" json:"name"`
	Script      string         `gorm:"type:text" json:"script"` // Voice language / script
	Description string         `gorm:"type:text" json:"description"`
	Rows        datatypes.JSON `gorm:"type:jsonb" json:"rows"` // Susunan grid aksara per baris
	CreatedAt   time.Time      `json:"createdAt"`
}

type UpsertTableRequest struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Script      string         `json:"script"`
	VoiceLang   string         `json:"voiceLang"` // fallback alias
	Description string         `json:"description"`
	Rows        datatypes.JSON `json:"rows"`
}
