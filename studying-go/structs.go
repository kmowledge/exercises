package main

import "fmt"

type Animal struct{
	genre, class, huntedBy string
	eats []string
	avgLifespan int
	gender bool
}

//default features of crocodile(animal)
type crocodile struct {
	genre: "crocodile"
	class "reptile"
	eats {"fish", "birds", "mammals"}
	huntedBy {"human"}
	avgLifespan 55
	gender bool
}

func main() {
	crocTom := new(crocodile) //instance of crocodile struct
	crocTom {
		gender: true
	}
	crocTom.eats{"pumpkin"} //I override the default values
}