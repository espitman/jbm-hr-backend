package router

import (
	"github.com/gin-gonic/gin"
)

// registerBaseRoutes registers all base routes
func (r *Router) registerBaseRoutes() {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Gin API",
		})
	})
}
