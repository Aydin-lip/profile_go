package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"userProfile/internal/controller"
	"userProfile/internal/repository"
	"userProfile/internal/service"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userRepo := repository.UserRepository(db)
	userService := service.UserService(userRepo)
	userController := controller.UserController(userService)

	userRoutes := r.Group("/Security")
	{
		userRoutes.POST("/Register", userController.Register)
		// userRoutes.POST("/Login", userController.LoginUser)
	}

	return r
}
