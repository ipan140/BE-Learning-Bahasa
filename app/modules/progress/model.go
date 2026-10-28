package progress

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type UserStudyProgress struct {
	UserID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"userId"`
	TableID         string         `gorm:"primaryKey;size:100" json:"tableId"`
	MemorizedChars  datatypes.JSON `gorm:"type:jsonb" json:"memorizedChars"` // Array huruf hafal ["あ", "い"]
	Score           int            `gorm:"default:0" json:"score"`
	LastPracticedAt time.Time      `json:"lastPracticedAt"`
}

type SaveProgressRequest struct {
	TableID        string         `json:"tableId"`
	MemorizedChars datatypes.JSON `json:"memorizedChars"`
	Score          int            `json:"score"`
}

type SaveQuizScoreRequest struct {
	TableID string `json:"tableId"`
	Score   int    `json:"score"`
}
