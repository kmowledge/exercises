package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	num := "123"
	fmt.Printf("Value: %v, Type: %T\n", num, num)
	val, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Value: %v, Type: %T\n", val, val)

}
