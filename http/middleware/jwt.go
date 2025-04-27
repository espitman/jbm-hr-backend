package middleware

import (
	"net/http"
	"strings"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/utils"
	"github.com/labstack/echo/v4"
)

// JWT middleware validates the JWT token in the Authorization header
func JWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get the Authorization header
			auth := c.Request().Header.Get("Authorization")
			if auth == "" {
				return dto.ErrorJSON(c, http.StatusUnauthorized, contract.ErrTokenMissing.Error())
			}

			// Check if the header has the Bearer prefix
			parts := strings.Split(auth, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return dto.ErrorJSON(c, http.StatusUnauthorized, contract.ErrTokenInvalid.Error())
			}

			// Validate the token
			claims, err := utils.ValidateToken(parts[1])
			if err != nil {
				if err == utils.ErrExpiredToken {
					return dto.ErrorJSON(c, http.StatusUnauthorized, contract.ErrTokenExpired.Error())
				}
				return dto.ErrorJSON(c, http.StatusUnauthorized, contract.ErrTokenInvalid.Error())
			}

			// Set the claims in the context for later use
			c.Set("user", claims)

			return next(c)
		}
	}
}
