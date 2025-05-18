package main

import (
	"context"
	"log"
	"net/http"
	"path/filepath"

	"github.com/espitman/jbm-hr-backend/database"
	"github.com/espitman/jbm-hr-backend/database/repository/album"
	"github.com/espitman/jbm-hr-backend/database/repository/alibaba_code"
	"github.com/espitman/jbm-hr-backend/database/repository/department"
	"github.com/espitman/jbm-hr-backend/database/repository/digikala_code"
	"github.com/espitman/jbm-hr-backend/database/repository/hrteam"
	"github.com/espitman/jbm-hr-backend/database/repository/otp"
	"github.com/espitman/jbm-hr-backend/database/repository/request"
	"github.com/espitman/jbm-hr-backend/database/repository/resume"
	"github.com/espitman/jbm-hr-backend/database/repository/user"
	"github.com/espitman/jbm-hr-backend/ent/migrate"
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
	"github.com/espitman/jbm-hr-backend/http/router"
	"github.com/espitman/jbm-hr-backend/service/albumservice"
	"github.com/espitman/jbm-hr-backend/service/alibabacodeservice"
	"github.com/espitman/jbm-hr-backend/service/departmentservice"
	"github.com/espitman/jbm-hr-backend/service/digikalacodeservice"
	"github.com/espitman/jbm-hr-backend/service/hrteamservice"
	"github.com/espitman/jbm-hr-backend/service/infoservice"
	"github.com/espitman/jbm-hr-backend/service/requestservice"
	"github.com/espitman/jbm-hr-backend/service/resumeservice"
	"github.com/espitman/jbm-hr-backend/service/uploadservice"
	"github.com/espitman/jbm-hr-backend/service/userservice"
	"github.com/espitman/jbm-hr-backend/utils/config"
	_ "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/files"
)

// @title           JBM HR Backend API
// @version         1.01
// @description     This is the backend API for JBM HR system.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  s.heidari@jabama.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Load environment variables
	config.LoadEnv()

	// Get port from environment variables
	port := config.GetConfig("PORT", "8080")

	// Initialize database connection
	client, err := database.NewClient()
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer client.Close()

	// Run database migrations
	err = client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Initialize repository
	albumRepo := album.NewEntRepository(client)
	userRepo := user.NewEntRepository(client)
	otpRepo := otp.NewEntRepository(client)
	departmentRepo := department.NewEntRepository(client)
	hrTeamRepo := hrteam.NewEntRepository(client)
	resumeRepo := resume.NewRepository(client)
	requestRepo := request.NewEntRepository(client)
	digikalaCodeRepo := digikala_code.NewEntRepository(client)
	alibabaCodeRepo := alibaba_code.NewEntRepository(client)

	// Initialize service
	albumService := albumservice.New(albumRepo)
	userService := userservice.New(userRepo, otpRepo)
	departmentService := departmentservice.New(departmentRepo)
	hrTeamService := hrteamservice.New(hrTeamRepo)
	resumeService := resumeservice.New(resumeRepo)
	requestService := requestservice.New(requestRepo)
	digikalaCodeService := digikalacodeservice.New(digikalaCodeRepo)
	alibabaCodeService := alibabacodeservice.NewService(alibabaCodeRepo)
	infoService := infoservice.NewInfoService(userRepo, requestRepo, resumeRepo, departmentRepo)

	// Initialize upload service
	uploadService, err := uploadservice.New()
	if err != nil {
		log.Fatalf("Failed to initialize upload service: %v", err)
	}

	// Initialize handlers
	albumHandler := albumhandler.NewAlbumHandler(albumService)
	albumAdminHandler := albumhandler.NewAlbumAdminHandler(albumService)
	userHandler := userhandler.NewUserHandler(userService)
	departmentHandler := departmenthandler.NewDepartmentHandler(departmentService)
	departmentAdminHandler := departmenthandler.NewDepartmentAdminHandler(departmentService)
	hrTeamHandler := hrteamhandler.NewHRTeamHandler(hrTeamService)
	hrTeamAdminHandler := hrteamhandler.NewHRTeamAdminHandler(hrTeamService)
	uploadHandler := uploadhandler.NewUploadHandler(uploadService)
	resumeHandler := resumehandler.NewResumeHandler(resumeService)
	resumeAdminHandler := resumehandler.NewResumeAdminHandler(resumeService)
	requestHandler := requesthandler.NewHandler(requestService)
	requestAdminHandler := requesthandler.NewAdminHandler(requestService)
	digikalaCodeAdminHandler := digikalacodehandler.NewDigikalaCodeAdminHandler(digikalaCodeService)
	alibabaCodeAdminHandler := alibabacodehandler.NewAlibabaCodeAdminHandler(alibabaCodeService)
	infoHandler := infohandler.NewInfoHandler(infoService)

	// Initialize UI handler
	webPath, _ := filepath.Abs("ui/web")
	adminPath, _ := filepath.Abs("ui/admin")
	uiHandler := uihandler.NewUIHandler(webPath, adminPath)

	// Initialize router
	r := router.NewRouter(
		albumHandler,
		albumAdminHandler,
		userHandler,
		departmentHandler,
		departmentAdminHandler,
		hrTeamHandler,
		hrTeamAdminHandler,
		uiHandler,
		uploadHandler,
		resumeHandler,
		resumeAdminHandler,
		requestHandler,
		requestAdminHandler,
		digikalaCodeAdminHandler,
		alibabaCodeAdminHandler,
		infoHandler,
	)
	r.SetupRoutes()

	// Start server
	log.Printf("Server starting on port %s...", port)
	if err := r.Start(":" + port); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed starting server: %v", err)
	}
}
