package main

import (
	"basic-jwt-auth/handlers"
	"basic-jwt-auth/initializer"
	"basic-jwt-auth/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func main() {
	initializer.LoadEnvVariables()
	// Create a new Fiber instance
	app := fiber.New()
	// Use the recover middleware
	app.Use(recover.New())
	// Create a new JWT middleware
	// Note: This is just an example, please use a secure secret key
	jwt := middleware.NewAuthMiddleware(initializer.Secret)
	// Create a Login route
	app.Post("/api/login", handlers.Login)
	// Create a protected route
	app.Get("/api/protected", jwt, handlers.Protected)
	// Listen on port 3000
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
