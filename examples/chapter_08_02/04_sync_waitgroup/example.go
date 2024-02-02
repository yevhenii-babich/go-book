package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1) // Додавання лічильника для Goroutine
		go func(i int) {
			defer wg.Done()          // Вказує на завершення Goroutine
			fmt.Println("Робота", i) // Вивід повідомлення
		}(i)
	}

	wg.Wait() // Чекати на завершення всіх Goroutines
}
