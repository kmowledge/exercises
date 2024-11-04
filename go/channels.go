func receiveLines(ch <-chan string) {
	for val := range ch {
		fmt.Println(val)
	}
}

func sendHello(ch chan<- string) {
	ch <- "Hello"
}

func sendWorld(ch chan string) {
	sendCh := (chan<- string)(ch)
	sendCh <- "World!"
}

func main() {
	ch := make(chan string)
	go receiveLines(ch)

	sendHello(ch)
	sendWorld(ch)

	close(ch)
	runtime.Gosched()
}

// Output:
// Hello
// World!