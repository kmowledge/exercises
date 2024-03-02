package main

import (
	"fmt"
	"regexp"
)

func formatCheck(a *string) {
	for {
		if b, _ := regexp.MatchString(`[0-23]\:[0-5][0-9]\:[0-5][0-9]`, *a); b {
			break
		}
		fmt.Println("stick to the (24h) time format ---> 00:00:00")
		fmt.Scanln(a) //If you scan to &address of a pointer, you get infinite loop of the print above. Why?
	}
}

// func 30minEarlier(a)

func main() {
	fmt.Println(`	Hello, gimme time in the following format 00:00:00
	and I will tell you what time was 30 minutes ago.\n
	To exit the program, answer "enough".`)
	// for/if+return jeśli enough to pomiń wszystko, zakończ program
	var timeIn string
	fmt.Scanln(&timeIn)
	formatCheck(&timeIn)
	// timeFormatted := time.parseString(timeIn)
	// timeOut := timeFormatted.Sub(time.Minute * 30)
	fmt.Println("You said: ", timeIn)
}
