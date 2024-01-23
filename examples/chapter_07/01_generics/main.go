package main

import (
	"fmt"
	"slices"
	"strings"
)

type ValidTypes interface {
	~string | ~int | ~int32 | ~int64 | ~float64
}

// Add додає два значення
func Add[T ValidTypes](one, two T) T {
	return one + two
}

// IsEqual порівнює два значення
func IsEqual[T ValidTypes](one, two T) bool {
	return one == two
}

// AddAndCompare додає два значення та порівнює їх
func AddAndCompare[T ValidTypes, C comparable](one, two T, x C, z C) (T, bool) {
	return Add(one, two), x == z
}

func AddMultiple[T ValidTypes, E []T](in ...T) (T, E) {
	var result T
	var slice E
	for _, v := range in {
		result += v
		slice = append(slice, v)
	}
	return result, slice
}

type MyString string

func (ms MyString) String() string {
	return "[" + string(ms) + "]"
}

type SomeString interface {
	~string
}
type MySlice[T SomeString] []T

func (ms MySlice[T]) String() string {
	return ms.Join("-")
}

func (ms MySlice[T]) TypedString() T {
	var out T
	for _, v := range ms {
		out += T(fmt.Sprint(v)) + "--"
	}
	return out[:len(out)-2]
}

func (ms MySlice[T]) Join(sep string) string {
	var stringSlice []string
	for _, v := range ms {
		stringSlice = append(stringSlice, string(v))
	}
	return strings.Join(stringSlice, sep)
}

func main() {
	filler := strings.Repeat("-", 20)
	fmt.Println(Add(1, 2))
	fmt.Println(IsEqual(1, 2))
	fmt.Println(IsEqual(2, 2))
	fmt.Println(Add("1", "2"))
	fmt.Println(Add(1.1, 2.2))
	fmt.Println(Add[float64](1, 2.3))
	var a, b MyString = "Hello ", "World!"
	fmt.Println(Add(a, b))
	fmt.Println(IsEqual(a, b))
	fmt.Println(IsEqual(a, a))
	println(Add(a, b))
	fmt.Printf("%sAddAndCompare%s\n", filler, filler)
	fmt.Println(AddAndCompare(a, b, a, b))
	fmt.Printf("%sAddMultiple%s\n", filler, filler)
	fmt.Println(AddMultiple(1, 2, 3, 4, 5, 6, 7, 8))
	fmt.Println(AddMultiple("1", "2", "3", "4", "5", "6", "7", "8"))
	fmt.Printf("%sAddMultiple[MyString]%s\n", filler, filler)
	fmt.Println(AddMultiple[MyString]())
	fmt.Println(AddMultiple[MyString]("One", "Two", "Three"))
	fmt.Printf("%sMySlice[MyString]%s\n", filler, filler)
	mySlice := MySlice[MyString]{"One", "Two", "Three"}
	fmt.Println(mySlice)
	slices.Sort(mySlice)
	fmt.Println(mySlice)
	fmt.Println(mySlice.TypedString())
	fmt.Printf("%sMySlice[string]%s\n", filler, filler)
	stringSlice := MySlice[string]{"One", "Two", "Three"}
	fmt.Println(stringSlice)
	slices.Sort(stringSlice)
	fmt.Println(stringSlice)
	fmt.Println(stringSlice.TypedString())
}
