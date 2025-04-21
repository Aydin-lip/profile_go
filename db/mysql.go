package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySqlDB() *gorm.DB {
	dbUser := os.Getenv("MYSQL_DB_USERNAME")
	dbPass := os.Getenv("MYSQL_DB_PASSWORD")
	dbHost := os.Getenv("MYSQL_DB_HOST")
	dbPort := os.Getenv("MYSQL_DB_PORT")
	dbName := os.Getenv("MYSQL_DB_NAME")

	if (dbUser == "") ||
		(dbPass == "") ||
		(dbHost == "") ||
		(dbName == "") ||
		(dbPort == "") {
		log.Fatal("One or more required environment variables are not set")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(dsn)
		fmt.Println(err)
		panic("Failed to connect database")
	}

	return db
}
