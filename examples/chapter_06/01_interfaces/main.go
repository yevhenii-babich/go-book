package main

import (
	"fmt"
	"math"
	"strings"
)

type Geometry interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	side float64
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

func (s *Square) Perimeter() float64 {
	return 4 * s.side
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (t Triangle) Perimeter() float64 {
	// Assuming an equilateral triangle for the sake of simplicity
	return 3 * t.base
}

func PrintGeometry(in Geometry) {
	// Get the type of the input
	inType := fmt.Sprintf("%T", in)
	inType = strings.TrimPrefix(inType, "*")
	inType = strings.TrimPrefix(inType, "*main.")
	// Print the area and perimeter of the input
	fmt.Printf("Area of %s: %0.3f\n", inType, in.Area())
	fmt.Printf("Perimeter of %s: %0.3f\n", inType, in.Perimeter())
}
func main() {
	// Creating a square of side 4
	PrintGeometry(&Square{4})
	// Creating a circle of radius 2
	PrintGeometry(&Circle{2})
	// Creating a triangle of base 3 and height 4
	PrintGeometry(&Triangle{3, 4})
	// Creating a square of side 4 (without pointer)
	PrintGeometry(Triangle{3, 4})
}
