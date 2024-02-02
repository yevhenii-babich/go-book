package main

import (
	"fmt"
	"sync"
)

type myCounter struct {
	mu    sync.Mutex
	count int
	wg    sync.WaitGroup
}

func (c *myCounter) increment() {
	c.mu.Lock()         // блокування доступу до спільного ресурсу
	defer c.mu.Unlock() // розблокування доступу після завершення функції
	defer c.wg.Done()   // вказує на завершення Goroutine
	// збільшення лічильника
	c.count++
}

func (c *myCounter) Run() {
	c.wg.Add(1)      // додавання лічильника для Goroutine
	go c.increment() // запуск Goroutine
}

func (c *myCounter) Wait() {
	c.wg.Wait() // чекати на завершення всіх Goroutines
}

func (c *myCounter) Value() int {
	c.mu.Lock()         // блокування доступу до спільного ресурсу
	defer c.mu.Unlock() // розблокування доступу після завершення функції
	return c.count      // повернення значення лічильника
}

func main() {
	var counter myCounter
	for i := 0; i < 1000; i++ { // запуск 1000 Goroutines
		counter.Run()
		if i%100 == 0 {
			fmt.Println("Значення count:", counter.Value())
		}
	}
	counter.Wait()
	fmt.Println("Значення count:", counter.Value())
}
