// Three simple interfaces of different specificities: basic, replacing value, empty.
package main

import "fmt"

// 1.basic: MediaPlayer

type TitleMp3 string

/*
func (mp3 TitleMp3) Play() {
	fmt.Println("<playing music track>")
}

type TitleMp4 struct{}

func (mp4 TitleMp4) Play() {
	fmt.Println("<playing video image>")
}

type MediaPlayer interface{}

func main() {
	var file MediaPlayer
	// file = file
	// file.Play()
	fmt.Print(file, "\n")

	var song TitleMp3
	file = song
	song.Play()

	var film TitleMp4
	file = film
	film.Play()

}
*/
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
	var capital Capital
	capital = "Cracow"
	fmt.Println("Old capital of Poland: ", capital)
	capital.CapitalUpdate()
	fmt.Println("New capital of Poland: ", capital)
}
*/

//Schemat z przypisywaniem do interfejsu instancji obiektu,zastosowany przy mp3player nie działa tu.
//Powód: tam interface był pusty, tu z metodą. Tam pass by value, tu: reference. Podmieniamy val.
// var capital Capital
// capital = "Cracow"
// cu = capital
// fmt.Println("Old capital of Poland: ", cu)
// cu.CapitalUpdate()
// fmt.Println("New capital of Poland: ", cu)

// 3.empty : type adnotation

func adnotate(i interface{}) {
	fmt.Printf("Value: %v\nType: %T\n", i, i)
}

func main() {
	var i interface{}
	i = "Writing a simple sentence."
	adnotate(i)
}
