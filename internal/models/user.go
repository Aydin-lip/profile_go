package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"userProfile/internal/schema"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName" gorm:"unique;not null" binding:"required,username"`
	Password  string `json:"password" gorm:"not null" binding:"required,gte=8"`
	Email     string `json:"email" gorm:"unique;not null" binding:"omitempty,email"`
	Phone     string `json:"phone" gorm:"unique" binding:"omitempty,e164"`
	Age       int    `json:"age"`
	gorm.Model
}

type UserResponse struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	UserName  string    `json:"userName"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"createdAt"`
}

func (User) TableName() string {
	return fmt.Sprintf("%s.Users", schema.Security())
}
