package main

type Counter int

func (ctr *Counter) Increment() {
	*ctr++
}

func (ctr *Counter) Decrement() {
	*ctr--
}

func (ctr *Counter) IsZero() bool {
	if ctr == nil {
		return true
	}
	return *ctr == 0
}

func main() {
	var counter Counter
	counter.Increment()
	println(counter.IsZero())
	counter.Decrement()
	println(counter.IsZero())
}
