package main

/*
In this example, we're hosting a costume party with five guests.
We use an array to store the guests' costumes and assign a costume to each guest.
Then, we loop through the array and print out the costumes.
The comments in the code explain what's happening at each step.
I hope this example brings a smile to your face! 🎃
*/
import (
	"fmt"
)

func main() {
	// Let's say we're hosting a costume party and have invited five guests.
	// We'll represent the guests' costumes using an array.

	// First, we declare an array with 5 string elements to store the costumes.
	var costumes [5]string

	// Now, let's assign costumes to each guest.
	costumes[0] = "Vampire"
	costumes[1] = "Zombie"
	costumes[2] = "Werewolf"
	costumes[3] = "Ghost"
	costumes[4] = "Alien"

	// We can loop through the array using a for loop and print each costume.
	for i, costume := range costumes {
		fmt.Printf("Guest %d is wearing a %s costume.\n", i+1, costume)
	}

	// Our guests are all dressed up and ready to party!
	// Output:
	// Guest 1 is wearing a Vampire costume.
	// Guest 2 is wearing a Zombie costume.
	// Guest 3 is wearing a Werewolf costume.
	// Guest 4 is wearing a Ghost costume.
	// Guest 5 is wearing a Alien costume.
}
