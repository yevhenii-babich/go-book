package main

import (
	"bufio"
	"bytes"
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
	data, err := os.ReadFile(dir + fn)  // читання файлу у змінну
	reader := bytes.NewReader(data)     // створення нового читача
	scanner := bufio.NewScanner(reader) // створення нового сканера для читання файлу
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // читання файлу рядок за рядком
		time.Sleep(time.Second / 3) // затримка виводу на екран
	}
}
