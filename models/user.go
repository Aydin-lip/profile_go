package models

import "gorm.io/gorm"

type User struct {
	// Id uint `json:"id" gorm:"primary"`
	Id        uint   `json:"id" gorm:"unique;not null;primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Phone     string `json:"phone" gorm:"unique"`
	Age       int    `json:"age"`
	gorm.Model
}
