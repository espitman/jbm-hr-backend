package router

import (
	"github.com/espitman/jbm-hr-backend/http/middleware"
	"github.com/labstack/echo/v4"
)

// registerDepartmentRoutes registers all department-related routes
func (r *Router) registerDepartmentRoutes(group *echo.Group) {
	departments := group.Group("/departments")
	{
		// Protected routes (JWT required)
		departments.Use(middleware.JWT())
		{
			// Public routes (for authenticated users)
			departments.GET("", r.departmentHandler.List)
			departments.GET("/:id", r.departmentHandler.Get)
		}
	}
}

// registerDepartmentAdminRoutes registers all admin department-related routes
func (r *Router) registerDepartmentAdminRoutes(group *echo.Group) {
	departments := group.Group("/departments")
	{
		// Admin routes (protected by admin middleware)
		departments.POST("", r.departmentAdminHandler.Create)
		departments.PUT("/:id", r.departmentAdminHandler.Update)
		departments.DELETE("/:id", r.departmentAdminHandler.Delete)
	}
}
