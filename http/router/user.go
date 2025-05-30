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
		group.POST("/admin/login", r.userHandler.AdminLogin)

		// Protected routes (JWT required)
		protected := users.Group("")
		protected.Use(middleware.JWT())
		{
			protected.GET("/me", r.userHandler.GetMe)
			protected.PUT("/avatar", r.userHandler.UpdateAvatar)
			protected.PUT("/confirm", r.userHandler.ConfirmUser)
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
		users.GET("/:id", r.userHandler.GetUserByID)
		users.PUT("/:id", r.userHandler.UpdateUser)
		users.PUT("/:id/password", r.userHandler.UpdateUserPassword)
		users.PUT("/:id/avatar", r.userHandler.UpdateUserAvatar)
		users.PUT("/:id/birthdate", r.userHandler.UpdateUserBirthdate)
		users.PUT("/:id/cooperation-start-date", r.userHandler.UpdateUserCooperationStartDate)
		users.PUT("/:id/active", r.userHandler.ActivateUser)
		users.PUT("/:id/deactivate", r.userHandler.DeactivateUser)
		users.GET("/search/:term", r.userHandler.SearchUsers)
		users.GET("/today-birthdate", r.userHandler.GetUsersWithTodayBirthdate)
		users.GET("/today-cooperation-start-date", r.userHandler.GetUsersWithTodayCooperationStartDate)
		users.GET("/jalali-month-birthdate", r.userHandler.GetUsersWithBirthdateInJalaliMonth, middleware.JWT(), middleware.Admin())
		users.GET("/jalali-month-cooperation-start-date", r.userHandler.GetUsersWithCooperationStartDateInJalaliMonth, middleware.JWT(), middleware.Admin())
	}
}
