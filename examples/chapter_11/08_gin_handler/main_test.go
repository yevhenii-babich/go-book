package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	// Встановлюємо режим тестування Gin
	gin.SetMode(gin.TestMode)

	// Ініціалізуємо маршрутизатор Gin
	r := gin.Default()
	r.GET("/hello", helloHandler)

	// Створюємо запит
	req, _ := http.NewRequest("GET", "/hello", nil)
	resp := httptest.NewRecorder()

	// Відправляємо запит до маршрутизатора Gin
	r.ServeHTTP(resp, req)

	// Перевіряємо результати
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "Hello, World!", resp.Body.String())
}
