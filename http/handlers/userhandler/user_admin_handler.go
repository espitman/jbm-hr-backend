package userhandler

import (
	"net/http"
	"strconv"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/utils"
	"github.com/labstack/echo/v4"
)

// RegisterUser handles the user registration
// @Summary Register a new user
// @Description Register a new user in the system (Admin only)
// @Tags admin - users
// @Accept json
// @Produce json
// @Param request body RegisterUserRequest true "Register User"
// @Success 201 {object} RegisterUserResponse
// @Failure 400 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/users/register [post]
func (h *UserHandler) RegisterUser(c echo.Context) error {
	var req RegisterUserRequest

	// Bind and validate request
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, "Invalid request format")
	}

	if err := utils.ValidateStruct(req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	// Register user
	user, err := h.userService.RegisterUser(c.Request().Context(), &contract.RegisterUserInput{
		Email:     req.Email,
		Phone:     req.Phone,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      req.Role,
		Avatar:    req.Avatar,
	})
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	// Prepare response
	return dto.CreatedJSON(c, RegisterUserData{
		ID:        user.ID,
		Email:     user.Email,
		Phone:     user.Phone,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		Avatar:    user.Avatar,
	})
}

// ListUsers handles listing all users with pagination
// @Summary List all users
// @Description Get a paginated list of all users in the system (Admin only)
// @Tags admin - users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} ListUsersResponse
// @Failure 400 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/users [get]
func (h *UserHandler) ListUsers(c echo.Context) error {
	// Get pagination parameters
	page := utils.GetQueryParamInt(c, "page", 1)
	limit := utils.GetQueryParamInt(c, "limit", 10)

	// Get users with pagination
	users, total, err := h.userService.ListUsers(c.Request().Context(), page, limit)
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	// Prepare response data
	usersData := make([]UserData, len(users))
	for i, user := range users {
		usersData[i] = UserData{
			ID:        user.ID,
			Email:     user.Email,
			Phone:     user.Phone,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role,
			Avatar:    user.Avatar,
		}
	}

	// Return paginated response
	return dto.SuccessJSON(c, ListUsersData{
		Users: usersData,
		Total: total,
	})
}

// UpdatePassword handles updating a user's password
// @Summary Update user password
// @Description Update a user's password
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body UpdatePasswordRequest true "Update Password"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/users/{id}/password [put]
func (h *UserHandler) UpdatePassword(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid user ID")
	}

	var req UpdatePasswordRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, "invalid request format")
	}

	if err := utils.ValidateStruct(req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	// Update the password
	err = h.userService.UpdatePassword(c.Request().Context(), id, &contract.UpdatePasswordInput{
		Password: req.Password,
	})
	if err != nil {
		if err == contract.ErrUserNotFound {
			return dto.ErrorJSON(c, http.StatusNotFound, err.Error())
		}
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, nil)
}
