package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"userProfile/config"
	"userProfile/internal/database"
	"userProfile/internal/models"
	"userProfile/routes"
	"userProfile/validation"
)

type person struct {
	firstName string
	lastName  string
	age       int
	phone     string
	email     string
}

func main() {
	config.LoadEnv()

	db := database.SqlServerDB()
	// defer db.Close()

	models.SetupModels(db)
	validation.SetupCustom()

	r := routes.SetupRouter(db)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	port := config.GetEnv("PORT", ":8080")
	r.Run(port)
}
