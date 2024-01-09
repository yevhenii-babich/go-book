package main

/* Identifiers and literals example  */
const (
	age      = 42         // decimal
	userName = "username" // string
	userId   = 0b101010   // binary
	UserData = 0o52       // octal
	a1       = 0x2a       // hexadecimal
)

type Shape int

const (
	Circle   Shape = -1
	Triangle       = 3 * (1 << iota)
	Square
	Pentagon
	Hexagon
)

func main() {
	var userAge int // userAge is 0
	println("userAge:", userAge)
	userAge = age // userAge is 42
	println("userAge:", userAge)
	name := "John Doe"
	println("name:", name)
	name = userName // name is "username"
	println("name:", name)
	// multiple variables can be declared in one line
	var x, y = 10, 20
	//can be written as: x, y := 10, 20
	println("x, y:", x, y)
	x, y = userId, UserData
	// 42 is main answer to everything in the universe and beyond :)
	println("x, y, a1:", x, y, a1)
	var somePi float64 = 3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253
	println("somePi:", somePi) //Output: somePi: +3.141593e+000
	println(Circle, Triangle, Square, Pentagon, Hexagon)
}
