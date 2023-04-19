package routes

import (
	"github.com/cheenusrivel/nft-go-api/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/", controller.GetUsers)
	router.POST("/", controller.CreateUser)
	router.DELETE("/:nric", controller.DeleteUser)
	router.PUT("/:nric", controller.UpdateUser)
}
