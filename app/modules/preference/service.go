package preference

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	GetPreferences(userID uuid.UUID) (*UserPreference, error)
	SavePreferences(userID uuid.UUID, req SavePreferenceRequest) (*UserPreference, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetPreferences(userID uuid.UUID) (*UserPreference, error) {
	pref, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	if pref == nil {
		// Nilai default jika belum pernah disimpan
		return &UserPreference{
			UserID:       userID,
			AudioSpeed:   1.0,
			AudioEnabled: true,
			Theme:        "light",
		}, nil
	}

	return pref, nil
}

func (s *service) SavePreferences(userID uuid.UUID, req SavePreferenceRequest) (*UserPreference, error) {
	current, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	if current == nil {
		current = &UserPreference{
			UserID:       userID,
			AudioSpeed:   1.0,
			AudioEnabled: true,
			Theme:        "light",
			UpdatedAt:    time.Now(),
		}
	}

	if req.AudioSpeed > 0 {
		current.AudioSpeed = req.AudioSpeed
	}
	if req.AudioEnabled != nil {
		current.AudioEnabled = *req.AudioEnabled
	}
	if strings.TrimSpace(req.Theme) != "" {
		current.Theme = strings.TrimSpace(req.Theme)
	}

	if err := s.repo.Upsert(current); err != nil {
		return nil, err
	}

	return current, nil
}
