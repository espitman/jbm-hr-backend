package router

import (
	"net/http"

	_ "github.com/espitman/jbm-hr-backend/docs" // This is important for Swagger
	"github.com/espitman/jbm-hr-backend/http/handlers/albumhandler"
	"github.com/espitman/jbm-hr-backend/http/handlers/departmenthandler"
	"github.com/espitman/jbm-hr-backend/http/handlers/hrteamhandler"
	"github.com/espitman/jbm-hr-backend/http/handlers/uihandler"
	"github.com/espitman/jbm-hr-backend/http/handlers/uploadhandler"
	"github.com/espitman/jbm-hr-backend/http/handlers/userhandler"
	customMiddleware "github.com/espitman/jbm-hr-backend/http/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Router holds the echo instance and handlers
type Router struct {
	*echo.Echo
	albumHandler           *albumhandler.AlbumHandler
	albumAdminHandler      *albumhandler.AlbumAdminHandler
	userHandler            *userhandler.UserHandler
	uiHandler              *uihandler.UIHandler
	departmentHandler      *departmenthandler.DepartmentHandler
	departmentAdminHandler *departmenthandler.DepartmentAdminHandler
	hrTeamHandler          *hrteamhandler.HRTeamHandler
	hrTeamAdminHandler     *hrteamhandler.HRTeamAdminHandler
	uploadHandler          *uploadhandler.UploadHandler
}

// NewRouter creates a new router instance
func NewRouter(
	albumHandler *albumhandler.AlbumHandler,
	albumAdminHandler *albumhandler.AlbumAdminHandler,
	userHandler *userhandler.UserHandler,
	departmentHandler *departmenthandler.DepartmentHandler,
	departmentAdminHandler *departmenthandler.DepartmentAdminHandler,
	hrTeamHandler *hrteamhandler.HRTeamHandler,
	hrTeamAdminHandler *hrteamhandler.HRTeamAdminHandler,
	uiHandler *uihandler.UIHandler,
	uploadHandler *uploadhandler.UploadHandler,
) *Router {
	e := echo.New()
	e.Use(customMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORS())

	return &Router{
		Echo:                   e,
		albumHandler:           albumHandler,
		albumAdminHandler:      albumAdminHandler,
		userHandler:            userHandler,
		uiHandler:              uiHandler,
		departmentHandler:      departmentHandler,
		departmentAdminHandler: departmentAdminHandler,
		hrTeamHandler:          hrTeamHandler,
		hrTeamAdminHandler:     hrTeamAdminHandler,
		uploadHandler:          uploadHandler,
	}
}

// SetupRoutes sets up all the routes in the application
func (r *Router) SetupRoutes() {
	// Serve UI static files first
	r.GET("/*", r.uiHandler.ServeUI)

	// Create API v1 group
	apiV1 := r.Group("/api/v1")

	// Admin API v1 group with admin middleware
	apiV1Admin := apiV1.Group("/admin")
	apiV1Admin.Use(customMiddleware.JWT())
	apiV1Admin.Use(customMiddleware.Admin())

	// Register routes
	r.registerAlbumRoutes(apiV1)
	r.registerAlbumAdminRoutes(apiV1Admin)
	r.registerUserRoutes(apiV1)
	r.registerUserAdminRoutes(apiV1Admin)
	r.registerDepartmentRoutes(apiV1)
	r.registerDepartmentAdminRoutes(apiV1Admin)
	r.registerHRTeamRoutes(apiV1)
	r.registerHRTeamAdminRoutes(apiV1Admin)
	r.registerUploadRoutes(apiV1)

	// Add Swagger
	r.GET("/swagger/*", echoSwagger.WrapHandler)
}

// GetEcho returns the echo instance
func (r *Router) GetEcho() *echo.Echo {
	return r.Echo
}

// ServeHTTP implements the http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Echo.ServeHTTP(w, req)
}

// registerUploadRoutes registers the upload routes
func (r *Router) registerUploadRoutes(api *echo.Group) {
	upload := api.Group("/upload")
	upload.POST("/image", r.uploadHandler.UploadImage)
}
