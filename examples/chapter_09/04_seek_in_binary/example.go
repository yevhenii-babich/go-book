package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("go.work")
	if err != nil {
		fmt.Println("Помилка при відкритті файлу:", err)
		return
	}
	defer file.Close()

	// Переміщення на 3 байти від початку файлу (skip "go ")
	_, err = file.Seek(3, 0)
	if err != nil {
		fmt.Println("Помилка при позіціонуванні у файлі:", err)
		return
	}

	// Читання даних після переміщення
	buffer := make([]byte, 6)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("Помилка при читанні з файлу:", err)
		return
	}
	fmt.Println("Version in go.work is: ", string(buffer))
}
