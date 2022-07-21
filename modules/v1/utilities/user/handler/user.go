package user

import (
	api "smartvoting/app/contracts"
	"smartvoting/modules/v1/utilities/user/repository"
	"smartvoting/modules/v1/utilities/user/service"

	"gorm.io/gorm"
)

func Handler(db *gorm.DB, blockchain *api.Api) *userHandler {
	userRepository := repository.NewRepository(db)
	userService := service.NewService(userRepository, blockchain)
	userHandler := NewUserHandler(userService)
	return userHandler
}
