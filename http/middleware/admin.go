package middleware

import (
	"net/http"

	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/utils"
	"github.com/labstack/echo/v4"
)

// Admin middleware ensures the user is authenticated and has admin role
func Admin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// First, apply JWT middleware to ensure user is authenticated
			if err := JWT()(func(c echo.Context) error {
				return nil
			})(c); err != nil {
				return err
			}

			// Get claims from context (set by JWT middleware)
			claims := c.Get("user").(*utils.Claims)

			// Verify admin role
			if claims.Role != "admin" {
				return dto.ErrorJSON(c, http.StatusForbidden, "Access denied. Admin role required.")
			}

			return next(c)
		}
	}
}
