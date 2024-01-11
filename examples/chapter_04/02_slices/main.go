package main

/*
In this example, we're hosting a pizza party with different types of pizzas.
We use a slice to store the pizzas and a slice of slices to store the pizzas and their toppings.
We also demonstrate the usage of cap, len, nil slices, and modifications affecting the underlying array of slices.
The comments in the code explain what's happening at each step. Enjoy the pizza party! 🍕
*/
import (
	"fmt"
)

func main() {
	// Imagine we're hosting a pizza party, and we have various types of pizza.
	// We'll use a slice to represent the pizzas we have on the table.

	// First, we declare a nil slice of strings.
	var pizzaTypes []string

	// Let's check if our pizzaTypes slice is nil.
	if pizzaTypes == nil {
		fmt.Println("Oh no! We don't have any pizzas yet!")
	}

	// Time to add some pizzas to our slice!
	pizzaTypes = append(pizzaTypes, "Pepperoni", "Margherita", "BBQ Chicken", "Hawaiian")

	// Now we have 4 pizzas on the table, let's check the length and capacity.
	fmt.Printf("We have %d pizzas on the table! (Capacity: %d)\n", len(pizzaTypes), cap(pizzaTypes))

	// Let's create a slice of slices to store pizzas and their toppings.
	var pizzas [][]string

	// We'll loop through the pizzaTypes and add each pizza along with its toppings to the pizzas slice.
	for _, pizzaType := range pizzaTypes {
		var toppings []string

		switch pizzaType {
		case "Pepperoni":
			toppings = []string{"Tomato Sauce", "Cheese", "Pepperoni"}
		case "Margherita":
			toppings = []string{"Tomato Sauce", "Cheese", "Basil"}
		case "BBQ Chicken":
			toppings = []string{"BBQ Sauce", "Cheese", "Chicken", "Red Onion"}
		case "Hawaiian":
			toppings = []string{"Tomato Sauce", "Cheese", "Ham", "Pineapple"}
		}

		pizzas = append(pizzas, append([]string{pizzaType}, toppings...))
	}

	// Let's print out the pizzas and their toppings.
	for _, pizza := range pizzas {
		fmt.Printf("%s pizza has these toppings: %v\n", pizza[0], pizza[1:])
	}

	// 🚀 Pro Tip: Be careful when modifying slices, as it can affect others that share the same underlying array!
	// Let's remove pineapple from the Hawaiian pizza and see what happens.
	hawaiianToppings := pizzas[3][1:]
	for i, topping := range hawaiianToppings {
		if topping == "Pineapple" {
			hawaiianToppings = append(hawaiianToppings[:i], hawaiianToppings[i+1:]...)
			break
		}
	}
	pizzas[3] = append([]string{"Modified Hawaiian"}, hawaiianToppings...)

	// Now let's see the updated list of pizzas and their toppings.
	fmt.Println("Updated pizza list:")
	for _, pizza := range pizzas {
		fmt.Printf("%s pizza has these toppings: %v\n", pizza[0], pizza[1:])
	}
	//simple slices
	mySlice2 := []int{10, 20, 30, 40, 50}
	fmt.Println(mySlice2)
	mySlice2 = append(mySlice2, 60)
	fmt.Println(mySlice2)
	subSlice := mySlice2[1:4] //sub slice: 20, 30, 40 (3 elements from index 1 to 1+3)
	fmt.Println(subSlice)
	clear(subSlice)       // new in GO 1.21
	fmt.Println(mySlice2) // [10 0 0 0 50 60]
	newSlice := make([]int, 3)
	copy(newSlice, mySlice2)
	fmt.Println(newSlice) // [10 0 0]
}
