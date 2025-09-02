package controllers

import (
	"kkl-v2/data/services"
)

type UserCtrl interface {
	//Methods
}

type userCtrl struct {
	userSvc services.UserService
}

//Implement controllers

func NewUserCtrl(userSvc services.UserService) UserCtrl {
	return &userCtrl{
		userSvc: userSvc,
	}
}
