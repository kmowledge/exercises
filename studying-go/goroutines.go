package main

import (
	"fmt"
	"time"
)

func doWork(num int) {
	fmt.Printf("Job %d started.\n", num)
	time.Sleep(time.Second * 1)
	fmt.Printf("Job %d finished.\n", num)
}

func goroutReturn(n int) int {
	return n * n
}

func main() {
	now := time.Now()
	fmt.Println("Main starts.")

	fmt.Printf("Execution of doWork started in %0.8s.\n", time.Since(now))
	for i := 0; i <= 10; i++ {
		go doWork(i)
	}

	time.Sleep(time.Second + time.Millisecond)
	// fmt.Printf("Execution of anon started in %0.8ss.\n", time.Since(now))
	go func(n time.Duration) {
		fmt.Println("Anon started in ", time.Since(now), "since the beginning.")
		time.Sleep(time.Duration(n/4) * time.Second)
		fmt.Println("Anon finished in ", time.Since(now), "since the beginning.")
	}(5)

	anonVar := func(n int) int {
		fmt.Println("Anon started in ", time.Since(now), "since the beginning.")
		time.Sleep(time.Duration(n/4) * time.Second)
		fmt.Println("Anon finished in ", time.Since(now), "since the beginning.")
		return 2 * 5
	}
	go anonVar(1)
	go anonVar(3)
	time.Sleep(2 * time.Second)
	returnVar := 123
	go func() {
		returnVar = goroutReturn(3)
	}()

	fmt.Println(`	Below we try to print values returned by anonVar and goroutReturn function.
	Only goroutReturn will succeed:
	1. it's a declared function goroutReturn,
	2. wrapped in anon func which assigns the output to an existing var returnVar.`)
	fmt.Println(anonVar)   //prints address, because it holds function
	fmt.Println(returnVar) //print int, because it holds return value of function
	fmt.Println("Main finishes.")
	fmt.Printf("Program executed in %0.8ss.\n", time.Since(now))
}
