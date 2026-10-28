package progress

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UserStudyProgress struct {
	ID              uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	UserID          uuid.UUID      `gorm:"type:uuid;not null;index:idx_user_table,unique" json:"userId"`
	TableID         string         `gorm:"type:text;not null;index:idx_user_table,unique" json:"tableId"`
	MemorizedChars  datatypes.JSON `gorm:"type:jsonb" json:"memorizedChars"` // Array huruf hafal
	Score           int            `gorm:"type:int4;default:0" json:"score"`
	StreakCount     int            `gorm:"type:int4;default:0" json:"streakCount"`
	LastPracticedAt time.Time      `json:"lastPracticedAt"`
	Metadata        datatypes.JSON `gorm:"type:jsonb" json:"metadata,omitempty"`
}

func (p *UserStudyProgress) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

type SaveProgressRequest struct {
	TableID        string         `json:"tableId"`
	MemorizedChars datatypes.JSON `json:"memorizedChars"`
	Score          int            `json:"score"`
	StreakCount    int            `json:"streakCount"`
	Metadata       datatypes.JSON `json:"metadata"`
}

type SaveQuizScoreRequest struct {
	TableID     string `json:"tableId"`
	Score       int    `json:"score"`
	StreakCount int    `json:"streakCount"`
}
