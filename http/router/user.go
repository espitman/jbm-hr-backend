package router

import (
	"github.com/labstack/echo/v4"
)

// registerUserRoutes registers all user-related routes
func (r *Router) registerUserRoutes(group *echo.Group) {
	users := group.Group("/users")
	{
		users.POST("/request-otp", r.userHandler.RequestOTP)
		users.POST("/verify-otp", r.userHandler.VerifyOTP)
	}
}
