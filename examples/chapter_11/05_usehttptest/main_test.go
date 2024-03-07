package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/product", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(productHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Неправильний HTTP статус")

	var product Product
	err = json.Unmarshal(rr.Body.Bytes(), &product)
	assert.NoError(t, err)

	assert.Equal(t, 1, product.ID)
	assert.Equal(t, "Кава", product.Name)
	// httptest.NewServer(handler)
}
