package main

import "fmt"

/*
What a struct is and how we can declare and initialize it in our Go program.

How to access or assign the data of an individual struct field.

How to compare one struct with another and check for equality or inequality.
*/

// declare only
type SomeStruct struct {
	f1 string
	f2 int
	f3 bool
	f5 []string
}

// initialize through literal
var Inst1 = SomeStruct{ //returns struct
	f1: "",
	f2: 10,
	//f4: "Added feature four", No such option to initialize non-declared feature.
	f5: make([]string, 5, 10),
}

func main() {
	Inst2 := new(SomeStruct)          // returns pointer to struct
	Inst2.f1 = "feature one of Inst2" //assign
	Inst1.f5 = append(Inst1.f5, "Added", "feature", "five", "to", "Inst1")
	fmt.Printf("%v\n%v\n%t\n", Inst1, Inst2.f1, &Inst1 == Inst2) //access, compare
}

/*
Want: Animal (declare feats) -> Crocodile (init default feats) -> crocTom (original feats)
Can: Animal (declare feats) -> Crocodile (Animal) func setFeats -> instance (not working)
Can: Animal (declare feats) -> Crocodile (Animal) var, instance.


type Animal struct {
	Genre, Class, huntedBy string
	Eats                   []string
	avgLifespan            int
	Gender                 bool
}

// type Crocodile struct {
// 	Animal
// }

// default features of Crocodile(animal)
func Crocodile() Animal {
Animal:
	Animal{
		Genre:       "Crocodile",
		Class:       "reptile",
		Eats:        []string("fish", "birds", "mammals"),
		huntedBy:    "human",
		avgLifespan: 55,
		Gender:      bool,
	}
}

func main() {
	crocTom := Crocodile() //instance of Crocodile struct
	crocTom{
		Gender: true,
	}
	crocTom.Eats{"pumpkin"} //I override the default values
	fmt.Printf("%#v", crocTom)
}
*/
