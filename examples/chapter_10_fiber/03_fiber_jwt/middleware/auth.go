package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// NewAuthMiddleware Middleware JWT function
func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
		//ErrorHandler: errorHandler,
	})
}

// errorHandler custom Fiber error handler
func errorHandler(c *fiber.Ctx, err error) error {
	return c.SendStatus(401)
}
