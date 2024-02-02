package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			mu.Lock()   // Блокуємо доступ до умови або чекаємо розблокування
			cond.Wait() // Чекаємо на умову
			mu.Unlock() // Розблоковуємо доступ до умови

			fmt.Println("Горутина", i, "пробуджена")
		}(i)
	}
	// Даємо час горутинам блокуватися на умові
	time.Sleep(time.Second)
	// Пробуджуємо одну горутину
	cond.Signal()
	time.Sleep(time.Second)
	// Пробуджуємо одну горутину
	cond.Signal()
	// Даємо час на пробудження
	time.Sleep(time.Second)
	// Пробуджуємо всі решту горутин
	fmt.Println("Пробуджуємо всі решту горутин")
	cond.Broadcast()
	wg.Wait()
}
