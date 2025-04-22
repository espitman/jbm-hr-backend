package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterBaseRoutes registers the base application routes
func RegisterBaseRoutes(r *gin.Engine) {
	// Define a simple route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Gin API",
		})
	})
}
