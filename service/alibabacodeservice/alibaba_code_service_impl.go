package alibabacodeservice

import (
	"context"
	"errors"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/database/repository/alibaba_code"
)

// ServiceImpl implements the Service interface
type ServiceImpl struct {
	repo alibaba_code.Repository
}

// NewService creates a new ServiceImpl
func NewService(repo alibaba_code.Repository) *ServiceImpl {
	return &ServiceImpl{
		repo: repo,
	}
}

// GetAll retrieves all Alibaba codes with optional filters
func (s *ServiceImpl) GetAll(ctx context.Context, filters *contract.AlibabaCodeFilters) ([]*contract.AlibabaCode, error) {
	return s.repo.GetAll(ctx, filters)
}

// GetByID retrieves an Alibaba code by its ID
func (s *ServiceImpl) GetByID(ctx context.Context, id int) (*contract.AlibabaCode, error) {
	return s.repo.GetByID(ctx, id)
}

// GetByCode retrieves an Alibaba code by its code
func (s *ServiceImpl) GetByCode(ctx context.Context, code string) (*contract.AlibabaCode, error) {
	return s.repo.GetByCode(ctx, code)
}

// Create creates a new Alibaba code
func (s *ServiceImpl) Create(ctx context.Context, req *contract.CreateAlibabaCodeInput) (*contract.AlibabaCode, error) {
	// Validate code type
	if !isValidCodeType(req.Type) {
		return nil, errors.New("invalid code type")
	}

	// Check if code already exists
	existingCode, err := s.repo.GetByCode(ctx, req.Code)
	if err == nil && existingCode != nil {
		return nil, errors.New("code already exists")
	}

	return s.repo.Create(ctx, req)
}

// Update updates an existing Alibaba code
func (s *ServiceImpl) Update(ctx context.Context, id int, req *contract.UpdateAlibabaCodeInput) (*contract.AlibabaCode, error) {
	// Check if code exists
	existingCode, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existingCode == nil {
		return nil, errors.New("code not found")
	}

	return s.repo.Update(ctx, id, req)
}

// Delete deletes an Alibaba code by its ID
func (s *ServiceImpl) Delete(ctx context.Context, id int) error {
	// Check if code exists
	existingCode, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingCode == nil {
		return errors.New("code not found")
	}

	return s.repo.Delete(ctx, id)
}

// isValidCodeType checks if the code type is valid
func isValidCodeType(codeType string) bool {
	validTypes := map[string]bool{
		"1m":  true,
		"3m":  true,
		"6m":  true,
		"12m": true,
		"25m": true,
	}
	return validTypes[codeType]
}
