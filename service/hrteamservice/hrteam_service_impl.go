package hrteamservice

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
	hrteamRepo "github.com/espitman/jbm-hr-backend/database/repository/hrteam"
)

// hrteamService implements the Service interface
type hrteamService struct {
	repository hrteamRepo.Repository
}

// New creates a new HR team service instance
func New(repository hrteamRepo.Repository) Service {
	return &hrteamService{
		repository: repository,
	}
}

// GetAllHRTeam returns all HR team members
func (s *hrteamService) GetAllHRTeam(ctx context.Context) ([]*contract.HRTeam, error) {
	return s.repository.GetAll(ctx)
}

// CreateHRTeam creates a new HR team member
func (s *hrteamService) CreateHRTeam(ctx context.Context, input *contract.HRTeamInput) (*contract.HRTeam, error) {
	return s.repository.Create(ctx, input)
}

// GetHRTeamByID returns an HR team member by their ID
func (s *hrteamService) GetHRTeamByID(ctx context.Context, id int) (*contract.HRTeam, error) {
	return s.repository.GetByID(ctx, id)
}

// UpdateHRTeam updates an existing HR team member
func (s *hrteamService) UpdateHRTeam(ctx context.Context, id int, input *contract.HRTeamInput) (*contract.HRTeam, error) {
	return s.repository.Update(ctx, id, input)
}

// DeleteHRTeam deletes an HR team member by their ID
func (s *hrteamService) DeleteHRTeam(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
