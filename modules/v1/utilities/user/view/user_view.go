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
	title := "Dashboard"
	c.HTML(http.StatusOK, "dashboard.html", gin.H{"title": title})
}
