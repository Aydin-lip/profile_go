package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"userProfile/internal/handler"
	"userProfile/internal/repository"
	"userProfile/internal/service"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	userRoutes := r.Group("/Security")
	{
		userRoutes.POST("/Register", userHandler.RegisterUser)
		// userRoutes.POST("/Login", userHandler.LoginUser)
	}

	return r
}
