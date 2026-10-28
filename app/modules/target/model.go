package target

import (
	"time"

	"github.com/google/uuid"
)

type UserTarget struct {
	UserID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"userId"`
	CountryID string    `gorm:"primaryKey;size:10" json:"countryId"`
	Notes     string    `gorm:"type:text" json:"notes"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type AddTargetRequest struct {
	CountryID string `json:"countryId"`
}

type BatchSyncRequest struct {
	CountryIDs []string `json:"countryIds"`
}

type SaveNotesRequest struct {
	Notes string `json:"notes"`
}

type TargetsResponse struct {
	Targets []string          `json:"targets"`
	Notes   map[string]string `json:"notes"`
}
