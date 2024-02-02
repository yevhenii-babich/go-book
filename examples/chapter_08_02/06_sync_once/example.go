package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once
var resource string
var mu sync.Mutex
var workerId int

func initResource() {
	fmt.Printf("Initializing resource...\n")
	resource = "Initialized"
	fmt.Printf("Initialized by worker #%d\n", workerId)
}

func main() {
	fmt.Println("Resource state:", resource)
	for i := 0; i < 10; i++ {
		go func(id int) {
			mu.Lock()
			workerId = id
			mu.Unlock()
			once.Do(initResource) // initResource буде викликано лише один раз
		}(i)
	}

	// Чекаємо достатньо часу, щоб всі goroutines встигли запуститись
	// У реальному коді слід використовувати sync.WaitGroup або інший механізм синхронізації
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Resource state:", resource)
}
