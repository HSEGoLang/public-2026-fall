package main

import "fmt"

// TYPES OMIT
type L3 struct{}

func (L3) Show() { fmt.Println("L3") }

type L2 struct{ L3 }

func (L2) Show() { fmt.Println("L2") }

type L1 struct{ L2 }

// END TYPES OMIT

func main() {
	var x L1
	x.Show()
	x.L2.Show()
	x.L2.L3.Show()
	x.L3.Show()
}
