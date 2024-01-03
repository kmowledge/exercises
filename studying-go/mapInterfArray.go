package main

import (
	"fmt"
)

func main() {
	var ranking map[int8]int32
	fmt.Println(ranking == nil)
	fmt.Println(ranking[1])
	ranking = make(map[int8]int32)
	fmt.Println(ranking)
	fmt.Println(ranking == nil)
	ranking[1] = 326000
	fmt.Println(ranking == nil)

	// key := int8(1) Type conversion is required, because "1" is not read as int8 by default
	var key int8
	_, err := fmt.Scan(&key)
	if err != nil {
		fmt.Print("Value error, expected decimal (int8). One more try.")
		fmt.Scan(&key)
	}
	if val, ok := ranking[key]; ok {
		fmt.Println("Wynik instrukcji warunkowej: zgadza się\nKey: ", key, " Value: ", val)
	} else {
		fmt.Println("Wynik instrukcji warunkowej: nie zgadza się\nKey: ", key, " doesn't exist in this map:\n", ranking)
	}
	var rank map[int8][2]interface{}
	rank = make(map[int8][2]interface{})
	rank[1] = [2]interface{}{36733, "BestPlayer"}
	fmt.Println(rank[1])
	fmt.Println(rank[2])

	fmt.Println("Check if the key exists in the second ranking.")

	// key1 := strconv.FormatInt(int64(key), 10)
	// fmt.Scanf("%s", &key1)
	// key = int8(len(key1))
	var key1 int8
	// fmt.Scan(&key1)
	if _, err := fmt.Scan(&key1); err != nil {
		fmt.Print("Value error, expected decimal (int8). One more try.")
		fmt.Scan(&key1)
	}
	if _, ok := rank[key1]; ok {
		fmt.Println("This key exists. On position", key1, " there is: ", rank[key1])
		return
	}
	fmt.Println("There is no such key:", key1, " in second ranking.")
	// if val, ok := rank[1].([2]interface{}); ok && 36900 == val[0] && "UnknownPlayer" == val[1] {
	// 	fmt.Println(rank[1])
	// 	return
	// }
	// fmt.Println("There is no mentioned value on position 1 in rank.")

}

//./prog.go:12:16: invalid operation: rank[1] (map index expression of type [2]interface{}) is not an interface
