package main

import (
	"fmt"
)

func main() {
	x := 42
	// 3.1.1 If Statement
	// Check is x is greater than 10.
	y := 5
	if x > 10 {
		fmt.Println("x is greater than 10")
		y := 5.6
		fmt.Println(y)
	}

	// 3.1.2 If-Else Statement
	// Check if y is greater than 10.
	if y > 10 {
		// If the condition is true, this block of code will be executed.
		fmt.Println("y is greater than 10")
	} else {
		// If the condition is false, this block of code will be executed.
		fmt.Println("y is not greater than 10")
	}

	z := 10

	// 3.1.3 If-Else If Statement
	// Check is z is greater than 20 or 10.
	if z > 20 {
		// If z is greater than 20, this block of code will be executed.
		fmt.Println("z is greater than 20")
	} else if z > 10 {
		// If z is greater than 10, this block of code will be executed.
		fmt.Println("z is greater than 10")
	} else {
		// If none of the conditions above are true, this block of code will be executed.
		fmt.Println("z is not greater than 10 or 20")
	}
}
