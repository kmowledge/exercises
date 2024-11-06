//Do dalszego rozpisania, printy zmienić, dodać dalsze ścieżki. zapisać specific question do nowej zmiennej.

// Used: "for" loop, "switch" statement
// Control flow: Right input exits the program. Otherwise the loop continues.
// Use case: command-line program (e.g. game), chatbot.

package main

import "fmt"

func main() {
	var selection string
	fmt.Println("Welcome to our restaurant fanpage!")
	fmt.Println("What would you like to know?")
	fmt.Println("Choose one of the options, by typing its number and sending back.")
	fmt.Println("Options:")
	fmt.Println("1. Opening hours")
	fmt.Println("2. Reserve a table")
	fmt.Println("3. Current promotions")
	fmt.Println("4. Menu")
	fmt.Println("5. Other")

	fmt.Println("If You have a specific question, we'll try to answer it within 24 hours.")

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
