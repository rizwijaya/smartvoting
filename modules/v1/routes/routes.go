package routes

import (
	"smartvoting/app/config"
	userHandlerV1 "smartvoting/modules/v1/utilities/user/handler"
	userViewV1 "smartvoting/modules/v1/utilities/user/view"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParseTmpl(router *gin.Engine) *gin.Engine { //Load HTML Template
	router.Static("/assets", "./public/assets")
	router.Static("/images", "./public/images")
	router.Static("/css", "./public/assets/css")
	router.Static("/js", "./public/assets/js")
	router.Static("/fonts", "./public/assets/fonts")
	return router
}

func Init(db *gorm.DB, conf config.Conf, router *gin.Engine) *gin.Engine {
	// votingHandlerV1 := votingHandlerV1.Handler(db)
	// votingViewV1 := votingviewV1.View(db)
	userHandlerV1 := userHandlerV1.Handler(db)
	userViewV1 := userViewV1.View(db)
	// Routing Website Service
	general := router.Group("/")
	general.GET("/", userViewV1.Index)
	general.GET("/dashboard", userViewV1.Dashboard)
	general.GET("/login", userViewV1.Login)
	general.POST("/login", userHandlerV1.Login)
	general.GET("/logout", userHandlerV1.Logout)
	//Routing API Service
	//api := router.Group("/api/v1")
	//api.POST("/newtoken", userHandlerV1.NewToken)

	router = ParseTmpl(router)
	return router
}
