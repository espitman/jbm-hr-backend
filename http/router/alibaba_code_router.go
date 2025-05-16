package router

import (
	"github.com/espitman/jbm-hr-backend/http/handlers/alibabacodehandler"
	"github.com/espitman/jbm-hr-backend/http/middleware"
	"github.com/labstack/echo/v4"
)

// RegisterAlibabaCodeRoutes registers all Alibaba code routes
func RegisterAlibabaCodeRoutes(e *echo.Echo, alibabaCodeAdminHandler *alibabacodehandler.AlibabaCodeAdminHandler) {
	// Admin routes
	adminGroup := e.Group("/api/v1/admin/alibaba-codes")
	adminGroup.Use(middleware.Admin())
	adminGroup.POST("", alibabaCodeAdminHandler.Create)
	adminGroup.GET("", alibabaCodeAdminHandler.List)
	adminGroup.GET("/:id", alibabaCodeAdminHandler.Get)
	adminGroup.PUT("/:id/assign", alibabaCodeAdminHandler.Assign)
}
