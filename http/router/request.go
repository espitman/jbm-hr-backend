package router

import (
	"github.com/espitman/jbm-hr-backend/http/middleware"
	"github.com/labstack/echo/v4"
)

// registerRequestRoutes registers all request-related routes
func (r *Router) registerRequestRoutes(group *echo.Group) {
	requests := group.Group("/requests")
	{
		// Protected routes (JWT required)
		requests.Use(middleware.JWT())
		{
			// User routes
			requests.POST("", r.requestHandler.CreateRequest)
		}
	}
}

// registerRequestAdminRoutes registers all admin request-related routes
func (r *Router) registerRequestAdminRoutes(group *echo.Group) {
	requests := group.Group("/requests")
	{
		// Admin routes (protected by admin middleware)
		requests.GET("", r.requestAdminHandler.GetRequests)
		requests.GET("/:id", r.requestAdminHandler.GetRequest)
		requests.PUT("/:id/status", r.requestAdminHandler.UpdateRequestStatus)
	}
}
