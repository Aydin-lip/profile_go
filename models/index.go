package models

import (
	"userProfile/schema"

	"gorm.io/gorm"
)

// type schema struct {
// 	name string
// 	db   *gorm.DB
// }

func SetupModels(db *gorm.DB) {
	schema.Create(db)

	if err := db.AutoMigrate(&User{}); err != nil {
		panic("Failed to migrate database")
	}

	// u1 := User{
	// 	FirstName: "Aydin",
	// 	LastName:  "Azakh",
	// 	Username:  "Aydin.lip",
	// 	Password:  "1234",
	// 	Email:     "ms1166760@gmail.com",
	// 	Phone:     "09037336131",
	// 	Age:       20,
	// }

	// result := db.Create(u1)

	// fmt.Println(result)
}
