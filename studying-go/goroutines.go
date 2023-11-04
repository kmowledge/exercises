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

func main() {
	now := time.Now()
	fmt.Println("Main starts.")

	fmt.Printf("Execution of doWork started in %0.8ss.\n", time.Since(now))
	for i := 0; i <= 10; i++ {
		go doWork(i)
	}

	time.Sleep(time.Second + time.Millisecond)
	// fmt.Printf("Execution of anon started in %0.8ss.\n", time.Since(now))
	go func(n time.Duration) {
		fmt.Println("Anon started in ", time.Since(now), "since beginning of the program")
		time.Sleep(n * time.Second)
		fmt.Println("Anon finished in ", time.Since(now), "since beginning of the program")
	}(5)

	time.Sleep(5 * time.Second)
	fmt.Println("Main finishes.")
	fmt.Printf("Program executed in %0.8ss.\n", time.Since(now))
}
