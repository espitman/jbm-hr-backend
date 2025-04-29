package router

import (
	"github.com/espitman/jbm-hr-backend/http/middleware"
	"github.com/labstack/echo/v4"
)

// registerHRTeamRoutes registers all HR team-related routes
func (r *Router) registerHRTeamRoutes(group *echo.Group) {
	hrTeam := group.Group("/hr-team")
	{
		// Protected routes (JWT required)
		hrTeam.Use(middleware.JWT())
		{
			// Public routes (for authenticated users)
			hrTeam.GET("", r.hrTeamHandler.List)
			hrTeam.GET("/:id", r.hrTeamHandler.Get)
		}
	}
}

// registerHRTeamAdminRoutes registers all admin HR team-related routes
func (r *Router) registerHRTeamAdminRoutes(group *echo.Group) {
	hrTeam := group.Group("/hr-team")
	{
		// Admin routes (protected by admin middleware)
		hrTeam.POST("", r.hrTeamAdminHandler.Create)
		hrTeam.PUT("/:id", r.hrTeamAdminHandler.Update)
		hrTeam.DELETE("/:id", r.hrTeamAdminHandler.Delete)
	}
}
