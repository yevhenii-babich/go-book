package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Middleware для всього додатку
	app.Use(
		func(c *fiber.Ctx) error {
			// Логіка загального middleware
			c.Response().Header.Set("Access-Control-Allow-Origin", "*")
			return c.Next()
		},
		logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Головна сторінка")
	})
	// Група API маршрутів
	api := app.Group("/api/v1", func(c *fiber.Ctx) error {
		c.Request().Header.Set("X-Version-Id", "1")
		c.Response().Header.Set("X-Version-Id", "1")
		return c.Next()
	})

	// Додавання маршрутів до групи
	api.Get("/user", func(c *fiber.Ctx) error {
		// Обробка запиту
		return c.SendString("API User Endpoint")
	})

	api.Get("/product/:name?", func(c *fiber.Ctx) error {
		// Обробка запиту
		product := c.Params("name")
		if c.Params("name") == "" {
			product = "unknown"
		}
		return c.SendString(fmt.Sprintf("API Product %s Endpoint", product))
	})

	// Запуск сервера
	app.Listen(":3000")
}
