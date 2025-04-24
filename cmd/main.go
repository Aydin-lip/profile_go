package main

import (
	"userProfile/config"
	"userProfile/internal/database"
	"userProfile/internal/models"
	"userProfile/routes"
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

	r := routes.SetupRouter(db)
	port := config.GetEnv("PORT", ":8080")
	r.Run(port)
}
