package resumehandler

import (
	"net/http"
	"strconv"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/service/resumeservice"
	"github.com/labstack/echo/v4"
)

// ResumeAdminHandler handles HTTP requests for resumes by admins
type ResumeAdminHandler struct {
	resumeService resumeservice.Service
}

// NewResumeAdminHandler creates a new ResumeAdminHandler
func NewResumeAdminHandler(resumeService resumeservice.Service) *ResumeAdminHandler {
	return &ResumeAdminHandler{
		resumeService: resumeService,
	}
}

// Get handles retrieving a resume by ID
// @Summary Get a resume
// @Description Get a resume by ID
// @Tags resumes - admin
// @Accept json
// @Produce json
// @Param id path int true "Resume ID"
// @Success 200 {object} GetResumeResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/resumes/{id} [get]
func (h *ResumeAdminHandler) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid id")
	}

	resume, err := h.resumeService.GetByID(c.Request().Context(), id)
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	if resume == nil {
		return dto.ErrorJSON(c, http.StatusNotFound, "resume not found")
	}

	return dto.SuccessJSON(c, resume)
}

// List handles retrieving all resumes
// @Summary List resumes
// @Description Get all resumes
// @Tags resumes - admin
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} ListResumeResponse
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/resumes [get]
func (h *ResumeAdminHandler) List(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit < 1 {
		limit = 10
	}

	resumes, total, err := h.resumeService.List(c.Request().Context(), page, limit)
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	// Convert []*contract.Resume to []contract.Resume
	convertedResumes := make([]contract.Resume, len(resumes))
	for i, r := range resumes {
		convertedResumes[i] = *r
	}

	return dto.SuccessJSON(c, ResumeListData{
		Items: convertedResumes,
		Total: total,
	})
}

// UpdateStatus handles updating a resume's status
// @Summary Update resume status
// @Description Update a resume's status
// @Tags resumes - admin
// @Accept json
// @Produce json
// @Param id path int true "Resume ID"
// @Param input body UpdateStatusRequest true "Status update input"
// @Success 200 {object} UpdateStatusResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/resumes/{id}/status [put]
func (h *ResumeAdminHandler) UpdateStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid id")
	}

	var req UpdateStatusRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, "invalid input")
	}

	// Get existing resume
	resume, err := h.resumeService.GetByID(c.Request().Context(), id)
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	if resume == nil {
		return dto.ErrorJSON(c, http.StatusNotFound, "resume not found")
	}

	// Update only the status
	input := &contract.ResumeInput{
		IntroducedName:  resume.IntroducedName,
		IntroducedPhone: resume.IntroducedPhone,
		Position:        resume.Position,
		File:            resume.File,
		UserID:          resume.UserID,
		Status:          req.Status,
	}

	updatedResume, err := h.resumeService.Update(c.Request().Context(), id, input)
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, updatedResume)
}
