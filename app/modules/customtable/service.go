package customtable

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	GetTables(userID uuid.UUID) ([]UserCustomTable, error)
	SaveTable(userID uuid.UUID, req UpsertTableRequest) (*UserCustomTable, error)
	DeleteTable(userID uuid.UUID, id string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetTables(userID uuid.UUID) ([]UserCustomTable, error) {
	return s.repo.FindByUserID(userID)
}

func (s *service) SaveTable(userID uuid.UUID, req UpsertTableRequest) (*UserCustomTable, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, errors.New("nama tabel wajib diisi")
	}

	tableID := strings.TrimSpace(req.ID)
	if tableID == "" {
		tableID = fmt.Sprintf("custom-%d", time.Now().UnixMilli())
	}

	now := time.Now()
	table := &UserCustomTable{
		ID:          tableID,
		UserID:      userID,
		Name:        name,
		VoiceLang:   req.VoiceLang,
		Description: req.Description,
		Rows:        req.Rows,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.repo.Upsert(table); err != nil {
		return nil, err
	}

	return table, nil
}

func (s *service) DeleteTable(userID uuid.UUID, id string) error {
	tableID := strings.TrimSpace(id)
	if tableID == "" {
		return errors.New("id tabel wajib diisi")
	}
	return s.repo.Delete(userID, tableID)
}
