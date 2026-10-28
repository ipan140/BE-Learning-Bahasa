package progress

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Service interface {
	GetProgress(userID uuid.UUID) (map[string][]string, error)
	SaveProgress(userID uuid.UUID, req SaveProgressRequest) error
	SaveQuizScore(userID uuid.UUID, req SaveQuizScoreRequest) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetProgress(userID uuid.UUID) (map[string][]string, error) {
	list, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	res := make(map[string][]string)
	for _, item := range list {
		var chars []string
		if len(item.MemorizedChars) > 0 {
			_ = json.Unmarshal(item.MemorizedChars, &chars)
		}
		if chars == nil {
			chars = []string{}
		}
		res[item.TableID] = chars
	}

	return res, nil
}

func (s *service) SaveProgress(userID uuid.UUID, req SaveProgressRequest) error {
	tableID := strings.TrimSpace(req.TableID)
	if tableID == "" {
		return errors.New("tableId wajib diisi")
	}

	rawChars := req.MemorizedChars
	if len(rawChars) == 0 {
		rawChars = datatypes.JSON([]byte("[]"))
	}

	return s.repo.Upsert(&UserStudyProgress{
		ID:              uuid.New(),
		UserID:          userID,
		TableID:         tableID,
		MemorizedChars:  rawChars,
		Score:           req.Score,
		StreakCount:     req.StreakCount,
		LastPracticedAt: time.Now(),
		Metadata:        req.Metadata,
	})
}

func (s *service) SaveQuizScore(userID uuid.UUID, req SaveQuizScoreRequest) error {
	tableID := strings.TrimSpace(req.TableID)
	if tableID == "" {
		return errors.New("tableId wajib diisi")
	}

	return s.repo.SaveScore(userID, tableID, req.Score, req.StreakCount)
}
