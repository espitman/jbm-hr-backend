package userhandler

import (
	"net/http"
	"strconv"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/utils"
	"github.com/labstack/echo/v4"
)

// convertToDepartmentDTO converts department data to DTO format
func convertToDepartmentDTO(departmentID *int, title *string, icon *string, shortName *string) *DepartmentDTO {
	if departmentID == nil {
		return nil
	}
	return &DepartmentDTO{
		ID:        *departmentID,
		Title:     *title,
		Icon:      *icon,
		ShortName: *shortName,
	}
}

// RegisterUser handles the user registration
// @Summary Register a new user
// @Description Register a new user in the system (Admin only)
// @Tags users - admin
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
		Email:                req.Email,
		Phone:                req.Phone,
		FirstName:            req.FirstName,
		LastName:             req.LastName,
		Role:                 req.Role,
		Avatar:               req.Avatar,
		DepartmentID:         req.DepartmentID,
		Birthdate:            req.Birthdate,
		CooperationStartDate: req.CooperationStartDate,
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
		Department: convertToDepartmentDTO(
			user.DepartmentID,
			user.DepartmentTitle,
			user.DepartmentIcon,
			user.DepartmentShortName,
		),
		Birthdate:            user.Birthdate,
		CooperationStartDate: user.CooperationStartDate,
	})
}

// ListUsers handles listing all users
// @Summary List all users
// @Description Get a list of all users with pagination
// @Tags users - admin
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} ListUsersResponse
// @Failure 400 {object} dto.Response
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
			Department: convertToDepartmentDTO(
				user.DepartmentID,
				user.DepartmentTitle,
				user.DepartmentIcon,
				user.DepartmentShortName,
			),
			Birthdate:            user.Birthdate,
			CooperationStartDate: user.CooperationStartDate,
		}
	}

	// Return paginated response
	return dto.SuccessJSON(c, ListUsersData{
		Users: usersData,
		Total: total,
	})
}

// AdminLogin handles admin user login
// @Summary Admin login
// @Description Authenticate admin user with email and password
// @Tags users - admin
// @Accept json
// @Produce json
// @Param request body AdminLoginRequest true "Admin Login"
// @Success 200 {object} AdminLoginResponse
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/admin/login [post]
func (h *UserHandler) AdminLogin(c echo.Context) error {
	var req AdminLoginRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, "invalid request format")
	}

	if err := utils.ValidateStruct(req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	token, user, err := h.userService.AdminLogin(c.Request().Context(), req.Email, req.Password)
	if err != nil {
		if err == contract.ErrUserNotFound {
			return dto.ErrorJSON(c, http.StatusNotFound, err.Error())
		}
		if err.Error() == "invalid password" || err.Error() == "only admin users can login with password" || err.Error() == "password not set for this user" {
			return dto.ErrorJSON(c, http.StatusUnauthorized, err.Error())
		}
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, AdminLoginResponse{
		Token: token,
		User: UserData{
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

// GetUserByID handles getting a user by ID
// @Summary Get user by ID
// @Description Get user details by ID
// @Tags users - admin
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} UserData
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/users/{id} [get]
func (h *UserHandler) GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid user ID")
	}

	user, err := h.userService.GetUserByID(c.Request().Context(), id)
	if err != nil {
		if err == contract.ErrUserNotFound {
			return dto.ErrorJSON(c, http.StatusNotFound, "user not found")
		}
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, UserData{
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
	})
}

// UpdateUser handles updating a user's information
// @Summary Update user information
// @Description Update a user's information (Admin only)
// @Tags users - admin
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body contract.UpdateUserInput true "Update User"
// @Success 200 {object} UpdateUserResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/users/{id} [put]
func (h *UserHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid user ID")
	}

	var input contract.UpdateUserInput
	if err := c.Bind(&input); err != nil {
		return dto.BadRequestJSON(c, "invalid request format")
	}

	if err := utils.ValidateStruct(input); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	user, err := h.userService.UpdateUser(c.Request().Context(), id, &input)
	if err != nil {
		if err == contract.ErrUserNotFound {
			return dto.ErrorJSON(c, http.StatusNotFound, "user not found")
		}
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, UpdateUserResponse{
		Data: UserData{
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
		},
	})
}

