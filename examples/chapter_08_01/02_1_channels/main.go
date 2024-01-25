package main

import "time"

// say sends a string to a channel
func say(s string, ch chan<- string) {
	ch <- s // sends data to channel
}

// out receives a string from a channel
func out(ch <-chan string, wc chan struct{}) {
	for {
		v, ok := <-ch // receive from ch until it's closed
		if !ok {      // channel closed
			close(wc) // wc - is a "wait channel" control, when you close it receiver receive the "event" with empty data
			println("closed, exiting")
			break
		}
		println(v)
	}

}

func main() {
	// Create a channel of empty struct to signal when the 'out' goroutine is done
	wc := make(chan struct{})

	// Create a channel of strings to send messages between goroutines
	ch := make(chan string)

	// Start a goroutine that listens for messages on the 'ch' channel and signals when it's done on the 'wc' channel
	go out(ch, wc)

	// Start three goroutines that send different messages on the 'ch' channel
	go say("world", ch)
	go say("hello", ch)
	go say("probably", ch)

	// Pause the main goroutine for a short time to allow other goroutines to process
	time.Sleep(100 * time.Millisecond)

	// Close the 'ch' channel to signal no more messages will be sent
	close(ch)

	println("waiting...")

	// Wait for the 'out' goroutine to signal it's done
	<-wc

	println("done")
}
