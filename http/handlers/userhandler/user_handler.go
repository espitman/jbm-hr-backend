package userhandler

import (
	"net/http"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/service/userservice"
	"github.com/espitman/jbm-hr-backend/utils"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService userservice.Service
}

func NewUserHandler(userService userservice.Service) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// RequestOTP handles the OTP request by email
// @Summary Request OTP
// @Description Send OTP to user's email
// @Tags users
// @Accept json
// @Produce json
// @Param request body RequestOTPRequest true "Request OTP"
// @Success 200 {object} RequestOTPResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/users/request-otp [post]
func (h *UserHandler) RequestOTP(c echo.Context) error {
	var req RequestOTPRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	if err := utils.ValidateStruct(req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	_, err := h.userService.RequestOTP(c.Request().Context(), req.Email)
	if err != nil {
		if err == contract.ErrUserNotFound {
			return dto.ErrorJSON(c, http.StatusNotFound, err.Error())
		}
		if err == contract.ErrActiveOTPExists {
			return dto.BadRequestJSON(c, err.Error())
		}
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	response := RequestOTPResponse{}
	return dto.SuccessJSON(c, response)
}

// VerifyOTP handles the OTP verification
// @Summary Verify OTP
// @Description Verify OTP and return JWT token if valid
// @Tags users
// @Accept json
// @Produce json
// @Param request body VerifyOTPRequest true "Verify OTP"
// @Success 200 {object} VerifyOTPResponse
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/users/verify-otp [post]
func (h *UserHandler) VerifyOTP(c echo.Context) error {
	var req VerifyOTPRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	if err := utils.ValidateStruct(req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	valid, err := h.userService.VerifyOTP(c.Request().Context(), req.Email, req.OTP)
	if err != nil {
		if err == contract.ErrUserNotFound {
			return dto.ErrorJSON(c, http.StatusNotFound, err.Error())
		}
		if err == contract.ErrOTPNotFound || err == contract.ErrOTPInvalid ||
			err == contract.ErrOTPExpired || err == contract.ErrOTPAlreadyUsed {
			return dto.BadRequestJSON(c, err.Error())
		}
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	if !valid {
		return dto.ErrorJSON(c, http.StatusUnauthorized, "Invalid OTP")
	}

	response := VerifyOTPResponse{}
	response.Data = VerifyOTPData{
		Token: "mock-jwt-token", // TODO: Generate actual JWT token
	}

	return dto.SuccessJSON(c, response)
}
