package server

import (
	"go_project/internal/logger"
	"go_project/internal/user"

	_ "go_project/docs" // swag docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// Server holds dependencies
type Server struct {
	UserService user.UserService
}

// RegisterRoutes registers all Gin routes
func (s *Server) RegisterRoutes(r *gin.Engine) {
	logger.Log.Info().Msg("Registering routes")

	r.GET("/users", s.GetUsersHandler)
	r.POST("/users", s.AddUserHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func SetupServer(dbConn *gorm.DB) *gin.Engine {
	userRepo := user.NewRepository(dbConn)
	userService := user.NewService(userRepo)
	srv := Server{UserService: userService}

	r := gin.New()
	r.Use(gin.Recovery(), logger.JsonLoggerMiddleware())
	srv.RegisterRoutes(r)
	return r
}
