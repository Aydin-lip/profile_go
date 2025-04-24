package repository

import (
	"gorm.io/gorm"

	"userProfile/internal/models"
)

type UserRepositoryType struct {
	DB *gorm.DB
}

func UserRepository(db *gorm.DB) *UserRepositoryType {
	return &UserRepositoryType{DB: db}
}

func (r *UserRepositoryType) Create(user models.User) error {
	return r.DB.Create(&user).Error
}
