package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"userProfile/internal/controller"
	"userProfile/internal/repository"
	"userProfile/internal/service"
	// "userProfile/middleware"
)

func SecurityRoute(route *gin.RouterGroup, db *gorm.DB) {
	userRepo := repository.UserRepository(db)
	userService := service.UserService(userRepo)
	userController := controller.UserController(userService)

	route.POST("/Register", userController.Register)
	route.POST("/Login", userController.Login)
}
