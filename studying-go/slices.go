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

// s1 := make([]int, a)
//
//	for i, _ := range s1 {
//		_ = rand.Intn(128) //Tak się nie przypisuje wartości w slice.
//		fmt.Print(s1[i])
//	}
func main() { //Zescanowany int chcę umieścić w int8 <-128,127>, da się? Próbowałem przez pointer.
	var a int
	fmt.Println("Podaj liczbę z przedziału <-127, 128>")
	fmt.Scanln(&a)
	for !(a > -127 && a < 128) {
		fmt.Println("Podana liczba jest poza przedziałem <-127, 128>. Podaj proszę inną.")
		fmt.Scanln(&a)
	}
	s2 := make([]int, a, 2*a) //len musi być a, żeby copy się udało,
	//cap 2*a, żeby uniknąć memo reallocation przy append.
	s1 := make([]int, a)
	for i := 0; i < a; i++ {
		s1[i] = rand.Intn(128)
	}
	fmt.Println(s1, "\n", s2, "len(s2)==2, cap(s2)==4")
	copy(s2, s1)
	fmt.Println(s1, "\n", s2)
	for i, val := range s1 {
		s1[i] = int(float32(val) / 0.8)
	}
	s2 = append(s2, s1...)
	sum2 := 0
	for _, val := range s2 {
		sum2 += val
	}
	fmt.Println(s1, "\n", s2)
	var b string
	switch {
	case sum2%2 != 0:
		b := "Sum of elements s2 is odd."
		fmt.Println(b)
	case sum2%3 != 0:
		b := "Sum of elements s2 is divisible by 3."
		fmt.Println(b)
	case sum2%5 != 0:
		b := "Sum of elements s2 is divisible by 5."
		fmt.Println(b)
	default:
		b := "Sum of elements s2 isn't divisible neither by 2, nor 3, nor 5. Koniec programu"
		fmt.Println(b)
		return
	}
	c := new(string)
	*c = b //c = &b -> oba przypisują adres, a wartości nie
	fmt.Println("Drukujemy c, czyli pointer stringa b, a przedtem jego adres.", "\n", c, "\n", *c)

}
