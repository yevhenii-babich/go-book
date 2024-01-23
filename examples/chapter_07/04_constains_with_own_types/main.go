package main

import "fmt"

type Person struct {
	Name     string
	Lastname string
	Age      int
}

func (p *Person) GetName() string {
	return p.Name
}

type Employee struct {
	Person
	Position   string
	Department string
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

type MyTypes interface {
	Person | Employee
}

func GetInfo[T MyTypes](t T) {
	//fmt.Println(t.Name, t.Lastname, t.Age) // do not work

	switch tmp := any(t).(type) {
	case Employee:
		fmt.Println(tmp.Position, tmp.Department, tmp.Name, tmp.Lastname, tmp.Age)
	case Person:
		fmt.Println(tmp.Name, tmp.Lastname, tmp.Age)
	default:
		fmt.Println("Unknown type") //never happen
	}
	fmt.Println("----")
}

func getInfo(t any) {
	switch tmp := t.(type) {
	case Employee:
		fmt.Println(tmp.Position, tmp.Department, tmp.Name, tmp.Lastname, tmp.Age)
	case Person:
		fmt.Println(tmp.Name, tmp.Lastname, tmp.Age)
	default:
		fmt.Println("Unknown type")
	}
	fmt.Println("----")

}

func main() {

	vadik := Person{
		Name:     "Vadik",
		Lastname: "Korobkin",
		Age:      30,
	}

	programer := Employee{
		Person:     vadik,
		Position:   "Programmer",
		Department: "IT",
	}
	GetInfo(vadik)
	GetInfo(programer)
	book := Book{Title: "Великий Гетсбі", Author: "Ф. Скотт Фіцджеральд", Pages: 180}
	//GetInfo(book) // do not work
	fmt.Println("====================================")
	getInfo(vadik)
	getInfo(programer)
	getInfo(book)
}
