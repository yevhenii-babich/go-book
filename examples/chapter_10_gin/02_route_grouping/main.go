package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Створення групи маршрутів для API
	apiRoutes := router.Group("/api")

	// Конфігурація BasicAuth Middleware
	authorized := apiRoutes.Group("/", gin.BasicAuth(gin.Accounts{
		"username1": "password1",
		"username2": "password2",
	}))

	// Маршрути для авторизованих користувачів
	authorized.GET("/secure-data", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ви маєте доступ до захищених даних!",
			"user":    c.MustGet(gin.AuthUserKey).(string),
		})
	})

	// Звичайний маршрут без автентифікації
	apiRoutes.GET("/open-data", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Це відкриті дані, доступні без автентифікації."})
	})

	router.Run(":8080")
}
