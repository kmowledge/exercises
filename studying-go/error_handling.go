package main

import (
	"errors"
	_ "fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./slices.go")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	divide(3, 0)
}

func divide(num1, num2 float32) (float32, error) {
	if num2 == 0 {
		return 0, errors.New("cannot divide by zero")
		//fmt.Errorf("Cannot divide %.2f by %.2f, num1, num2")
	}
	return num1 / num2, nil
}
