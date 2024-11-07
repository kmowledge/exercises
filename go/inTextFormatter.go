package main

import "fmt"

type Text string

func (t Text) Tabformat() {
	fmt.Printf("\t%v\n", t)
}

func (t Text) DoubleQuoteFormat() {
	fmt.Printf("\"%v\"\n", t)
}

func (t Text) SingleQuoteFormat() {
	fmt.Printf("'%v'\n", t)
	// return string("'t'\n" + t) To try: method II, with return
}

type Formatter interface {
	Tabformat()
	DoubleQuoteFormat()
	SingleQuoteFormat()
}

func main() {
	var f Formatter
	var a Text
	var b string
	fmt.Println(`		Write one word which you want to format and a command,
		which specifies the formatting type You want to apply:,
		tab - inserts long space before the word,
		sq - puts the word into single quotes,
		dq - puts the word into double quotes,
		Example: myword sq -> 'myword',
		When you want to exit, answer /exit/.
		`)

	// fmt.Scanln(&a, &b)
sLoop:
	for {
		fmt.Scanln(&a, &b)
		f = a
		if a == "/exit/" {
			return
		} else {
			switch b {
			case "tab":
				f.Tabformat()
				break sLoop
			case "sq":
				f.SingleQuoteFormat()
				break sLoop
			case "dq":
				f.DoubleQuoteFormat()
				break sLoop
			default:
				fmt.Println("Wrong pattern, check the instruction and try again.")
			}
		}
	}
}
