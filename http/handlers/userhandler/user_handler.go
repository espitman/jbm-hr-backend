package userhandler

import (
	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	// Add dependencies here later
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// RequestOTP handles the OTP request by email
// @Summary Request OTP
// @Description Send OTP to user's email
// @Tags user
// @Accept json
// @Produce json
// @Param request body RequestOTPRequest true "Request OTP"
// @Success 200 {object} RequestOTPResponse
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/users/request-otp [post]
func (h *UserHandler) RequestOTP(c echo.Context) error {
	var req RequestOTPRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	// TODO: Implement actual OTP generation and email sending
	// For now, just return a mock response
	response := RequestOTPResponse{}
	return dto.SuccessJSON(c, response)
}

// VerifyOTP handles the OTP verification
// @Summary Verify OTP
// @Description Verify OTP and return JWT token if valid
// @Tags user
// @Accept json
// @Produce json
// @Param request body VerifyOTPRequest true "Verify OTP"
// @Success 200 {object} VerifyOTPResponse
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/users/verify-otp [post]
func (h *UserHandler) VerifyOTP(c echo.Context) error {
	var req VerifyOTPRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	// TODO: Implement actual OTP verification
	// For now, just return a mock response
	response := VerifyOTPResponse{}
	response.Data.Token = "mock-jwt-token"

	return dto.SuccessJSON(c, response)
}
