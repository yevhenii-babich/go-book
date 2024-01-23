package main

import (
	"fmt"
	"reflect"
	"strings"
)

// PrintStructFields виводить назви та значення полів будь-якої структури
func PrintStructFields[T any](s T, level ...int) {
	var cl int
	var filler string
	if len(level) > 0 {
		cl = level[0]
		for i := 0; i < cl; i++ {
			filler += "\t"
		}
	}
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Struct {
		fmt.Println("Подано не структуру")
		return
	}

	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Kind() == reflect.Struct {
			fmt.Printf("%s%v:\n", filler, val.Type().Field(i).Name)
			PrintStructFields(val.Field(i).Interface(), cl+1)
			continue
		}

		field := val.Type().Field(i)
		if tag := field.Tag.Get("json"); tag != "" {
			tags := strings.Split(tag, ",")
			fmt.Printf("%s serialized to: %s", field.Name, tags[0])
			if len(tags) > 1 && tags[1] == "omitempty" {
				fmt.Printf(",can be ommited if empty")
			}
			fmt.Printf(", ")
		}
		value := val.Field(i)
		if value.Kind() == reflect.Slice {
			for j := 0; j < value.Len(); j++ {
				sliceValue := value.Index(j)
				if sliceValue.Kind() == reflect.Struct {
					fmt.Printf("%s\t* %v:\n", filler, field.Name)
					PrintStructFields(sliceValue.Interface(), cl+2)
					continue
				}
				fmt.Printf("%s\t* %v: %v\n", filler, field.Name, sliceValue.Interface())
			}
			continue
		}
		fmt.Printf("%s%v: %v\n", filler, field.Name, value.Interface())
	}
}

type Person struct {
	Name       string   `json:"name,omitempty"`
	Lastname   string   `json:"lastname"`
	Age        int      `json:"age,omitempty"`
	IsEmployed bool     `json:"isEmployed"`
	Employee   Employee `json:"employee"`
}
type Article struct {
	Title       string
	Description string
}
type Employee struct {
	Position   string
	Department string
}
type Book struct {
	Title    string
	Author   string
	Pages    int
	Articles []Article
}

func main() {
	employee := Employee{Position: "Менеджер", Department: "Відділ продажів"}
	person := Person{Name: "Аліса", Lastname: "Eліс", Age: 30, Employee: employee, IsEmployed: true}
	articles := []Article{{Title: "Стаття 1", Description: "Опис 1"}, {Title: "Стаття 2", Description: "Опис 2"}}
	book := Book{Title: "Великий Гетсбі", Author: "Ф. Скотт Фіцджеральд", Pages: 180, Articles: articles}

	PrintStructFields(person)
	PrintStructFields(book)
	PrintStructFields(articles) // Подано не структуру
}
