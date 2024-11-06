func sender(data []int) <-chan int {
	fmt.Println("[sender] make channel")
	ch := make(chan int)

	go func() {
		defer func() {
			fmt.Println("[sender] close channel")
			close(ch)
		}()
		for _, item := range data {
			ch <- item
		}
		fmt.Println("[sender] all data sent to channel")
	}()

	fmt.Println("[sender] return <-channel")
	return ch
}

func receiver(name string, delay time.Duration, ch <-chan int) {
	for val := range ch {
		fmt.Printf("[%6s] data = %d\n", name, val)
		time.Sleep(delay * time.Millisecond)
	}
	fmt.Printf("[%6s] stopped\n", name)
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6}
	ch := sender(data)

	fmt.Println("[  main] receivers")
	go receiver("first", 10, ch)
	go receiver("second", 3, ch)

	fmt.Println("[  main] waiting for execution")
	time.Sleep(time.Second)
}