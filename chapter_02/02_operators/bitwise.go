package main

import "fmt"

func main() {
	bitwiseAnd := 7 & 5 // bitwiseAnd is 5 (binary 0111 & 0101 = 0101)
	fmt.Printf("bitwiseAnd := 7 & 5, = %d (%b & %b = %b)\n", bitwiseAnd, 7, 5, bitwiseAnd)
	bitwiseOr := 7 | 5 // bitwiseOr is 7 (binary 0111 | 0101 = 0111)
	fmt.Printf("bitwiseOr := 7 | 5, = %d (%b | %b = %b)\n", bitwiseOr, 7, 5, bitwiseOr)
	bitwiseXor := 7 ^ 5 // bitwiseXor is 2 (binary 0111 ^ 0101 = 0010)
	fmt.Printf("bitwiseXor := 7 ^ 5, = %d (%b ^ %b = %b)\n", bitwiseXor, 7, 5, bitwiseXor)
	bitwiseAndNot := 7 &^ 5 // bitwiseAndNot is 2 (binary 0111 &^ 0101 = 0010)
	fmt.Printf("bitwiseAndNot := 7 &^ 5, = %d (%b &^ %b = %b)\n", bitwiseAndNot, 7, 5, bitwiseAndNot)
	leftShift := 7 << 5 // leftShift is 224 (binary 0111 << 5 = 11100000)
	fmt.Printf("leftShift := 7 << 5, = %d (%b << 5 = %b)\n", leftShift, 7, leftShift)
	rightShift := 7 >> 5 // rightShift is 0 (binary 0111 >> 5 = 00000000)
	fmt.Printf("rightShift := 7 >> 5, = %d (%b >> 5 = %b)\n", rightShift, 7, rightShift)
}
