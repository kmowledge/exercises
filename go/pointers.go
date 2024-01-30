package main

import "fmt"

func main() {
	// var p = new(string)                                 //def pointer funkcją new()
	// var r = new(string)                                 //%#v Co robi kratka w stringu? Dereferencję tak jak *?
	// fmt.Printf("%#v a teraz bez kratki: %v.\n", p, r)   //pointers
	// fmt.Printf("%#v a teraz bez kratki: %v.\n", *p, *r) //dereferenced pointers
	var s = new(string) //Dereferencja możliwa.
	fmt.Printf("s  Value:%v Type:%T \n", s, s)
	var u *string //Dereferencja niemożliwa. nil pointer dereference
	fmt.Printf("u  Value:%v Type:%T \n", u, u)
	var t string
	fmt.Printf("t  Value:%v Type:%T \n", t, t)
	fmt.Printf("&t Value:%v Type:%T \n", &t, &t) // Adres &t rozpoznawany jako pointer
	//fmt.Printf("Value:%v Type:%T \n", *t, *t) cannot indirect t (variable of type string)
	*s = "kot" //Przypisanie "kot" do s nie zmieniło wartości wskaźnika, tj. adresu.  cannot use s (untyped string constant) as *string value in assignment
	fmt.Printf("*s=Value:%#v Type:%#T  var s=new(string)\n", *s, *s)
	fmt.Println("*u= nil pointer dereference var u *string")
}
