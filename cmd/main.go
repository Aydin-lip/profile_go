package main

import (
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

func init() {
	config.LoadEnv()
}

func main() {
	db := database.SqlServerDB()

	models.SetupModels(db)
	validation.SetupCustom()

	r := routes.SetupRouter(db)

	port := config.GetEnv("PORT", ":8080")
	r.Run(port)
}
