package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodOptions {
			http.Error(w, "Метод не підтримується", http.StatusMethodNotAllowed)
			slog.Error("Метод не підтримується", "method", r.Method)
			return
		}
		slog.Info("Запит", "request", r)
		switch r.Method {
		case http.MethodOptions:
			w.Header().Set("Allow", "GET")
			w.WriteHeader(http.StatusNoContent)
		case http.MethodGet:
			switch r.URL.Path {
			case "/":
				_, _ = fmt.Fprintf(w, "<h1>Головна сторінка</h1><p>Вітаємо на головній сторінці</p><a href='/about'>Про нас</a><br><a href='/contact'>Контакти</a>")
				slog.Info("Головна сторінка")
			case "/about":
				_, _ = fmt.Fprintf(w, "<h1>Про нас</h1><p>Ми - команда розробників</p> <a href='/'>На головну</a>")
				slog.Info("Про нас")
			case "/contact":
				_, _ = fmt.Fprintf(w, "<h1>Контакти</h1><p>Наші контакти</p> <a href='/'>На головну</a>")
				slog.Info("Контакти")
			default:
				http.NotFound(w, r)
			}
		}
	})

	fmt.Println("Сервер запущено на http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Помилка запуску сервера", "error", err)
	}
}
