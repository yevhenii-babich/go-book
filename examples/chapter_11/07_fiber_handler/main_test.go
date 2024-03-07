package testfiberhander

import (
	"io"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Хендлер, який ми хочемо протестувати
func helloWorld(c *fiber.Ctx) error {
	time.Sleep(10 * time.Millisecond) // Штучна затримка для тесту
	return c.SendString("Hello, World!")
}

// Тестова функція
func TestHelloWorld(t *testing.T) {
	// Створюємо Fiber app
	app := fiber.New()

	// Реєструємо наш хендлер
	app.Get("/hello", helloWorld)

	// Створюємо httptest запит і відповідь
	req := httptest.NewRequest("GET", "/hello", nil)
	resp, err := app.Test(req, 11) // 11 мілісекунд - максимальний час виконання

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Читаємо відповідь і перевіряємо її
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, World!", string(body))
}
