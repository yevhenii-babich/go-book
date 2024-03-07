package simplesttest

import (
	"sync"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	expected := 5
	var result int
	go func() {
		defer wg.Done()
		result = Add(2, 3)
	}()
	time.Sleep(10 * time.Millisecond)
	wg.Wait()
	t.Log("Результат", result)
	if result != expected {
		t.Errorf("Результат %d; Очікувано %d", result, expected)
	}
}
