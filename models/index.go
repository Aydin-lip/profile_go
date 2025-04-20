package models

import "gorm.io/gorm"

func SetupModels(db *gorm.DB) {
	if err := db.AutoMigrate(&User{}); err != nil {
		panic("Failed to migrate database")
	}

}
