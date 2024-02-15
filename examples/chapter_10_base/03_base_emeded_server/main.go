package main

import (
	"embed"
	"log"
	"net/http"
	"strings"
)

//go:embed public/*
var content embed.FS

func main() {
	// Створення файлового сервера з використанням вбудованої файлової системи
	fs := http.FileServer(http.FS(content))
	//http.Handle("/", fs)
	// Налаштовуємо маршрутизацію
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/":
			w.Header().Set("Location", "/public/")
			w.WriteHeader(http.StatusMovedPermanently)
			log.Println("Redirect to /public")
		case strings.HasPrefix(r.URL.Path, "/public"):
			log.Println("Serve : ", r.URL.Path)
			fs.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
			log.Println("Not found : ", r.URL.Path)
		}
	})

	// Виводимо інформаційне повідомлення та запускаємо сервер
	log.Println("Server started on: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
