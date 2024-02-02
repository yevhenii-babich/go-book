package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"time"
)

func main() {
	const fn = "/ua/examples/chapter_09/01_read_files/line_by_line_bufio.go"
	dir, err := os.Getwd()
	if err != nil {
		slog.Error("Помилка при отриманні поточної директорії", "error", err)
	}
	file, err := os.Open(dir + fn) // відкриття файлу
	if err != nil {
		slog.Error("Помилка при відкритті файлу", "error", err)
		return
	}
	defer file.Close()                // закриття файлу після завершення роботи
	scanner := bufio.NewScanner(file) // створення нового сканера для читання файлу
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // читання файлу рядок за рядком
		time.Sleep(time.Second / 3) // затримка виводу на екран
	}
	if err := scanner.Err(); err != nil { // помилка при читанні файлу
		slog.Error("Помилка при читанні з файлу", "error", err)
	}
}
