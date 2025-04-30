package router

import (
	"github.com/espitman/jbm-hr-backend/http/middleware"
	"github.com/labstack/echo/v4"
)

// registerResumeRoutes registers all resume-related routes
func (r *Router) registerResumeRoutes(apiV1 *echo.Group) {
	// User routes
	resume := apiV1.Group("/resumes")
	resume.Use(middleware.JWT())
	resume.POST("", r.resumeHandler.Create)
}

// registerResumeAdminRoutes registers all admin resume-related routes
func (r *Router) registerResumeAdminRoutes(apiV1Admin *echo.Group) {
	// Admin routes
	resumes := apiV1Admin.Group("/resumes")
	resumes.GET("", r.resumeAdminHandler.List)
	resumes.GET("/:id", r.resumeAdminHandler.Get)
	resumes.PUT("/:id/status", r.resumeAdminHandler.UpdateStatus)
}
