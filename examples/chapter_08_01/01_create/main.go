package main

import (
	"time"
)

func say(s string) {
	defer println(s, ":Done!") // defer is a stack of functions to be executed after the function returns
	for i := 0; i < 5; i++ {
		// Sleep pauses the current goroutine for at least the duration d.
		// A negative or zero duration causes Sleep to return immediately.
		time.Sleep(1 * time.Millisecond)
		println(s) // println is a function that prints to stdout
	}
}

func main() { // main is a goroutine
	go say("world") // go is a keyword that starts a goroutine
	say("hello")    // say is a function that prints to stdout
	//wait for goroutine to finish
	<-time.After(200 * time.Millisecond) // time.After returns a channel that will send the time after the duration
}
