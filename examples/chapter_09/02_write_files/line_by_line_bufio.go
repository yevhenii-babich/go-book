package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil { // помилка при створенні файлу
		fmt.Println("Помилка при створенні файлу:", err)
		return
	}
	defer file.Close() // закриття файлу після завершення роботи

	writer := bufio.NewWriter(file)             // створення буферизованого записувача
	_, err = writer.WriteString("Hello, Go!\n") // запис у файл
	if err != nil {                             // помилка при записі у файл
		fmt.Println("Помилка при записі у файл:", err)
		return
	}

	err = writer.Flush() // запис у файл з буферу
	if err != nil {      // помилка при записі у файл
		fmt.Println("Помилка при збереженні даних у файл:", err)
	}
}
