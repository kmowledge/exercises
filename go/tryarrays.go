package main

import "fmt"

// func main() {
// 	const a = 10
// 	const b = 4
// 	var c = 54
// 	var d = [a]int{ //don't forget '=' in declaration
// 		b: c, //don't forget the comma
// 		//other elements remain default
// 	}
// 	fmt.Print(d)
// }

// func main() {
// 	var arr = [10]int{
// 		4: 54,
// 	}
// 	fmt.Println(arr)

// }

func main() {
	const a = 10
	//2 sposoby zapisu:
	var arr1 [a]int
	var arr2 = [a]int{}
	fmt.Print(arr1, "\n", arr2)
}

// func main() {
// 	const b = 10
// 	var a [b]int
// 	fmt.Println(a)
// }
