package voting

import (
	"smartvoting/modules/v1/utilities/voting/repository"
	"smartvoting/modules/v1/utilities/voting/service"

	"gorm.io/gorm"
)

type VotingHandler interface {
}

type votingHandler struct {
	votingService service.Service
}

func NewVotingHandler(votingService service.Service) *votingHandler {
	return &votingHandler{votingService}
}

func Handler(db *gorm.DB) *votingHandler {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	Handler := NewVotingHandler(Service)
	return Handler
}
