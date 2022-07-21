package voting

import (
	"log"
	"net/http"
	"smartvoting/app/blockchain"
	"smartvoting/modules/v1/utilities/voting/models"
	"smartvoting/modules/v1/utilities/voting/service"

	"github.com/gin-gonic/gin"
)

type VotingHandler interface {
	CreateElection(c *gin.Context)
}

type votingHandler struct {
	votingService service.Service
}

func NewVotingHandler(votingService service.Service) *votingHandler {
	return &votingHandler{votingService}
}

func (h *votingHandler) CreateElection(c *gin.Context) {
	var input models.NewElection

	err := c.ShouldBindJSON(&input)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusConflict, "create_election.html", gin.H{
			"title":   "Buat Event Pemilihan - Smart Voting",
			"message": "Gagal menambahkan event pemilihan",
		})
		return
	}

	secretKey, err := h.votingService.GetServerKey(input.Server)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusConflict, "create_election.html", gin.H{
			"title":   "Buat Event Pemilihan - Smart Voting",
			"message": "Gagal menambahkan event pemilihan",
		})
		return
	}

	auth := blockchain.GetAccountAuth(blockchain.Connect(), secretKey)
	_, err = h.votingService.CreateElection(input, auth)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusConflict, "create_election.html", gin.H{
			"title":   "Buat Event Pemilihan - Smart Voting",
			"message": "Gagal menambahkan event pemilihan",
		})
		return
	}
	c.HTML(http.StatusOK, "create_election.html", gin.H{
		"title":   "Buat Event Pemilihan - Smart Voting",
		"message": "Berhasil menambahkan event pemilihan",
	})
}
