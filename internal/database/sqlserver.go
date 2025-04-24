package database

import (
	"fmt"
	"log"

	"userProfile/config"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func SqlServerDB() *gorm.DB {
	dbUser := config.GetEnv("SQLSERVER_DB_USERNAME", "")
	dbPass := config.GetEnv("SQLSERVER_DB_PASSWORD", "")
	dbHost := config.GetEnv("SQLSERVER_DB_HOST", "")
	dbPort := config.GetEnv("SQLSERVER_DB_PORT", "")
	dbName := config.GetEnv("SQLSERVER_DB_NAME", "")

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
		panic("Failed to connect database (before create database)")
	}

	createDatabaseCommand := fmt.Sprintf("IF DB_ID('%s') IS NULL CREATE DATABASE [%s]", dbName, dbName)
	if err := db.Exec(createDatabaseCommand).Error; err != nil {
		fmt.Println("Error creating database:", err)
	}

	dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(dsn)
		fmt.Println(err)
		panic("Failed to connect database (after create database)")
	}

	return db
}
