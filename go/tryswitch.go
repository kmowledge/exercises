/*We can use it to match a variable to a single value, multiple values,
match it based on expressions and logic, or even use it without an expression.
np. switch letter{case "a", "e", "i", "o", "u": ... case "t", "f", "k", "l"} ...*/

package main

import "fmt"

// ///////////switch + for
func main() {
	var selection string
	fmt.Println("Hello, choose option: 1, 2, 3")

	for {
		fmt.Scan(&selection) //Inny zapis
		switch {             //switch selection {
		case selection == "1": //case "1":
			fmt.Println("Wybrałeś 1.")
			return
		case selection == "2":
			fmt.Println("Wybrałeś 2.")
			return
		case selection == "3":
			fmt.Println("Wybrałeś 3.")
			return
		default:
			fmt.Println("Niepoprawna opcja, spróbuj ponownie: 1,2,3.")
		}
	}
}
