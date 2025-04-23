package main

import (
	"log"
	"os"

	"userProfile/internal/database"
	"userProfile/internal/models"
	"userProfile/routes"

	"github.com/joho/godotenv"
)

type person struct {
	firstName string
	lastName  string
	age       int
	phone     string
	email     string
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// db := db.MySqlDB()
	db := database.SqlServerDB()
	// defer db.Close()

	models.SetupModels(db)

	// r := gin.Default()

	// r.GET("/Security")
	// // r.POST("/user", \addUserHandler)

	r := routes.SetupRouter(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	r.Run(port)
}
