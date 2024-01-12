// Three simple interfaces of different specificities: basic, replacing value, empty.
package main

import "fmt"

// 1.basic: MediaPlayer

type TitleMp3 string

func (mp3 TitleMp3) Play() {
	if mp3 == "" {
		fmt.Printf("/untitled/: <playing video image>\n")
	} else {
		fmt.Printf("%v: <playing music track>\n", mp3)
	}
}

type TitleMp4 struct {
	name string
}

func (mp4 TitleMp4) Play() {
	if mp4.name == "" {
		fmt.Printf("/untitled/: <playing video image>\n")
	} else {
		fmt.Printf("%v: <playing video image>\n", mp4.name)
	}
}

type MediaPlayer interface {
	Play()
}

func main() {
	TitleMp3("Beautiful Day").Play()
	var film TitleMp4
	film.Play() //casting methods on custom types' instances

	var interf MediaPlayer
	interf = TitleMp3("Californication")
	interf.Play() //using interface to cast the common method of types

	var film2 = TitleMp4{name: "Titanic"}
	interf = film2
	interf.Play()
	//interf.Play() // invalid memory address or nil pointer dereference
	// fmt.Print(interf, "\n")

	// interf = film

}

// 2.replacing value: CapitalUpdater

/*
type Capital string

func (c *Capital) CapitalUpdate() {
	*c = "Warsaw"
}

type CapitalUpdater interface {
	CapitalUpdate()
}

func main() {
	var capital Capital	     // def variable via signature
	capital := Capital("Cracow") // def literally via type constructor
	fmt.Println("Old capital of Poland: ", capital)
	capital.CapitalUpdate()
	fmt.Println("New capital of Poland: ", capital)
}
*/

//Schemat z przypisywaniem do interfejsu instancji obiektu,zastosowany przy mp3player nie działa tu.
//Powód: tam interface był pusty, tu z metodą. Tam pass by value, tu: reference. Podmieniamy val.
// Tried every possible way to put var str or var Capital 'into' interface,
// but acquired only memory address in the output. (errors: nil, can't indirect, neither addressable nor a map index expression)
// var cu CapitalUpdater
// cu = new(Capital)
// fmt.Println("Old capital of Poland: ", cu) //Old: 0xc000014070
// cu.CapitalUpdate()
// fmt.Println("New capital of Poland: ", cu) //New: 0xc000014070
// Conclusion:
// Treat interface as a box contating value in case of standard methods,
// whereas for interface that changes value of variable,
// write method for pointer receiver and cast it on instance("object") of this receiver (type/"class").

// otherstr := "Gniezno"
// var cu CapitalUpdater
// cu = otherstr //cannot use otherstr (variable of type string) as CapitalUpdater value in assignment: string does not implement CapitalUpdater (missing method CapitalUpdate)
// cu.CapitalUpdate()
// fmt.Println("New capital of Poland: ", otherstr)

// 3.empty : type adnotation

/*
func adnotate(i interface{}) {
	fmt.Printf("Value: %v\nType: %T\n", i, i)
}

func main() {
	var i interface{}
	i = "Writing a simple sentence."
	adnotate(i)

	i = 23 + 9i
	adnotate(i)
}
*/
