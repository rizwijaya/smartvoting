package user

import (
	ss "smartvoting/modules/v1/utilities/user/service"
)

type UserHandler interface {
}

type userHandler struct {
	userService ss.Service
}

func NewUserHandler(userService ss.Service) *userHandler {
	return &userHandler{userService}
}
