package utils

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetQueryParamString gets a string query parameter from the Echo context
func GetQueryParamString(c echo.Context, key string) *string {
	value := c.QueryParam(key)
	if value == "" {
		return nil
	}
	return &value
}

// GetQueryParamInt gets an integer query parameter with a default value
func GetQueryParamInt(c echo.Context, param string, defaultValue int) int {
	value := c.QueryParam(param)
	if value == "" {
		return defaultValue
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return result
}
