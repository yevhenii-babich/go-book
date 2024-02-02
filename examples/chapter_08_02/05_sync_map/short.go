package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Map
	go m.Store("hello", "world") // Збереження значень
	time.Sleep(10 * time.Millisecond)
	value, ok := m.Load("hello") // Отримання значення
	fmt.Println("hello:", value, ok)
	m.Store("hello", "Go") // Оновлення значення
	go m.Store("int", 42)  // Збереження значень
	m.Delete("int")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("*\t%s: %v (%T)\n", key, value, value)
		return true
	})
}
