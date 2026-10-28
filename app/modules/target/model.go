package target

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserTarget struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index:idx_user_country,unique" json:"userId"`
	CountryID string    `gorm:"type:text;not null;index:idx_user_country,unique" json:"countryId"`
	Notes     string    `gorm:"type:text" json:"notes"`
	CreatedAt time.Time `json:"createdAt"`
}

func (UserTarget) TableName() string {
	return "user_targets"
}

func (t *UserTarget) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
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
