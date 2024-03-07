package emulateapi

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func fakeAPIHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: (%s) %s %s", r.Proto, r.Method, r.URL.Path)
	product := Product{ID: 1, Name: "Кава"}
	_ = json.NewEncoder(w).Encode(product)
}

func TestExternalAPI(t *testing.T) {
	// Створення фейкового сервера
	server := httptest.NewTLSServer(http.HandlerFunc(fakeAPIHandler))
	defer server.Close()

	// URL нашого фейкового сервера
	url := server.URL

	// Виконання запиту до фейкового API
	client := server.Client()
	log.Printf("Request: GET %s", url)
	resp, err := client.Get(url)
	assert.NoError(t, err)
	defer resp.Body.Close()

	var product Product
	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&product))

	// Перевірка відповіді
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, 1, product.ID)
	assert.Equal(t, "Кава", product.Name)
}
