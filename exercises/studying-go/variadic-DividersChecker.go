package main

import "fmt"

func dividersChecker(divisible int, dividers ...int) (b bool, divs []int) {
	b = true
	for _, d := range dividers {
		if divisible%d == 0 {
			divs = append(divs, d)
		} else {
			b = false
		}
	}
	return
}

func main() {
	fmt.Println(dividersChecker(320, 2, 4, 5))
	fmt.Println(dividersChecker(320, 2, 4, 5, 7))
}

// Output:
// true [2 4 5]
// false [2 4 5]
