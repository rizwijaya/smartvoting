package user

import (
	"smartvoting/modules/v1/utilities/user/repository"
	"smartvoting/modules/v1/utilities/user/service"

	"gorm.io/gorm"
)

func Handler(db *gorm.DB) *userHandler {
	userRepository := repository.NewRepository(db)
	userService := service.NewService(userRepository)
	userHandler := NewUserHandler(userService)
	return userHandler
}
