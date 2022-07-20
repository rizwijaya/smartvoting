package view

import (
	"smartvoting/modules/v1/utilities/voting/repository"
	"smartvoting/modules/v1/utilities/voting/service"

	"gorm.io/gorm"
)

type votingView struct {
	votingService service.Service
}

func NewVotingView(votingService service.Service) *votingView {
	return &votingView{votingService}
}

func View(db *gorm.DB) *votingView {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	View := NewVotingView(Service)
	return View
}
