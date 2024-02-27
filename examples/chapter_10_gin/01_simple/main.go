package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := router.Run(":8080"); err != nil {
	}
}
