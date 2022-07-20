package user

import (
	"log"
	"net/http"
	"smartvoting/modules/v1/utilities/user/models"
	ss "smartvoting/modules/v1/utilities/user/service"
	token "smartvoting/pkg/jwt"
	notify "smartvoting/pkg/notification"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Login(c *gin.Context)
}

type userHandler struct {
	userService ss.Service
}

func NewUserHandler(userService ss.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Login(c *gin.Context) {
	session := sessions.Default(c)
	var input models.LoginInput
	err := c.ShouldBind(&input)
	if err != nil {
		log.Println(err) //Print log error
		session.AddFlash("Token yang anda masukkan salah!")
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title":   "Login Smart Voting",
			"flashes": session.Flashes(),
		})
		return
	}

	//Verifikasi Login
	user, err := h.userService.Login(input)
	if err != nil {
		log.Println(err) //Print log error
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title":   "Login Smart Voting",
			"message": "Token yang anda masukkan Salah!",
		})
		return
	}

	//Generate JWT Token (Please Uncomment for Generate Token)
	jwt, err := token.GenerateToken(user.User_id)
	if err != nil {
		log.Println(err) //Cetak log error
		return
	}

	fm := []byte("Berhasil Login")
	notify.SetMessage(c.Writer, "message", fm)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:  "token",
		Value: jwt,
	})

	//Create Session Login
	session.Set("userID", user.User_id)
	session.Set("userToken", user.Token)
	session.Set("userRole", user.Role_id)
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}
