package works

import "fmt"

var state = "not visible"

func Example() {
	// This example shows how to use the works package.
	// See the package.go file for the package code.
	fmt.Println("works.Example()")
}

func GetState() string {
	return state
}

func init() {
	fmt.Println("works.init()")
	state = "initialized"
}
