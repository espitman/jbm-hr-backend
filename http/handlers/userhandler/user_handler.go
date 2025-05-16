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

	return dto.SuccessJSON(c, interface{}(nil))
}

// VerifyOTP handles the OTP verification
// @Summary Verify OTP
// @Description Verify OTP and return JWT token and user data if valid
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

	token, user, err := h.userService.VerifyOTP(c.Request().Context(), req.Email, req.OTP)
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

	return dto.SuccessJSON(c, VerifyOTPData{
		Token: token,
		User: VerifyOTPUserData{
			ID:        user.ID,
			Email:     user.Email,
			Phone:     user.Phone,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role,
			Avatar:    user.Avatar,
		},
	})
}

// GetMe handles the request to get the current user's information
// @Summary Get current user data
// @Description Get the current user's data from the JWT token
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} GetMeResponse
// @Failure 401 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/users/me [get]
func (h *UserHandler) GetMe(c echo.Context) error {
	// Get user claims from context (set by JWT middleware)
	claims := c.Get("user").(*utils.Claims)

	user, err := h.userService.GetUserByID(c.Request().Context(), claims.UserID)
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, GetMeData{
		ID:        user.ID,
		Email:     user.Email,
		Phone:     user.Phone,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		Avatar:    user.Avatar,
		Department: convertToDepartmentDTO(
			user.DepartmentID,
			user.DepartmentTitle,
			user.DepartmentIcon,
			user.DepartmentShortName,
		),
		Birthdate:            user.Birthdate,
		CooperationStartDate: user.CooperationStartDate,
		PersonnelNumber:      &user.PersonnelNumber,
		NationalCode:         &user.NationalCode,
	})
}

// UpdateAvatar handles updating the current user's avatar
// @Summary Update current user avatar
// @Description Update the current user's avatar URL
// @Tags users
// @Accept json
// @Produce json
// @Param request body UpdateUserAvatarRequest true "Update Avatar"
// @Success 200 {object} UpdateUserResponse
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/users/avatar [put]
func (h *UserHandler) UpdateAvatar(c echo.Context) error {
	// Get user claims from context (set by JWT middleware)
	claims := c.Get("user").(*utils.Claims)

	var req UpdateUserAvatarRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, "invalid request format")
	}

	if err := utils.ValidateStruct(req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	// Update avatar
	updatedUser, err := h.userService.UpdateAvatar(c.Request().Context(), claims.UserID, req.Avatar)
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, UpdateUserResponse{
		Data: UserData{
			ID:        updatedUser.ID,
			Email:     updatedUser.Email,
			Phone:     updatedUser.Phone,
			FirstName: updatedUser.FirstName,
			LastName:  updatedUser.LastName,
			Role:      updatedUser.Role,
			Avatar:    updatedUser.Avatar,
			Department: convertToDepartmentDTO(
				updatedUser.DepartmentID,
				updatedUser.DepartmentTitle,
				updatedUser.DepartmentIcon,
				updatedUser.DepartmentShortName,
			),
			Birthdate:            updatedUser.Birthdate,
			CooperationStartDate: updatedUser.CooperationStartDate,
			PersonnelNumber:      updatedUser.PersonnelNumber,
			NationalCode:         updatedUser.NationalCode,
		},
	})
}
