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

	resources := make(map[string]int, 3)
	fmt.Println(len(resources))                                        //no capacity cap() concept in maps, as it is in slices
	resources["gold"], resources["wood"], resources["stone"] = 0, 0, 0 //only method for multiple adding
	fmt.Println(len(resources))
	delete(resources, "stone")
	fmt.Println(len(resources))
	resources["silver"] = 0
	for key = range resources {
		resources[key] += 100
		fmt.Println("Adding 100.")
	}
	fmt.Println(resources)

	elements := map[string]struct{}{ //---> set implementation (bool or empty struct as val)
		"fire":  struct{}{},
		"water": struct{}{},
		"earth": struct{}{},
		"wind":  struct{}{},
	}
	fmt.Println(elements, " --- extra-lightweight set, 0 bytes")
}
