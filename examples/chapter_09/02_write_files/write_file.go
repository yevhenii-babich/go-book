package main

import (
	"log/slog"
	"os"
)

func main() {
	if err := os.WriteFile("output.txt", []byte("Hello, Go!\n"), 0666); err != nil {
		slog.Error("Помилка при записі у файл", "error", err)
	}

}
