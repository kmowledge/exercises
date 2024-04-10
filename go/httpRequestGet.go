package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("http://httpbin.org/get")
	if err != nil {
		log.Fatal()
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(string(body))
	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		fmt.Println("Request successful")
	}
}
