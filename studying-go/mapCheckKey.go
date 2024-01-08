package main

import "fmt"

var fruits map[string]string

func main() {
	fmt.Println("Enter a key")
	var key string
	fmt.Scanf("%s", &key)

	fruits = make(map[string]string)
	fruits["pear"] = "\U0001F350"
	fruits["grapes"] = "\U0001F347"

	// fruits := map[string]string{
	// 	"pear":   "\U0001F350",
	// 	"grapes": "\U0001F347",
	// } decl+init via literal

	if _, ok := fruits[key]; ok {
		fmt.Printf("The fruit %s %s is in the map\n", key, fruits[key])
		return
	}
	fmt.Printf("The fruit %s is not in the map\n", key)

}
