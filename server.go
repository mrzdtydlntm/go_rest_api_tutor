package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrzdtydlntm/golang_backend/config"
	"github.com/mrzdtydlntm/golang_backend/controller"
	"github.com/mrzdtydlntm/golang_backend/repository"
	"github.com/mrzdtydlntm/golang_backend/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	authRoutes := r.Group("api/auth") //gaperlu middleware
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	r.Run()
}
