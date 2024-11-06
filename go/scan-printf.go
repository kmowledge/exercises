package main

import "fmt"

func main() {
	var name string
	var age string
	var cont bool
	fmt.Print("Jak masz na imię?")
	fmt.Scan(&name)
	fmt.Print("Ile masz lat?")
	fmt.Scan(&age)
	fmt.Println(name, " ", age)

	var p *string
	fmt.Println(p, "a teraz przypiszemy adres, ok?")
	fmt.Scan(&cont)
	// if cont != true; panic
	p = new(string)
	fmt.Println(p)

	var year int
	var mes string
	mes = "Welcome to the new year:"
	year = 2024
	liczba_znakow, blad := fmt.Printf("%s %d. Best wishes.\n", mes, year)
	if blad == nil {
		fmt.Printf("Wydrukowałeś %d znaków. Obyło się bez błędów.\n", liczba_znakow)
	}
}
