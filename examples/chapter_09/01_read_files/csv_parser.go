package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
)

func main() {
	const fn = "/ua/examples/chapter_09/01_read_files/testdata/example.csv" // шлях до файлу з даними у форматі CSV
	dir, err := os.Getwd()
	if err != nil {
		slog.Error("Помилка при отриманні поточної директорії", "error", err)
	}
	file, err := os.Open(dir + fn) // відкриття файлу
	if err != nil {
		slog.Error("Помилка при відкритті файлу", "error", err)
		return
	}
	defer file.Close() // закриття файлу після завершення роботи
	scanner := csv.NewReader(file)
	for {
		record, err := scanner.Read() // читання CSV файлу рядок за рядком
		if err != nil {
			if !errors.Is(err, io.EOF) {
				slog.Error("Помилка при читанні з файлу", "error", err)
			}
			break
		}
		for _, value := range record { // вивід рядка на екран
			fmt.Printf("[%s]\t", value)
		}
		fmt.Println()
	}
}
