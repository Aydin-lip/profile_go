package main

import (
	"log"
	"os"

	"example.com/db"
	"example.com/models"
	"github.com/gin-gonic/gin"
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

	db := db.MySqlDB()
	// defer db.Close()

	models.SetupModels(db)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		var users []models.User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch users"})
			return
		}
		c.JSON(200, gin.H{
			"users": users,
		})
	})
	// r.POST("/user", \addUserHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	r.Run(port)
}
