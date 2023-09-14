package main

import "fmt"

func main() {
	var selection string
	fmt.Scan(&selection)

	for selection != "new" {
		switch selection {
		case "new":
			fmt.Println("Starting a new game!")
		case "load":
			fmt.Println("Loading a saved game.")
		case "exit":
			fmt.Println("Exit the game.")
		default:
			fmt.Println("Invalid selection. Try again.")
		}
	}
}
