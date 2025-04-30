package resumehandler

import (
	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
)

// CreateResumeResponse represents the response for creating a resume
type CreateResumeResponse struct {
	dto.Response
	Resume contract.Resume `json:"resume"`
}

// GetResumeResponse represents the response for getting a resume
type GetResumeResponse struct {
	dto.Response
	Resume contract.Resume `json:"resume"`
}

// ListResumeResponse represents the response for listing resumes
type ListResumeResponse struct {
	dto.Response
	Resumes []contract.Resume `json:"resumes"`
}

// UpdateStatusResponse represents the response for updating a resume's status
type UpdateStatusResponse struct {
	dto.Response
	Resume contract.Resume `json:"resume"`
}
