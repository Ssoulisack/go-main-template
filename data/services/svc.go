package services

import (
	"kkl-v2/data/repositories"
)

type UserService interface {
}

type userService struct {
	userRepo repositories.UserRepository
}

//Implement services

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
