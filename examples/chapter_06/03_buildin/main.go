package main

// Step 1: Defining Basic Types First,
// we'll define the basic type Person and the interface Worker:

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Person Basic type
type Person struct {
	Name string
	Age  int
	Sex  bool
}

// String Function that returns a description of the person
func (p Person) String() string {
	return p.Name + ", age: " + strconv.Itoa(p.Age)
}
func (p Person) Read(in []byte) (n int, err error) {
	tmp := []byte(p.String())
	copy(in, tmp)
	return len(tmp), nil
}

type Employee struct {
	Person
	Position string
}

func (e Employee) String() string {
	if e.Sex {
		e.Name = "Mr. " + e.Name
	}
	return e.Position + " " + e.Person.String()
}

// Worker Interface  with method Work()
type Worker interface {
	Work() string
}

// Step 2: Defining Specific Worker Types Next,
// we'll create several structures for different types of workers,
// each of which embeds the `Person` and implements the `Worker` interface.

// Manager type that implements Worker
type Manager struct {
	Employee
	Department string
}

func (m Manager) Work() string {
	return m.Name + " manages the department " + m.Department
}

// Developer type that implements Worker
type Developer struct {
	Employee
	Language string
}

func (d Developer) Work() string {
	return d.Name + " writes code in " + d.Language
}

type HR struct {
	Employee
	Departments []string
}

func (h HR) Work() string {
	return h.Name + " recruiter for departments " + strings.Join(h.Departments, ", ")
}

// doIt function that uses io.Reader interface
func doIt(in io.Reader) error {
	buf := make([]byte, 1024)
	n, err := in.Read(buf)
	if err != nil {
		return err
	}
	fmt.Println("read", n, "bytes:", string(buf))
	return nil
}

// WorkerStringer interface that embeds Worker and fmt.Stringer
type WorkerStringer interface {
	Worker
	fmt.Stringer
}

// RatedWorker structure that embeds Worker interface and contains Rate
type RatedWorker struct {
	Worker
	Rate int
}

// AddRate method that increments Rate
func (r *RatedWorker) AddRate() {
	r.Rate++
}

// RatedWorker2 structure that embeds RatedWorker structure and WorkerStringer interface
type RatedWorker2 struct {
	RatedWorker    // embed structure with interface, contains Worker, Rate, inherit AddRate()
	WorkerStringer // embed interface, contains Worker and fmt.Stringer
}

// Step 3: Using Polymorphism Now that we have our structures,
// we can write a function that takes a Worker and executes their Work method:

func report(w Worker) {
	fmt.Println(w.Work())
	switch wrk := w.(type) {
	case Developer:
		fmt.Println(wrk.Name + " is a developer on " + wrk.Language)
	case Manager:
		fmt.Println(wrk.Name + " is a manager for department " + wrk.Department)
	case RatedWorker2: // it's not a Worker, but we can use it as Worker
		fmt.Println(wrk.String(), ", ", wrk.Work(), ", rate: ", wrk.Rate)
	default:
		fmt.Printf("I don't know about type %T (%v)", wrk, wrk)
	}
}

//Step 4: Main Function Finally, we utilize our types in the main function:

func main() {
	manager := Manager{
		Employee: Employee{
			Person{"Alice", 35, false},
			"CTO",
		},
		Department: "Development",
	}

	developer := Developer{
		Employee: Employee{
			Person{"Bob", 30, true},
			"Senior Developer",
		},
		Language: "Go",
	}
	hr := HR{
		Employee: Employee{
			Person: Person{"July", 20, false},
		},
		Departments: []string{"Development", "HR"},
	}
	// Who is the manager?
	fmt.Println(manager)
	report(manager)
	// Who is the developer?
	fmt.Println(developer)
	report(developer)

	fmt.Println(hr)
	report(hr)
	if err := doIt(hr); err != nil {
		fmt.Println(err)
	}
	// How we can assign a value of type Developer to a structure witch embed Worker?
	var rated RatedWorker
	rated.Worker = developer // assign developer (struct) to Worker (interface)
	fmt.Println(strings.Repeat("-", 20))
	for i := 0; i < 5; i++ {
		rated.AddRate()
	}
	fmt.Println(rated) // do not contain Stringer, so have standard output
	report(rated)      // but we can use it as Worker

	var rated2 RatedWorker2
	rated2.RatedWorker = rated
	rated2.WorkerStringer = manager // init interface (required)
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println(rated2) // contains Stringer, so we can use it as Stringer and Worker
	report(rated2)
	//doIt(rated2) did not implement io.Reader (missing Read method, even if it contains Worker)
}
