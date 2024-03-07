package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

func main() {
	r := gin.Default()
	r.GET("/hello", helloHandler)
	r.Run() // Слухаємо на 0.0.0.0:8080
}
