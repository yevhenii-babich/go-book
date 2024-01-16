package main

import (
	"errors"
	"fmt"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

type MyData struct {
	A, B int
}

// Функція modifyData змінює поля структури, на яку вказує 'd'.
func modifyData(d *MyData) {
	d.A = 100
	d.B = 200
}

// Incrementor Функція повертає іншу функцію, яка використовує змінну 'start'.
func Incrementor(start int) func() int {
	// 'count' захоплюється і зберігається між викликами повернутої функції.
	count := start
	return func() int {
		// Кожен раз, коли викликається ця функція, 'count' збільшується на одиницю.
		count++
		return count
	}
}

// Add Функція приймає будь-яку кількість аргументів типу int
func Add(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(Divide(5, 0))
	if res, err := Divide(7, 2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	data := MyData{A: 1, B: 2}
	modifyData(&data) // Передаємо покажчик на структуру 'data'
	fmt.Println(data) // Виведе: {100 200}
	// Створюємо замикання з початковим значенням 10.
	inc := Incrementor(10)

	// Кожен виклик inc() збільшує 'count' і повертає нове значення.
	fmt.Println(inc())           // Виведе: 11
	fmt.Println(inc())           // Виведе: 12
	fmt.Println(inc())           // Виведе: 13
	fmt.Println(Add(1, 2))       // Виведе: 3
	fmt.Println(Add(1, 2, 3, 4)) // Виведе: 10
}
