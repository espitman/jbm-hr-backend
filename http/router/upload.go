package router

import (
	"github.com/espitman/jbm-hr-backend/http/middleware"
	"github.com/labstack/echo/v4"
)

// registerUploadRoutes registers all upload-related routes
func (r *Router) registerUploadRoutes(group *echo.Group) {
	upload := group.Group("/upload")
	{
		// Protected routes (JWT required)
		upload.Use(middleware.JWT())
		{
			// Upload routes
			upload.POST("/image", r.uploadHandler.UploadImage)
			upload.POST("/document", r.uploadHandler.UploadDocument)

			// Pre-signed URL route
			upload.GET("/presigned/:key", r.uploadHandler.GetPresignedURL)
		}

		// Public routes (no JWT required)
		upload.POST("/image/public/:dir", r.uploadHandler.UploadPublicImage)
	}
}
