package main

import (
	"fmt"
	"time"
)

func main() {
	var a time.Time
	var b time.Duration
	var aa string = "13:25:00"
	a, _ = time.Parse(time.TimeOnly, aa)
	var bb string = "2h30min05s"
	b, _ = time.ParseDuration(bb)
	c := a.Add(b)
	fmt.Println(c.Format(time.TimeOnly))
}
