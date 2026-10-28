package preference

import (
	"time"

	"github.com/google/uuid"
)

type UserPreference struct {
	UserID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"userId"`
	Theme        string    `gorm:"type:text;default:'light'" json:"theme"`
	AudioSpeed   float64   `gorm:"type:numeric;default:1.0" json:"audioSpeed"`
	AudioEnabled bool      `gorm:"type:bool;default:true" json:"audioEnabled"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type SavePreferenceRequest struct {
	AudioSpeed   float64 `json:"audioSpeed"`
	AudioEnabled *bool   `json:"audioEnabled"`
	Theme        string  `json:"theme"`
}
