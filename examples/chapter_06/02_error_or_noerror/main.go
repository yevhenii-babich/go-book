package main

import "fmt"

type MyCoolError string

func (e *MyCoolError) Error() string {
	return string(*e)
}

func newErr(in string) *MyCoolError {
	if in != "" {
		out := MyCoolError(in)
		return &out
	}
	return nil
}

func newCorrectErr(in string) error {
	if in != "" {
		out := MyCoolError(in)
		return &out
	}
	return nil
}

func checkErr(err error) error {
	return err
}

func wrappedError(err error) error {
	if checkErr(err) != nil {
		return fmt.Errorf("something went wrong, error is: %w", err)
	}
	return nil
}

func main() {
	const okError = "ok, some error"
	var err error
	if err = checkErr(newErr(okError)); err != nil {
		// Prints "Not empty : error is [ok, some error]"
		fmt.Printf("Not empty : error is [%v]\n", err)
	}
	if err = checkErr(newErr("")); err != nil {
		// Prints: Hello Rob Pike: error is [<nil>] (*main.MyCoolError)
		fmt.Printf("Hello Rob Pike: error is [%v] (%T)\n", err, err)
	}
	if err = wrappedError(newErr("")); err != nil {
		// Prints: Hello Rob Pike: wrapped error is [something went wrong : <nil>] (*fmt.wrapError)
		fmt.Printf("Hello Rob Pike: wrapped error is [%v] (%T)\n", err, err)
	}
	// Correct way
	if err = checkErr(newCorrectErr("")); err != nil {
		// Do not enter here
		fmt.Printf("Hello Rob Pike: error is [%v] (%T)\n", err, err)
	}
	if err = wrappedError(newCorrectErr("")); err != nil {
		// Do not enter here
		fmt.Printf("Hello Rob Pike: wrapped error is [%v] (%T)\n", err, err)
	}
	if err = wrappedError(newCorrectErr(okError)); err != nil {
		// Correct way: wrapped error is [something went wrong : ok, some error] (*fmt.wrapError)
		fmt.Printf("Correct way: wrapped error is [%v] (%T)\n", err, err)
	}
}
