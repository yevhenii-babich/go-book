package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	product := Product{ID: 1, Name: "Кава"}
	json.NewEncoder(w).Encode(product)
}

func main() {
	http.HandleFunc("/product", productHandler)
	http.ListenAndServe(":8080", nil)
}
