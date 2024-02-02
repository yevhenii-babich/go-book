package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var sharedResource int

func worker(id int) {
	mu.Lock() // Блокування доступу до спільного ресурсу
	sharedResource++
	fmt.Printf("worker #%d counter:%d\n", id, sharedResource)
	time.Sleep(100 * time.Millisecond)
	mu.Unlock() // Розблокування доступу
}

func main() {
	//mu.Lock()
	for i := 0; i < 10; i++ {
		go worker(i)
	}
	//time.Sleep(time.Second)
}
