package handlers

import (
	"basic-jwt-auth/dal"
	"basic-jwt-auth/initializer"
	"basic-jwt-auth/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// Login Маршрут входу
func Login(c *fiber.Ctx) error {
	// Витягніть облікові дані з тіла запиту
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Знайдіть користувача за обліковими даними
	user, err := dal.FindByCredentials(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	day := time.Hour * 24
	// Створіть вимоги JWT, які включають ID користувача та час закінчення дії
	claims := jwt.MapClaims{
		"sub": user.ID,
		"fav": user.NickName,
		"exp": time.Now().Add(day * 1).Unix(),
	}
	// Створіть токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Згенеруйте закодований токен та надішліть його у відповідь.
	t, err := token.SignedString([]byte(initializer.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Поверніть токен
	return c.JSON(models.LoginResponse{
		Token: t,
	})
}

// Protected
// Захищений маршрут
func Protected(c *fiber.Ctx) error {
	// Отримайте користувача з контексту та поверніть його
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	favPhrase := claims["fav"].(string)
	id := int(claims["sub"].(float64))
	return c.SendString(fmt.Sprintf("Вітаємо 👋 %s (#id: %d)", favPhrase, id))
}
