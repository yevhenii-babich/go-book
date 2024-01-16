package main

import (
	"errors"
	"fmt"
)

func replaceReturn() (a int, err error) {
	defer func() {
		// This code will be executed before the return statement.
		// But it will replace the return values. (Named return values issue)
		// So, the return values will be 0 and "error".
		// To avoid this issue, do not use named return values.
		if err != nil {
			a = 0
		}
	}()
	a = 5
	err = errors.New("error")
	return a, err
}
func deferredPrint() {
	for i := 0; i < 5; i++ {
		defer println(i) //possible resource leak here
	}
}

func main() {
	deferredPrint()
	println("--------------------")
	fmt.Println(replaceReturn())
}
