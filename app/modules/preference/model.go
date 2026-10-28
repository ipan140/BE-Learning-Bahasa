package preference

import (
	"time"

	"github.com/google/uuid"
)

type UserPreference struct {
	UserID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"userId"`
	AudioSpeed   float64   `gorm:"default:1.0" json:"audioSpeed"`
	AudioEnabled bool      `gorm:"default:true" json:"audioEnabled"`
	Theme        string    `gorm:"size:20;default:'light'" json:"theme"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type SavePreferenceRequest struct {
	AudioSpeed   float64 `json:"audioSpeed"`
	AudioEnabled *bool   `json:"audioEnabled"`
	Theme        string  `json:"theme"`
}
