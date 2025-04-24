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
	UserName  string `json:"userName" gorm:"unique;not null" binding:"required,username"`
	Password  string `json:"password" gorm:"not null" binding:"required,gte=8"`
	Email     string `json:"email" gorm:"unique;not null" binding:"omitempty,email"`
	Phone     string `json:"phone" gorm:"unique" binding:"omitempty,e164"`
	Age       int    `json:"age"`
	gorm.Model
}

func (User) TableName() string {
	return fmt.Sprintf("%s.Users", schema.Security())
}
