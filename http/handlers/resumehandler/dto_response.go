package resumehandler

import (
	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
)

// CreateResumeResponse represents the response for creating a resume
type CreateResumeResponse struct {
	dto.Response
	Data contract.Resume `json:"data"`
}

// GetResumeResponse represents the response for getting a resume
type GetResumeResponse struct {
	dto.Response
	Data contract.Resume `json:"data"`
}

// ResumeListData represents the data structure for listing resumes
type ResumeListData struct {
	Items []contract.Resume `json:"items"`
	Total int               `json:"total"`
}

// ListResumeResponse represents the response for listing resumes
type ListResumeResponse struct {
	dto.Response
	Data ResumeListData `json:"data"`
}

// UpdateStatusResponse represents the response for updating a resume's status
type UpdateStatusResponse struct {
	dto.Response
	Data contract.Resume `json:"data"`
}
