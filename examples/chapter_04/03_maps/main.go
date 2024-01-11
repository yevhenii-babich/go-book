package main

/*
In this example, we're organizing a superhero meetup and have a map of their secret identities.
We print out their secret identities, change them to protect the superheroes, add a new superhero to the meetup,
and check if we have the secret identity for an unknown superhero.
The comments in the code explain what's happening at each step.
Have fun with this superhero party! 🦸
*/
import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Let's imagine we're organizing a superhero meetup.
	// We'll create a map to store each superhero's secret identity.

	// First, we'll initialize a map with some key-value pairs.
	secretIdentities := map[string]string{
		"Superman":    "Clark Kent",
		"Batman":      "Bruce Wayne",
		"Spider-Man":  "Peter Parker",
		"Iron Man":    "Tony Stark",
		"Black Widow": "Natasha Romanoff",
	}

	// Let's print out their secret identities.
	for superhero, identity := range secretIdentities {
		fmt.Printf("%s is secretly %s.\n", superhero, identity)
	}
	// let sort the map
	var keys []string
	for k := range secretIdentities {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Sorted secret identities:")
	for _, k := range keys {
		fmt.Printf("%s is secretly %s.\n", k, secretIdentities[k])
	}

	// Oh no! Our secret identity map was leaked to the public!
	// Let's quickly change their secret identities to protect them.
	secretIdentities["Superman"] = "Carlos Krypton"
	secretIdentities["Batman"] = "Bart Whalen"
	secretIdentities["Spider-Man"] = "Pablo Parker"
	secretIdentities["Iron Man"] = "Tim Stork"
	secretIdentities["Black Widow"] = "Natalie Roman"
	fmt.Println(strings.Repeat("-", 50))
	// Let's print out their new secret identities.
	fmt.Println("Updated secret identities:")
	for superhero, identity := range secretIdentities {
		fmt.Printf("%s is now secretly %s.\n", superhero, identity)
	}

	// Let's add a new superhero to the meetup.
	secretIdentities["Captain America"] = "Steve Rogers"
	fmt.Println(strings.Repeat("-", 50))

	// Someone saw a superhero they don't recognize, let's check if we have their identity.
	unknownHero := "Thor"
	identity, ok := secretIdentities[unknownHero]
	if ok {
		fmt.Printf("The unknown superhero is %s, and their secret identity is %s.\n", unknownHero, identity)
	} else {
		fmt.Printf("The unknown superhero %s is not in our secret identity map. They must be new!\n", unknownHero)
	}
}
