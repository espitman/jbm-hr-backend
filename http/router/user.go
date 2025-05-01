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
		users.PUT("/:id/password", r.userHandler.UpdatePassword)
		group.POST("/admin/login", r.userHandler.AdminLogin)

		// Protected routes (JWT required)
		protected := users.Group("")
		protected.Use(middleware.JWT())
		{
			protected.GET("/me", r.userHandler.GetMe)
		}
	}
}

// registerUserAdminRoutes registers all admin user-related routes
func (r *Router) registerUserAdminRoutes(group *echo.Group) {
	users := group.Group("/users")
	{
		// Admin routes (protected by admin middleware)
		users.POST("/register", r.userHandler.RegisterUser)
		users.GET("", r.userHandler.ListUsers)
		// Add other admin routes here
		// Example: users.GET("/all", r.userHandler.GetAllUsers)
	}
}
