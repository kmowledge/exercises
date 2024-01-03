package main

import "fmt"

func main() {
	var ranking map[int8]int32
	fmt.Println(ranking == nil)
	fmt.Println(ranking[1])
	ranking = make(map[int8]int32)
	fmt.Println(ranking)
	fmt.Println(ranking == nil)
	ranking[1] = 326000
	fmt.Println(ranking == nil)

	var rank map[int8][2]interface{}
	rank = make(map[int8][2]interface{})
	rank[1] = [2]interface{}{36733, "BestPlayer"}
	fmt.Println(rank[1])
	fmt.Println(rank[2])

	if val, ok := rank[1].([2]interface{}); ok && 36900 == val[0] && "UnknownPlayer" == val[1] {
		fmt.Println(rank[1])
		return
	}
	fmt.Println("There is no mentioned value on position 1 in rank.")

}

//./prog.go:12:16: invalid operation: rank[1] (map index expression of type [2]interface{}) is not an interface
