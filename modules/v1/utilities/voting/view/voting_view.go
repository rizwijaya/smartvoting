package view

import (
	"net/http"
	api "smartvoting/app/contracts"
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

func View(db *gorm.DB, blockchain *api.Api) *votingView {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository, blockchain)
	View := NewVotingView(Service)
	return View
}

func (h *votingView) CreateElection(c *gin.Context) {
	title := "Buat Event Pemilihan - Smart Voting"
	status := 2                     //get server status nonAktif
	event := " AND event = \"nul\"" //belum pernah dipakai
	server, err := h.votingService.GetServer(status, event)
	if err != nil {
		c.HTML(http.StatusNoContent, "create_election.html", gin.H{
			"title":   title,
			"server":  server,
			"message": "server not found",
		})
		return
	}

	c.HTML(http.StatusOK, "create_election.html", gin.H{
		"title":  title,
		"server": server,
	})
}
