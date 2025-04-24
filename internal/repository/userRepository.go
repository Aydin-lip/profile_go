package repository

import (
	"errors"
	"strings"

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
	if err := r.DB.Create(&user).Error; err != nil {
		// Check for unique constraint error
		if strings.Contains(err.Error(), "UNIQUE") || strings.Contains(err.Error(), "unique") {
			if strings.Contains(err.Error(), "username") {
				return errors.New("نام کاربری قبلاً گرفته شده است")
			}
			if strings.Contains(err.Error(), "email") {
				return errors.New("ایمیل قبلا ثبت شده است")
			}
			if strings.Contains(err.Error(), "phone") {
				return errors.New("شماره تلفن قبلا ثبت شده است")
			}
		}
		return err
	}
	return nil
}
