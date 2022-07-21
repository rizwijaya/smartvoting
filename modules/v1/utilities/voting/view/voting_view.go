package view

import (
	"net/http"
	"smartvoting/modules/v1/utilities/voting/repository"
	"smartvoting/modules/v1/utilities/voting/service"

	"github.com/gin-gonic/gin"
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

func (h *votingView) CreateElection(c *gin.Context) {
	title := "Buat Event Pemilihan - Smart Voting"
	c.HTML(http.StatusOK, "create_election.html", gin.H{"title": title})
}
