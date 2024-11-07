// Used: "for" loop, "switch" statement
// Control flow: Right input exits the program. Otherwise the loop continues.
// Use case: command-line program (e.g. game), chatbot.

package main

import "fmt"

func main() {
	var selection string
	fmt.Print("Enter an option: new, load, exit. ")

	for {
		fmt.Scan(&selection)
		switch selection {
		case "new":
			fmt.Println("Starting a new game!")
			return
		case "load":
			fmt.Println("Loading a saved game.")
			return
		case "exit":
			fmt.Println("Exit the game.")
			return
		default:
			fmt.Println("Invalid selection. Try again.")
		}
	}
}
