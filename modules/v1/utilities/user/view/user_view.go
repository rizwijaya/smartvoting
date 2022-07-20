package view

import (
	"net/http"
	"smartvoting/modules/v1/utilities/user/repository"
	"smartvoting/modules/v1/utilities/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userView struct {
	userService service.Service
}

func NewUserView(userService service.Service) *userView {
	return &userView{userService}
}

func View(db *gorm.DB) *userView {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
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
	title := "Dashboard Smart Voting"
	c.HTML(http.StatusOK, "beranda.html", gin.H{"title": title})
}
