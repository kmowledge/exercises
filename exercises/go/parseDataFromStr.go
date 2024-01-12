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
	//strconv.ParseInt(num, 10, 64) //0, 8, 16, 32, 64 = int, int8...
	//strconv.ParseFloat(num, 64) //32,64 Decimal system is implicit here.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Value: %v,  Type: %T\n", val, val)

}
