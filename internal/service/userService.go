package service

import (
	"userProfile/internal/models"
	"userProfile/internal/repository"

)

type UserServiceType struct {
	Repo *repository.UserRepositoryType
}

func UserService(repo *repository.UserRepositoryType) *UserServiceType {
	return &UserServiceType{Repo: repo}
}

func (s *UserServiceType) CreateUser(user models.User) error {
	return s.Repo.Create(user)
}

func (s *UserServiceType) LoginUser(username string) (*models.User, error) {
	user, err := s.Repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
