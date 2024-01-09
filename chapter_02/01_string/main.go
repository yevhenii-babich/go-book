package main

import "fmt"

// Example of changing a string value by index.
func main() {
	originalString := "Hello, World!"
	//originalString[0] = 'J' // Error: cannot assign to originalString[0]
	runeSlice := []rune(originalString)
	// Change the first character "H" to "J"
	runeSlice[0] = 'J'
	// here we are changing the value of originalString
	originalString = string(runeSlice)
	fmt.Println(originalString) // Output: Jello, World!
}
