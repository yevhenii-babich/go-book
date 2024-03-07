package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "net/http/pprof" //nolint:gosec
)

func helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

func main() {
	r := gin.Default()
	r.GET("/hello", helloHandler)
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil)) //nolint:gosec
	}()
	r.Run() // Слухаємо на 0.0.0.0:8080
}
