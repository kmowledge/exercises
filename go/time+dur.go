package main

import (
	"fmt"
	"time"
)

func main() {
	var a time.Time
	var b time.Duration
	var aa string = "13:25:00"
	var bb string = "2h30min05s"
	b, _ = time.ParseDuration(bb)
	a, _ = time.Parse(time.TimeOnly, aa)
	c := a.Add(b)
	fmt.Println(c.Format(time.TimeOnly))
}
