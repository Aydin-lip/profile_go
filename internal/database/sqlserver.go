package database

import (
	"fmt"
	"log"

	"userProfile/config"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func SqlServerDB() *gorm.DB {
	// Load environment variables
	dbUser := config.GetEnv("SQLSERVER_DB_USERNAME", "")
	dbPass := config.GetEnv("SQLSERVER_DB_PASSWORD", "")
	dbHost := config.GetEnv("SQLSERVER_DB_HOST", "")
	dbPort := config.GetEnv("SQLSERVER_DB_PORT", "")
	dbName := config.GetEnv("SQLSERVER_DB_NAME", "")

	// Check if any of the required environment variables are empty
	if (dbUser == "") ||
		(dbPass == "") ||
		(dbHost == "") ||
		(dbName == "") ||
		(dbPort == "") {
		log.Fatal("One or more required environment variables are not set")
	}

	// Connect to the SQL Server
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(dsn)
		fmt.Println(err)
		panic("Failed to connect database (before create database)")
	}

	// Check if the database exists, and create it if it doesn't
	createDatabaseCommand := fmt.Sprintf("IF DB_ID('%s') IS NULL CREATE DATABASE [%s]", dbName, dbName)
	if err := db.Exec(createDatabaseCommand).Error; err != nil {
		fmt.Println("Error creating database:", err)
	}

	// Connect to the SQL Server database
	dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(dsn)
		fmt.Println(err)
		panic("Failed to connect database (after create database)")
	}

	return db
}
