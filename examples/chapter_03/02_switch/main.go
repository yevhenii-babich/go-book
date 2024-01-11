package main

import (
	"fmt"
)

func main() {
	x := 3

	// 3.1.4 Switch Statement
	// Check the value of x and execute the corresponding case.
	switch x {
	case 1:
		// If x is 1, this block of code will be executed.
		fmt.Println("x is 1")
	case 2:
		// If x is 2, this block of code will be executed.
		fmt.Println("x is 2")
	case 3:
		// If x is 3, this block of code will be executed.
		fmt.Println("x is 3")
	default:
		// If none of the cases match, this block of code will be executed.
		fmt.Println("x is something else")
	}
}
