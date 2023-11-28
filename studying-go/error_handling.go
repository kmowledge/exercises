package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("./slices.go")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
