package main

import (
	"fmt"
	"log"
	"net/http"
	"resfulsimple/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", handlers.ProductHandler)

	fmt.Println("Сервер запущено на http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
