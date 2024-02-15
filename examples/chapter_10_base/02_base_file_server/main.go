package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Вказуємо директорію, яка буде кореневою для файлового сервера
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.Dir(wd + "/ua/examples/chapter_10_base/02_base_file_server/public"))

	// Налаштовуємо маршрутизацію
	http.Handle("/", fs)

	// Виводимо інформаційне повідомлення та запускаємо сервер
	log.Println("Файловий сервер запущено на http://localhost:8080/")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
