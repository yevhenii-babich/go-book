package simpletest

import "errors"

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("ділення на нуль")
	}
	return a / b, nil
}
