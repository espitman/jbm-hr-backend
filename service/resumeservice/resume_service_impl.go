package resumeservice

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/database/repository/resume"
)

type service struct {
	resumeRepo resume.Repository
}

// New creates a new ResumeService instance
func New(resumeRepo resume.Repository) Service {
	return &service{
		resumeRepo: resumeRepo,
	}
}

// Create creates a new resume
func (s *service) Create(ctx context.Context, input *contract.ResumeInput) (*contract.Resume, error) {
	return s.resumeRepo.Create(ctx, input)
}

// Update updates an existing resume
func (s *service) Update(ctx context.Context, id int, input *contract.ResumeInput) (*contract.Resume, error) {
	// Check if resume exists
	_, err := s.resumeRepo.GetByID(ctx, id)
	if err != nil {
		return nil, contract.ErrResumeNotFound
	}

	return s.resumeRepo.Update(ctx, id, input)
}

// GetByID retrieves a resume by its ID
func (s *service) GetByID(ctx context.Context, id int) (*contract.Resume, error) {
	resume, err := s.resumeRepo.GetByID(ctx, id)
	if err != nil {
		return nil, contract.ErrResumeNotFound
	}
	return resume, nil
}

// List retrieves a paginated list of resumes
func (s *service) List(ctx context.Context, page, limit int) ([]*contract.Resume, int, error) {
	return s.resumeRepo.List(ctx, page, limit)
}

// Delete deletes a resume by its ID
func (s *service) Delete(ctx context.Context, id int) error {
	// Check if resume exists
	_, err := s.resumeRepo.GetByID(ctx, id)
	if err != nil {
		return contract.ErrResumeNotFound
	}

	return s.resumeRepo.Delete(ctx, id)
}
