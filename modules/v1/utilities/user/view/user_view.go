package view

import (
	"fmt"
	"log"
	"net/http"
	api "smartvoting/app/contracts"
	"smartvoting/modules/v1/utilities/user/repository"
	"smartvoting/modules/v1/utilities/user/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userView struct {
	userService service.Service
}

func NewUserView(userService service.Service) *userView {
	return &userView{userService}
}

func View(db *gorm.DB, blockchain *api.Api) *userView {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository, blockchain)
	View := NewUserView(Service)
	return View
}

func (h *userView) Index(c *gin.Context) {
	title := "Smart Voting"
	c.HTML(http.StatusOK, "dashboard.html", gin.H{"title": title})
}

func (h *userView) Login(c *gin.Context) {
	title := "Login Smart Voting"
	c.HTML(http.StatusOK, "login.html", gin.H{"title": title})
}

func (h *userView) Dashboard(c *gin.Context) {
	session := sessions.Default(c)
	address, err := h.userService.GetUserAddress(fmt.Sprintf("%v", session.Get("userID")))
	if err != nil {
		log.Println(err)
		return
	}
	userData, _ := h.userService.GetUserData(address)

	title := "Dashboard Smart Voting"
	c.HTML(http.StatusOK, "beranda.html", gin.H{
		"title": title,
		"user":  userData,
		"role":  session.Get("role"),
	})
}
