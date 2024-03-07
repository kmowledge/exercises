package main

import "fmt"

const Limit = 7

func receiver(ch chan int) {
	count := 0
	for value := range ch {
		fmt.Println(value)
		count++
		if count == Limit {
			break
		}
	}
}

// do not change the code bellow
func sender(send int) chan int {
	ch := make(chan int, send)

	go func() {
		defer close(ch)
		for i := 0; i < send; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	var send int
	fmt.Scan(&send)

	ch := sender(send)
	receiver(ch)
}
