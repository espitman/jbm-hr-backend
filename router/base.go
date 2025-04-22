package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// registerBaseRoutes registers all base routes
func (r *Router) registerBaseRoutes() {
	r.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Welcome to Echo API",
		})
	})
}
