package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func SqlServerDB() *gorm.DB {
	dbUser := os.Getenv("SQLSERVER_DB_USERNAME")
	dbPass := os.Getenv("SQLSERVER_DB_PASSWORD")
	dbHost := os.Getenv("SQLSERVER_DB_HOST")
	dbPort := os.Getenv("SQLSERVER_DB_PORT")
	dbName := os.Getenv("SQLSERVER_DB_NAME")

	if (dbUser == "") ||
		(dbPass == "") ||
		(dbHost == "") ||
		(dbName == "") ||
		(dbPort == "") {
		log.Fatal("One or more required environment variables are not set")
	}

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(dsn)
		fmt.Println(err)
		panic("Failed to connect database")
	}

	createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", dbName)
	db.Exec(createDatabaseCommand)

	return db
}
