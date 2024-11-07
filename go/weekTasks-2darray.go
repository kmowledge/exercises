///////////////array
// array of arrays, 2-3D, lista to-do na tydzień, planowanie tygodnia
// problem: array of strings can't be a const. I need it to pass as argument for podrzędny array(y)

package main

import "fmt"

const (
	Monday    = "Monday"
	Tuesday   = "Tuesday"
	Wednesday = "Wednesday"
	Thursday  = "Thursday"
	Friday    = "Friday"
	Saturday  = "Saturday"
	Sunday    = "Sunday"
)

func main() {
	const day = 7
	var days = [day]string{
		0: Monday,
		1: Tuesday,
		2: Wednesday,
		3: Thursday,
		4: Friday,
		5: Saturday,
		6: Sunday,
	} //czy to może być array?

	fmt.Println(days)

	// const tasks int = 3
	var week = [days][...]int{ //array of slices(zmienna długość list zadań)
		//jednak dict of slices "Monday":

	}
	fmt.Println(week)
}
