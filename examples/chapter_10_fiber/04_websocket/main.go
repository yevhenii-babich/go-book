package main

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade повертає true, якщо клієнт
		// запросив оновлення до протоколу WebSocket.
		if websocket.IsWebSocketUpgrade(c) {
			// c.Locals додається allowed: true до *fiber.Ctx для використання в наступних обробниках
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		// c.Locals додається до *websocket.Conn
		log.Println("allowed:", c.Locals("allowed"))        // true
		log.Println("id:", c.Params("id"))                  // 123
		log.Println("v:", c.Query("v"))                     // 1.0
		log.Printf("session: [%s]\n", c.Cookies("session")) // ""

		// websocket.Conn зв'язки https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			mt  int
			msg []byte
			err error
		)
		for { // Нескінченний цикл для обробки повідомлень
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("читання:", err)
				break
			}
			log.Printf("отримано: %s", string(msg))
			msg = []byte("повернення: " + string(msg))
			// Повернення повідомлення
			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("запис:", err)
				break
			}
		}
		log.Printf("вихід\n")
	}))

	if err := app.Listen(":3000"); err != nil {
		log.Fatal()
	}
	// Доступ до сервера websocket: ws://localhost:3000/ws/123?v=1.0
}
