package main

import (
	"github.com/gin-gonic/gin"
	"jwt_gorm/controllers"
	"jwt_gorm/initializers"
	"jwt_gorm/middleware"
	"log"
)

func main() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	if err := initializers.SyncDatabase(); err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	authorized := r.Group("/", middleware.RequireAuth)
	authorized.GET("/validate", controllers.Validate)
	authorized.POST("/profile", controllers.CreateProfile)
	authorized.GET("/profile", controllers.GetProfile)
	authorized.PUT("/profile", controllers.UpdateProfile)
	authorized.GET("/profile/:id", controllers.GetProfileByID)  // This is an admin route
	authorized.GET("/all", controllers.AllProfiles)             // This is an admin route
	authorized.GET("/all/:profile_id", controllers.GetDataByID) // This is an admin route
	validated := authorized.Group("/validated")
	validated.GET("/validate", controllers.Validate) //GET http://validated/validate"
	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Println(err)
	}
	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}
