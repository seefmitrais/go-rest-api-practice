package service

import (
	"github.com/seefmitrais/go-rest-api-practice/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{repo: *userRepository}
}

func (us UserService) CreateNewUser(user *repository.User) error {
	if err := us.repo.Save(user); err != nil {
		return err
	}
	return nil
}
