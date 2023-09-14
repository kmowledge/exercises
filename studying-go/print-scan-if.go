package main

import "fmt"

var odp bool

func main() {
	fmt.Println("Witajcie w królestwie susła")
	fmt.Print("Czeka Was tu wiele nieporadności\n")
	fmt.Println("Zanim staniecie się ninja.")
	/*Teraz instrukcja warunkowa
	oraz wzięcie inputu.*/
	fmt.Println("Czy podejmiesz wyzwanie <znak zapyt> Odpowiedz: true/false") //dobrze, żenie misiami
	fmt.Scan(&odp)
	if odp != true {
		fmt.Printf("Żegnaj więc.\n<Budzisz się. To był sen.>\n")
	} else {
		fmt.Printf("Podążaj za mną.\n")
	}
}
