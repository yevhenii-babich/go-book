package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Основний маршрут
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Головна сторінка")
	})

	// Маршрут з параметром
	app.Get("/greeting/:name", func(c *fiber.Ctx) error {
		name := c.Params("name") // Отримання параметра маршруту
		return c.SendString(fmt.Sprintf("Привіт, %s!", name))
	})

	// Маршрут з необов'язковим параметром
	app.Get("/hello/:name?", func(c *fiber.Ctx) error {
		name := c.Params("name", "Невідомий") // Значення за замовчуванням, якщо параметр відсутній
		return c.SendString(fmt.Sprintf("Привіт, %s!", name))
	})

	// Маршрут з декількома параметрами
	app.Get("/greet/:name/:age/:city", func(c *fiber.Ctx) error {
		name := c.Params("name")
		age := c.Params("age")
		city := c.Params("city")
		return c.SendString(fmt.Sprintf("%s з міста %s, тобі %s років", name, city, age))
	})

	// Запуск сервера
	log.Fatal(app.Listen(":3000"))
}
