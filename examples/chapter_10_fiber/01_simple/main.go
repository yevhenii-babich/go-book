package main

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	// Створення нового Fiber додатку
	app := fiber.New()
	// Використання логера
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	// Визначення GET маршруту
	app.Get("/", func(c *fiber.Ctx) error {
		// Відправлення відповіді
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})

	// Запуск сервера на порту 8080
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
