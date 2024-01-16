package main

import "sync"

// Counter is a counter that may be incremented or decremented safely.
type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Decrement() {
	c.mu.Lock()
	c.count--
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func (c *Counter) IsZero() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count == 0

}

func New(initial int) *Counter {
	return &Counter{count: initial}

}

func main() {
	var wg sync.WaitGroup
	counter := New(1)
	wg.Add(2)
	go func() {
		counter.Increment()
		println(counter.IsZero())
		println(counter.Value())
		wg.Done()
	}()
	go func() {
		counter.Decrement()
		println(counter.IsZero())
		println(counter.Value())
		wg.Done()
	}()
	wg.Wait()
}
