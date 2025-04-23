package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

// Logger is a middleware that logs the request method, path, status code, and latency
func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Start timer
			start := time.Now()

			// Process request
			err := next(c)

			// Calculate latency
			latency := time.Since(start)

			// Log request details
			fmt.Printf("[%s] %s %s %d %s\n",
				time.Now().Format(time.RFC3339),
				c.Request().Method,
				c.Request().URL.Path,
				c.Response().Status,
				latency,
			)

			return err
		}
	}
}
