package router

import (
	"github.com/espitman/jbm-hr-backend/http/middleware"
	"github.com/labstack/echo/v4"
)

// registerUserRoutes registers all user-related routes
func (r *Router) registerUserRoutes(group *echo.Group) {
	users := group.Group("/users")
	{
		// Public routes (no JWT required)
		users.POST("/request-otp", r.userHandler.RequestOTP)
		users.POST("/verify-otp", r.userHandler.VerifyOTP)

		// Protected routes (JWT required)
		protected := users.Group("")
		protected.Use(middleware.JWT())
		{
			// Add protected user routes here
			// Example: protected.GET("/profile", r.userHandler.GetProfile)
		}
	}
}

// registerUserAdminRoutes registers all admin user-related routes
func (r *Router) registerUserAdminRoutes(group *echo.Group) {
	users := group.Group("/users")
	{
		// Register endpoint is public (no JWT required)
		users.POST("/register", r.userHandler.RegisterUser)

		// Other admin routes require JWT
		adminProtected := users.Group("")
		adminProtected.Use(middleware.JWT())
		{
			// Add other admin routes here that require JWT
		}
	}
}
