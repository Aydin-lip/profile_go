package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userRoutes := r.Group("/Security")
	SecurityRoute(userRoutes, db)

	return r
}
