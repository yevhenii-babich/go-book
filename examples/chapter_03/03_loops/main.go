package main

import (
	"fmt"
)

func main() {
	// 3.2.1 For Loop

	// Basic for loop
	for i := 0; i < 5; i++ {
		// This block of code will be executed five times.
		fmt.Println("Value of i:", i)
	}
	// While-style loop: for loop without the init and post statements.
	i := 0
	for i < 5 {
		fmt.Println("Value of i:", i)
		i++
	}
	// Infinite loop: for loop without the condition.
	i = 0
	for {
		fmt.Println("Value of i:", i)
		i++
		if i == 5 {
			break
		}
	}

	// 3.2.2 Range
	// Looping through a slice
	numbers := []int{10, 20, 30, 40, 50}
	{
		var i string
		for i, num := range numbers {
			// This block of code will be executed once for each element in the slice.
			fmt.Printf("Index: %d, Value: %d\n", i, num)
		}
		println(i)
	}
	// Looping through a string
	for i, char := range "Hello World" {
		// This block of code will be executed once for each character in the string.
		fmt.Printf("Index: %d, Value: %c\n", i, char)
	}
	// Loop without index
	for _, char := range "Hello World" {
		// This block of code will be executed once for each character in the string.
		fmt.Printf("Value: %c\n", char)
	}
	// loop without value
	for i := range "Hello World" {
		// This block of code will be executed once for each character in the string.
		fmt.Printf("Index: %d\n", i)
	}
	// Looping through a map
	fruits := map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "cherry",
	}

	for key, value := range fruits {
		// This block of code will be executed once for each element in the map.
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
