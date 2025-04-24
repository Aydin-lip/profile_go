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
		// Look for SQL Server UNIQUE constraint errors
		errMsg := err.Error()

		switch {
		case strings.Contains(errMsg, "uni_Security_Users_username"):
			return errors.New("نام کاربری قبلاً گرفته شده است")
		case strings.Contains(errMsg, "uni_Security_Users_email"):
			return errors.New("ایمیل قبلا ثبت شده است")
		case strings.Contains(errMsg, "uni_Security_Users_phone"):
			return errors.New("شماره تلفن قبلا ثبت شده است")
		}

		return err
	}
	return nil
}

func (r *UserRepositoryType) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("userName = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
