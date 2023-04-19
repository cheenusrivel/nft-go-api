package main

import (
	"github.com/cheenusrivel/nft-go-api/config"
	"github.com/cheenusrivel/nft-go-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoutes(router)
	router.Run(":8080")
}
