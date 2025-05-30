package router

import (
	"net/http"

	_ "github.com/espitman/jbm-hr-backend/docs" // This is important for Swagger
	"github.com/espitman/jbm-hr-backend/http/handlers/albumhandler"
	"github.com/espitman/jbm-hr-backend/http/handlers/alibabacodehandler"
	"github.com/espitman/jbm-hr-backend/http/handlers/departmenthandler"
	"github.com/espitman/jbm-hr-backend/http/handlers/digikalacodehandler"
	"github.com/espitman/jbm-hr-backend/http/handlers/hrteamhandler"
	"github.com/espitman/jbm-hr-backend/http/handlers/infohandler"
	"github.com/espitman/jbm-hr-backend/http/handlers/requesthandler"
	"github.com/espitman/jbm-hr-backend/http/handlers/resumehandler"
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
	albumHandler             *albumhandler.AlbumHandler
	albumAdminHandler        *albumhandler.AlbumAdminHandler
	userHandler              *userhandler.UserHandler
	uiHandler                *uihandler.UIHandler
	departmentHandler        *departmenthandler.DepartmentHandler
	departmentAdminHandler   *departmenthandler.DepartmentAdminHandler
	hrTeamHandler            *hrteamhandler.HRTeamHandler
	hrTeamAdminHandler       *hrteamhandler.HRTeamAdminHandler
	uploadHandler            *uploadhandler.UploadHandler
	resumeHandler            *resumehandler.ResumeHandler
	resumeAdminHandler       *resumehandler.ResumeAdminHandler
	requestHandler           *requesthandler.Handler
	requestAdminHandler      *requesthandler.AdminHandler
	digikalaCodeAdminHandler *digikalacodehandler.DigikalaCodeAdminHandler
	alibabaCodeAdminHandler  *alibabacodehandler.AlibabaCodeAdminHandler
	infoHandler              *infohandler.InfoHandler
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
	resumeHandler *resumehandler.ResumeHandler,
	resumeAdminHandler *resumehandler.ResumeAdminHandler,
	requestHandler *requesthandler.Handler,
	requestAdminHandler *requesthandler.AdminHandler,
	digikalaCodeAdminHandler *digikalacodehandler.DigikalaCodeAdminHandler,
	alibabaCodeAdminHandler *alibabacodehandler.AlibabaCodeAdminHandler,
	infoHandler *infohandler.InfoHandler,
) *Router {
	e := echo.New()
	e.Use(customMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORS())

	return &Router{
		Echo:                     e,
		albumHandler:             albumHandler,
		albumAdminHandler:        albumAdminHandler,
		userHandler:              userHandler,
		uiHandler:                uiHandler,
		departmentHandler:        departmentHandler,
		departmentAdminHandler:   departmentAdminHandler,
		hrTeamHandler:            hrTeamHandler,
		hrTeamAdminHandler:       hrTeamAdminHandler,
		uploadHandler:            uploadHandler,
		resumeHandler:            resumeHandler,
		resumeAdminHandler:       resumeAdminHandler,
		requestHandler:           requestHandler,
		requestAdminHandler:      requestAdminHandler,
		digikalaCodeAdminHandler: digikalaCodeAdminHandler,
		alibabaCodeAdminHandler:  alibabaCodeAdminHandler,
		infoHandler:              infoHandler,
	}
}

// SetupRoutes sets up all the routes in the application
func (r *Router) SetupRoutes() {
	// Serve admin UI static files first
	r.GET("/admin", r.uiHandler.ServeAdminUI)
	r.GET("/admin/*", r.uiHandler.ServeAdminUI)

	// Serve web UI static files
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
	r.registerResumeRoutes(apiV1)
	r.registerResumeAdminRoutes(apiV1Admin)
	r.registerRequestRoutes(apiV1)
	r.registerRequestAdminRoutes(apiV1Admin)
	r.registerDigikalaCodeAdminRoutes(apiV1Admin)
	r.registerAlibabaCodeAdminRoutes(apiV1Admin)
	r.registerInfoAdminRoutes(apiV1Admin)

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

// registerAlibabaCodeAdminRoutes registers all Alibaba code admin routes
func (r *Router) registerAlibabaCodeAdminRoutes(apiV1Admin *echo.Group) {
	alibabaCodeAdminGroup := apiV1Admin.Group("/alibaba-codes")
	alibabaCodeAdminGroup.POST("", r.alibabaCodeAdminHandler.Create)
	alibabaCodeAdminGroup.GET("", r.alibabaCodeAdminHandler.List)
	alibabaCodeAdminGroup.GET("/:id", r.alibabaCodeAdminHandler.Get)
	alibabaCodeAdminGroup.PUT("/:id/assign", r.alibabaCodeAdminHandler.Assign)
}
