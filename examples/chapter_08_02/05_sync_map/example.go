package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Map

	// Збереження значень
	go m.Store("hello", "world")
	go m.Store("int", 42)
	go m.Store("float", 42.42)
	time.Sleep(10 * time.Millisecond) // чекати 10 мілісекунд на завершення всіх Goroutines
	// Отримання значення
	if value, ok := m.Load("hello"); ok {
		fmt.Println("hello:", value)
	}

	// Оновлення значення
	m.Store("hello", "Go")

	// Видалення значення
	m.Delete("int")
	fmt.Println("Перебір елементів карти:")
	// Перебір елементів карти
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("*\t%s: %v (%T)\n", key, value, value)
		return true
	})
	// Ризики, пов'язані із потенційно можливою зміною типу даних в мапі(карті)
	v, ok := m.Swap("hello", struct {
		name string
		age  int
	}{"John", 42}) //заміна значення з отриманням попередього значення (якщо воно існує)
	fmt.Println("swap key [hello]", v, ok)
	v, ok = m.Swap("int", "float") //заміна значення з отриманням попередього значення (якщо воно існує)
	fmt.Println("swap key [int]", v, ok)
	v, ok = m.Swap("float", 42) //заміна значення з отриманням попередього значення (якщо воно існує)
	fmt.Println("swap key [float]", v, ok)
	// Перебір елементів карти
	fmt.Println("Перебір елементів карти:")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("*\t%s: %+v (%T)\n", key, value, value)
		return true
	})
}
