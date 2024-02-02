// Більш складний приклад, з використанням пулу оброблювачів для типових завдань
package main

import (
	"fmt"
	"log"
	"time"

	"worker_pool/pool"
)

type ExampleTask string

func (e ExampleTask) Execute(no int) {
	fmt.Printf("executing:%s by worker #%d\n", string(e), no)
	time.Sleep(50 * time.Millisecond) // emulate work
	fmt.Printf("finishing:%s by worker #%d\n", string(e), no)
}

func main() {
	wp := pool.NewPool(5, 50)
	wp.Exec(ExampleTask("foo"))
	wp.Exec(ExampleTask("bar"))
	wp.Resize(3)
	poolSize := 3
	for i := 0; i < 500; i++ {
		wp.Exec(ExampleTask(fmt.Sprintf("additional_%d", i+1)))
		log.Printf("added task #%d", i)
		if i%100 == 0 {
			poolSize++
			wp.Resize(poolSize)
		}
	}
	wp.Close() // close pool and wait for all tasks to finish
	wp.Wait()
}
