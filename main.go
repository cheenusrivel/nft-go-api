package main

import (
	"github.com/cheenusrivel/nft-go-api/config"
	"github.com/cheenusrivel/nft-go-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use((CORSMiddleware()))
	config.Connect()
	routes.UserRoutes(router)
	router.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