// UpdateUserPassword handles updating a user's password by admin
// @Summary Update user password
// @Description Update a user's password (Admin only)
// @Tags users - admin
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body contract.UpdatePasswordInput true "Update Password"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/users/{id}/password [put]
func (h *UserHandler) UpdateUserPassword(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid user ID")
	}

	var input contract.UpdatePasswordInput
	if err := c.Bind(&input); err != nil {
		return dto.BadRequestJSON(c, "invalid request format")
	}

	if err := utils.ValidateStruct(input); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	err = h.userService.UpdateUserPassword(c.Request().Context(), id, &input)
	if err != nil {
		if err == contract.ErrUserNotFound {
			return dto.ErrorJSON(c, http.StatusNotFound, "user not found")
		}
		if err.Error() == "only admin users can set passwords" {
			return dto.ErrorJSON(c, http.StatusForbidden, err.Error())
		}
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, nil)
}

// UpdateUserAvatar handles updating a user's avatar
// @Summary Update user avatar
// @Description Update a user's avatar URL
// @Tags users - admin
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body UpdateUserAvatarRequest true "Update Avatar"
// @Success 200 {object} UpdateUserResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/users/{id}/avatar [put]
func (h *UserHandler) UpdateUserAvatar(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid user ID")
	}

	var req UpdateUserAvatarRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, "invalid request format")
	}

	if err := utils.ValidateStruct(req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	// Update avatar
	updatedUser, err := h.userService.UpdateAvatar(c.Request().Context(), id, req.Avatar)
	if err != nil {
		if err == contract.ErrUserNotFound {
			return dto.ErrorJSON(c, http.StatusNotFound, "user not found")
		}
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
		},
	})
}

// UpdateUserBirthdate handles updating a user's birthdate
// @Summary Update user birthdate
// @Description Update a user's birthdate (Admin only)
// @Tags users - admin
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body UpdateUserBirthdateRequest true "Update Birthdate"
// @Success 200 {object} UpdateUserResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/users/{id}/birthdate [put]
func (h *UserHandler) UpdateUserBirthdate(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid user ID")
	}

	var req UpdateUserBirthdateRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, "invalid request format")
	}

	if err := utils.ValidateStruct(req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	user, err := h.userService.UpdateBirthdate(c.Request().Context(), id, req.Birthdate)
	if err != nil {
		if err == contract.ErrUserNotFound {
			return dto.ErrorJSON(c, http.StatusNotFound, "user not found")
		}
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, UpdateUserResponse{
		Data: UserData{
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
		},
	})
}

// UpdateUserCooperationStartDate handles updating a user's cooperation start date
// @Summary Update user cooperation start date
// @Description Update a user's cooperation start date (Admin only)
// @Tags users - admin
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body UpdateUserCooperationStartDateRequest true "Update Cooperation Start Date"
// @Success 200 {object} UpdateUserResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/users/{id}/cooperation-start-date [put]
func (h *UserHandler) UpdateUserCooperationStartDate(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid user ID")
	}

	var req UpdateUserCooperationStartDateRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, "invalid request format")
	}

	if err := utils.ValidateStruct(req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	user, err := h.userService.UpdateCooperationStartDate(c.Request().Context(), id, req.CooperationStartDate)
	if err != nil {
		if err == contract.ErrUserNotFound {
			return dto.ErrorJSON(c, http.StatusNotFound, "user not found")
		}
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, UpdateUserResponse{
		Data: UserData{
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
		},
	})
}

// SearchUsers handles searching users by term
// @Summary Search users
// @Description Search users by full name, email, or phone
// @Tags users - admin
// @Accept json
// @Produce json
// @Param term path string true "Search term"
// @Success 200 {object} SearchUsersResponse
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/admin/users/search/{term} [get]
func (h *UserHandler) SearchUsers(c echo.Context) error {
	term := c.Param("term")
	if term == "" {
		return dto.BadRequestJSON(c, "search term is required")
	}

	users, err := h.userService.SearchUsers(c.Request().Context(), term)
	if err != nil {
		return dto.InternalServerErrorJSON(c, "failed to search users")
	}

	return dto.SuccessJSON(c, SearchUsersResponse{
		Items: users,
		Total: len(users),
	})
}
