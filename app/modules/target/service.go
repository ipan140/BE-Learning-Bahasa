package target

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	GetTargets(userID uuid.UUID) (*TargetsResponse, error)
	AddTarget(userID uuid.UUID, countryID string) error
	RemoveTarget(userID uuid.UUID, countryID string) error
	BatchSync(userID uuid.UUID, countryIDs []string) error
	SaveNotes(userID uuid.UUID, countryID string, notes string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetTargets(userID uuid.UUID) (*TargetsResponse, error) {
	items, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	targets := make([]string, 0, len(items))
	notes := make(map[string]string)

	for _, item := range items {
		targets = append(targets, item.CountryID)
		if item.Notes != "" {
			notes[item.CountryID] = item.Notes
		}
	}

	return &TargetsResponse{
		Targets: targets,
		Notes:   notes,
	}, nil
}

func (s *service) AddTarget(userID uuid.UUID, countryID string) error {
	cid := strings.TrimSpace(countryID)
	if cid == "" {
		return errors.New("countryId wajib diisi")
	}

	now := time.Now()
	return s.repo.Add(&UserTarget{
		UserID:    userID,
		CountryID: cid,
		CreatedAt: now,
	})
}

func (s *service) RemoveTarget(userID uuid.UUID, countryID string) error {
	cid := strings.TrimSpace(countryID)
	if cid == "" {
		return errors.New("countryId wajib diisi")
	}
	return s.repo.Remove(userID, cid)
}

func (s *service) BatchSync(userID uuid.UUID, countryIDs []string) error {
	cleanIDs := make([]string, 0, len(countryIDs))
	for _, cid := range countryIDs {
		trimmed := strings.TrimSpace(cid)
		if trimmed != "" {
			cleanIDs = append(cleanIDs, trimmed)
		}
	}
	return s.repo.BatchSync(userID, cleanIDs)
}

func (s *service) SaveNotes(userID uuid.UUID, countryID string, notes string) error {
	cid := strings.TrimSpace(countryID)
	if cid == "" {
		return errors.New("countryId wajib diisi")
	}
	return s.repo.SaveNotes(userID, cid, strings.TrimSpace(notes))
}
