package resumehandler

import (
	"net/http"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/service/resumeservice"
	"github.com/labstack/echo/v4"
)

// ResumeHandler handles HTTP requests for resumes by users
type ResumeHandler struct {
	resumeService resumeservice.Service
}

// NewResumeHandler creates a new ResumeHandler
func NewResumeHandler(resumeService resumeservice.Service) *ResumeHandler {
	return &ResumeHandler{
		resumeService: resumeService,
	}
}

// Create handles creating a new resume
// @Summary Create a resume
// @Description Create a new resume
// @Tags resumes
// @Accept json
// @Produce json
// @Param input body CreateResumeRequest true "Resume input"
// @Success 201 {object} CreateResumeResponse
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/resumes [post]
func (h *ResumeHandler) Create(c echo.Context) error {
	var req CreateResumeRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, "invalid input")
	}

	// Get user ID from JWT claims
	userID := c.Get("user_id").(int)

	// Convert request to service input
	input := &contract.ResumeInput{
		IntroducedName:  req.IntroducedName,
		IntroducedPhone: req.IntroducedPhone,
		Position:        req.Position,
		File:            req.File,
		UserID:          userID,
	}

	resume, err := h.resumeService.Create(c.Request().Context(), input)
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, CreateResumeResponse{Resume: *resume})
}
