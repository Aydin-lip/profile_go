package models

import (
	"fmt"
	"gorm.io/gorm"
	"userProfile/internal/schema"
)

type User struct {
	// Id uint `json:"id" gorm:"primary"`
	// Id        uint   `json:"id" gorm:"unique;not null;primaryKey"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Phone     string `json:"phone" gorm:"unique"`
	Age       int    `json:"age"`
	gorm.Model
}

func (User) TableName() string {
	return fmt.Sprintf("%s.Users", schema.Security())
}
