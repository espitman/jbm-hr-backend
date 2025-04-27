package userhandler

import (
	"net/http"

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
	// Get the user from the context (set by JWT middleware)
	userClaims, ok := c.Get("user").(map[string]interface{})
	if !ok {
		return dto.ErrorJSON(c, http.StatusUnauthorized, "User information not found")
	}

	// Check if the user has admin role
	role, ok := userClaims["role"].(string)
	if !ok || role != "admin" {
		return dto.ErrorJSON(c, http.StatusForbidden, "Only admin users can register new users")
	}

	var req RegisterUserRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	if err := utils.ValidateStruct(req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	user, err := h.userService.RegisterUser(c.Request().Context(), &contract.RegisterUserInput{
		Email:  req.Email,
		Phone:  req.Phone,
		Role:   req.Role,
		Avatar: req.Avatar,
	})
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	response := RegisterUserResponse{}
	response.Data = RegisterUserData{
		ID:     user.ID,
		Email:  user.Email,
		Phone:  user.Phone,
		Role:   user.Role,
		Avatar: user.Avatar,
	}

	return dto.CreatedJSON(c, response)
}
