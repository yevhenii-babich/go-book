package main

import (
	"log"
	"net/http"
	"routers/handlers"
)

func main() {
	mux := http.NewServeMux()
	// Налаштовуємо маршрутизацію
	// GET /
	mux.HandleFunc("/", handlers.DefaultRoute)
	// POST /greet
	mux.HandleFunc("/greet", handlers.PostGreet)
	// GET /help
	mux.HandleFunc("/help", handlers.GetHelp)
	// Виводимо інформаційне повідомлення та запускаємо сервер
	log.Println("Сервер запущено на http://localhost:8080/")
	if err := http.ListenAndServe(":8080", mux); err != nil { //nolint:gosec
		log.Fatal(err)
	}
	log.Println("Сервер зупинено")
}
