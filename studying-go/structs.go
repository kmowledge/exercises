package main

import "fmt"

// ↓↓ BASIC STRUCT OPERATIONS ↓↓
// declare, initialize with literal, initialie pointer with new(), assign, access, compare for equality
/*
// declare only
type SomeStruct struct {
	f1 string
	f2 int
	f3 bool
	f5 []string
}

// initialize through literal
var Inst1 = SomeStruct{ //returns struct
	f1: "",
	f2: 10,
	//f4: "Added feature four", No such option to initialize non-declared feature.
	f5: make([]string, 5, 10),
}

func main() {
	Inst2 := new(SomeStruct)          // returns pointer to struct
	Inst2.f1 = "feature one of Inst2" //assign
	Inst1.f5 = append(Inst1.f5, "Added", "feature", "five", "to", "Inst1")
	fmt.Printf("%v\n%v\n%t\n", Inst1, Inst2.f1, &Inst1 == Inst2) //access, compare
}
*/

// ↓↓ ADVANCED STRUCT OPERATIONS ↓↓
// nested struct, anonymous struct, anonymous field, promoted field, struct tag
//for this exercise I don't enumerate all administrative units

// anonymous fields //struct json tags
type Świnoujście struct {
	Town       string //`json: "Localization"`
	Population int    //`json: "UserID"`
}

type Koszalin = struct { //Can't initialize anon fields (int: 22,) in literal.assign them.
	string
	float32
}

type WestPom struct {
	Szczecin string
	Koszalin
	Świnoujście Świnoujście //promoted fields
}

var WestPol = struct { //anonymous struct //nested struct
	WestPom WestPom
	Lubus   string
	LowSil  string
}{ //nested struct initialisation doesn't require providing innner content, just {}
	WestPom: WestPom{Szczecin: "Szczecin", Koszalin: Koszalin{},
		Świnoujście: Świnoujście{}},
	Lubus:  "Lubuskie",
	LowSil: "Dolnośląskie",
}

func main() {
	var WestPom = WestPom{ //remember about '=' when you initialize
		Szczecin:    "Szczecin",
		Koszalin:    Koszalin{},
		Świnoujście: Świnoujście{Town: "Świnoujście", Population: 40},
	}
	WestPom.Świnoujście.Population = 39 //standard nested struct, reassign value
	WestPom.float32 = 106.235           //promoted fields of nested struct (declaration Koszalin empty)
	WestPom.string = "Koszalin"
	fmt.Printf("%#v\nAccessing promoted fields: %#v %#v\n", WestPom.Świnoujście, WestPom.string, WestPom.float32)
	//fmt.Println("%#v", WestPom.string)
}
