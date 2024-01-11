package main

/*
Using together maps, slices, structs, and sorting.

In this example, we create a map of people with their names as keys and Person structs as values.
We then create a slice to store the keys, add the keys from the map to the slice, sort the slice,
and finally iterate over the sorted slice, accessing the map values using the sorted keys.
This way, we can print out the information about people in alphabetical order by their names.
We also serialize the map to JSON and print it out.
Finally, we create a slice of Person structs, sort it by age, and print it out.
*/
import (
	"fmt"
	"log/slog"
	"os"
	"sort"
)

type Person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"job,omitempty"`
	Tags     []string
}

func main() {
	// Let's create the map of people with their names as keys
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	people := map[string]Person{
		"David":   {Name: "David", Age: 30, Position: "Doctor", Tags: []string{"Dentist", "Surgeon"}},
		"Eve":     {Name: "Eve", Age: 35, Position: "Engineer", Tags: []string{"Go", "C++"}},
		"Alice":   {Name: "Alice", Age: 32},
		"Charlie": {Name: "Charlie", Age: 28, Position: "Programmer", Tags: []string{"PHP", "JavaScript", "Python"}},
		"Bob":     {Name: "Bob", Age: 25},
	}
	logger.Info("Probably unsorted  people:", "people", people)
	// Create ints slice to store the keys
	keys := make([]string, 0, len(people))

	// Add the keys from the map to the slice
	for key := range people {
		keys = append(keys, key)
	}

	// Sort the slice containing the keys
	sort.Strings(keys)

	// Iterate over the sorted slice and access the map values using the sorted keys
	logger.Info("Sorted people:")
	for _, key := range keys {
		person := people[key]
		logger.Info("person", "name", person.Name, "age", person.Age, "position", person.Position)
	}
	// slice with data
	var ageSorted []Person
	for _, key := range keys {
		person := people[key]
		ageSorted = append(ageSorted, person)
	}
	sort.Slice(ageSorted, func(i, j int) bool {
		return ageSorted[i].Age < ageSorted[j].Age
	})
	// slice ordered by age
	logger.Info("Sorted by age", "people", ageSorted)
	sort.Slice(ageSorted, func(i, j int) bool {
		return ageSorted[i].Name < ageSorted[j].Name
	})
	logger.Info("Sorted by name", "people", ageSorted)
	// clear slice data: zero out all values
	fmt.Printf("\n%+v, %p\n", ageSorted[0], &(ageSorted[0]))
	clear(ageSorted)
	fmt.Printf("\n%+v, %p\n", ageSorted[0], &(ageSorted[0]))
	ageSorted[0].Name = "Newbee"
	fmt.Printf("\n%+v, %p\n", ageSorted[0], &(ageSorted[0]))
	// clear structure: zero out all values
	eve, ok := people["Eve"]
	logger.Info("get data by key with check", "key exists", ok, "data", eve)
	// clear map: delete all keys and values
	clear(people)
	logger.Info("Cleared map:", "people", people)
	var ints = []int{0, 1, 2, 3, 4, 5, 6, 7}
	logger.Info("slice", "values", ints)
	clear(ints)
	logger.Info("cleared slice", "values", ints)
	// array do not support clear function
	var arrayInts = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//clear(arrayInts) // not allowed
	logger.Info("array", "values", arrayInts)
	ais := arrayInts[:]                       // convert array to slice
	clear(ais)                                // clear slice
	logger.Info("array", "values", arrayInts) // array is cleared
}
