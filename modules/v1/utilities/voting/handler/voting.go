package voting

import (
	api "smartvoting/app/contracts"
	"smartvoting/modules/v1/utilities/voting/repository"
	"smartvoting/modules/v1/utilities/voting/service"

	"gorm.io/gorm"
)

func Handler(db *gorm.DB, blockchain *api.Api) *votingHandler {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository, blockchain)
	Handler := NewVotingHandler(Service)
	return Handler
}
