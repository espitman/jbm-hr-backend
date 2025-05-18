package router

import (
	"github.com/labstack/echo/v4"
)

// registerInfoAdminRoutes registers all info admin routes
func (r *Router) registerInfoAdminRoutes(group *echo.Group) {
	infoAdminGroup := group.Group("/info")
	infoAdminGroup.GET("/dashboard", r.infoHandler.GetDashboardInfo)
}
