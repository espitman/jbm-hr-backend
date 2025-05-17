package infoservice

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/database/repository/department"
	"github.com/espitman/jbm-hr-backend/database/repository/request"
	"github.com/espitman/jbm-hr-backend/database/repository/resume"
	"github.com/espitman/jbm-hr-backend/database/repository/user"
)

// InfoService handles info-related business logic
type InfoService struct {
	userRepository       user.Repository
	requestRepository    request.Repository
	resumeRepository     resume.Repository
	departmentRepository department.Repository
}

// NewInfoService creates a new instance of InfoService
func NewInfoService(
	userRepository user.Repository,
	requestRepository request.Repository,
	resumeRepository resume.Repository,
	departmentRepository department.Repository,
) *InfoService {
	return &InfoService{
		userRepository:       userRepository,
		requestRepository:    requestRepository,
		resumeRepository:     resumeRepository,
		departmentRepository: departmentRepository,
	}
}

// GetDashboardInfo returns the dashboard information
func (s *InfoService) GetDashboardInfo(ctx context.Context) (*contract.DashboardInfo, error) {
	// Get total counts
	userCount, err := s.userRepository.GetTotalCount(ctx)
	if err != nil {
		return nil, err
	}

	requestCount, err := s.requestRepository.GetTotalCount(ctx, &contract.RequestFilter{})
	if err != nil {
		return nil, err
	}

	resumeCount, err := s.resumeRepository.GetTotalCount(ctx)
	if err != nil {
		return nil, err
	}

	departmentCount, err := s.departmentRepository.GetTotalCount(ctx)
	if err != nil {
		return nil, err
	}

	return &contract.DashboardInfo{
		Users:       userCount,
		Requests:    requestCount,
		Resumes:     resumeCount,
		Departments: departmentCount,
	}, nil
}
