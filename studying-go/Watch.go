package main

import (
	"fmt"
	"time"
)

type Display struct {
	Mode []string
	Time time.Time //Time?
}

func (t Time) Clock() (hour, min, sec int)

type Snap uint8

const (
	Button1Start Snap = iota
	Button1Stop  Snap
	Button1Reset Snap
	Button2Mode  Snap
	Button3Plus  Snap
	Button4Minus Snap
)

func (w Watch) Snap(s Snap) {
	switch s {
	case Button1Start:
		fmt.Println("Start")
		Start()
	case Button1Stop:
		fmt.Println("Stop")
		Stop()
	case Button1Reset:
		fmt.Println("Reset")
		Reset()
	case Button2Mode:
		fmt.Println("2Mode")
		Mode()
	case Button3Plus:
		fmt.Println("3Plus")
		Plus()
	case Button4Minus:
		fmt.Println("Minus")
		Minus()
	}
}

type Mode uint8

const (
	Watch     Mode = iota
	Stopwatch Mode
	Timer     Mode
)

type Starter interface {
	Start()
}

func (w *Watch) Start() {
	time.Start()
}

type Stopwatch interface {
	Starter()
	Stopper()
	Resetter()
}

type Timer interface {
	Starter()
	Stopper()
	Resetter()
	Setter()
}

func main() {

}
