package main

import (
	"gin/config"
	"gin/controllers"
	"gin/middleware"
	"gin/repositories"
	"gin/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                    = config.SetupDatabaseConnection()
	userRepository repositories.UserRepository = repositories.NewUserRepository(db)
	jwtService     services.JWTService         = services.NewJWTService()
	userService    services.UserService        = services.NewUserService(userRepository)
	authService    services.AuthService        = services.NewAuthService(userRepository)
	authController controllers.AuthController  = controllers.NewAuthController(authService, jwtService)
	userController controllers.UserController  = controllers.NewUserController(userService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	authRoutes := r.Group("api/auth" /*, middleware.AuthorizeJWT(jwtService)*/)
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	r.Run()
}
