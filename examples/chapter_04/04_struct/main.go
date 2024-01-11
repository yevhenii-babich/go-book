package main

/*
In this example, we're creating a struct to store information about pets, initializing a few pet structs,
printing out their information in a fun way, serializing them to JSON, and deserializing a JSON representation back
to a Pet struct. The comments in the code explain what's happening at each step.
Enjoy playing around with your virtual pets!
*/
import (
	"encoding/json"
	"fmt"
	"log"
)

// Let's imagine we're creating a struct to store information about pets

type Pet struct {
	Name    string `json:"name"`
	Species string `json:"species"`
	Age     int    `json:"age"`
	Hobby   string `json:"pet-hobby,omitempty"`
}

func main() {
	// We'll initialize a few pet structs using struct literals
	fluffy := Pet{
		Name:    "Fluffy",
		Species: "Cat",
		Age:     4,
		Hobby:   "scratching the couch",
	}

	spike := Pet{
		Name:    "Spike",
		Species: "Dog",
		Age:     2,
		Hobby:   "chewing shoes",
	}

	goldie := Pet{
		Name:    "Goldie",
		Species: "Fish",
		Age:     1,
	}

	// Let's print out the pet information in a fun way
	fmt.Printf("%s is a %d-year-old %s who loves %s.\n", fluffy.Name, fluffy.Age, fluffy.Species, fluffy.Hobby)
	fmt.Printf("%s is a %d-year-old %s who loves %s.\n", spike.Name, spike.Age, spike.Species, spike.Hobby)
	fmt.Printf("%s is a %d-year-old %s who swims around all day.\n", goldie.Name, goldie.Age, goldie.Species)

	// Now let's serialize our pets to JSON
	fluffyJSON, _ := json.Marshal(fluffy)
	spikeJSON, _ := json.Marshal(spike)
	goldieJSON, _ := json.Marshal(goldie)

	// Let's print the JSON representation of our pets
	fmt.Printf("Fluffy in JSON: %s\n", fluffyJSON)
	fmt.Printf("Spike in JSON: %s\n", spikeJSON)
	fmt.Printf("Goldie in JSON: %s\n", goldieJSON)

	// And now let's pretend we're creating a new pet adoption app
	// We'll deserialize the JSON back to a Pet struct
	var newPet Pet
	err := json.Unmarshal(fluffyJSON, &newPet)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %s", err)
	}

	// Let's print out the new pet's information
	fmt.Printf(
		"Welcome our new pet, %s! They're a %d-year-old %s who loves %s.\n",
		newPet.Name,
		newPet.Age,
		newPet.Species,
		newPet.Hobby,
	)
}
