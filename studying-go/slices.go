package main

import (
	"fmt"
	"math/rand"
)

// func main() {
// 	var y int
// 	fmt.Scan(&y)
// 	var s []string
// 	s = make([]string, y)
// 	fmt.Println(len(s))
// }

// func main() {
// 	var y int
// 	fmt.Scan(&y)
// 	sl := make([]string, y)
// 	fmt.Println(cap(sl))
// }

// func main() {
// 	// var s = []int{12, 23, 34}
// 	var s = make([]int, 3)
// 	var sn = make([]int, len(s))
// 	var n = copy(sn, s) // var n is the number of copied elements
// 	sn[0] = 0  // Here, we assign a new value to the slice elements in order
// 	sn[1] = 11 // to see whether it is a proper copy or the same slice
// 	fmt.Println(n)  // 3
// 	fmt.Println(s)  // [12 23 34] - the initial slice with no changes
// 	fmt.Println(sn) // [ 0 11 34] - the copied slice with modified elements
// }

// func main() {
// 	var s1 = []int{12, 23, 34}
// 	var s2 = []int{45, 56, 67}
// 	var s = append(s1, s2...)
// }

// //////////////////
// var (
// 	num1 int
// 	num2 int
// 	num3 int
// )

// func main() {
// 	// var num1, num2, num3 int //hyperskill
// 	// fmt.Scanln(&num1, &num2, &num3) //hyperskill
// 	// intSlice := []int{num1, num2, num3} //hyperskill
// 	fmt.Scan(&num1)
// 	fmt.Scan(&num2)
// 	fmt.Scan(&num3)
// 	var du []int
// 	var intSlice = []int{num1, num2, num3}
// 	for index, element := range intSlice {
// 		intSlice[index] = element
// 	}
// 	fmt.Println(du, "\n", intSlice)
// }

//[Użyj wiedzy] Zbierz input. Zadeklaruj slice. Inicjalizuj slice. Skopiuj do nowego wycinka.
//Za pomocą pętli for range podziel elementy wycinka przez 0,8 bez reszty.
//Pierwszy dopisz do tamtego. Zsmuj elementy nowego wycinka. Przypisz do tego pointer.
//Jeśli ²(kwadrat) tej liczby będzie mniejszy od liczby podanej na początku, wydrukuj:
//Twój input jest mniejszy od liczby zwróconej przez algorytm.

func main() { //Zescanowany int chcę umieścić w int8 <-128,127>, da się? Próbowałem przez pointer.
	var a int
	fmt.Println("Podaj liczbę z przedziału <-127, 128>")
	fmt.Scanln(&a)
	for !(a > -127 && a < 128) {
		fmt.Println("Podana liczba jest poza przedziałem <-127, 128>. Podaj proszę inną.")
		fmt.Scanln(&a)
	}

	s2 := make([]int, 2*a)
	s1 := make([]int, a)
	for _, element := range make([]int, a); _<a-1;_++ {
		s1[element] = rand.Intn(128)
		
	}

	fmt.Println(a, "\n", cap(s2), "\n", s1)
}

// var o int
// for;
