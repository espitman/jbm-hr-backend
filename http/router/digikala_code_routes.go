package router

import (
	"github.com/labstack/echo/v4"
)

// registerDigikalaCodeAdminRoutes registers the Digikala code admin routes
func (r *Router) registerDigikalaCodeAdminRoutes(admin *echo.Group) {
	digikalaCodes := admin.Group("/digikala-codes")
	digikalaCodes.POST("", r.digikalaCodeAdminHandler.Create)
	digikalaCodes.GET("", r.digikalaCodeAdminHandler.List)
	digikalaCodes.GET("/:id", r.digikalaCodeAdminHandler.Get)
	digikalaCodes.PUT("/:id/assign", r.digikalaCodeAdminHandler.Assign)
}
