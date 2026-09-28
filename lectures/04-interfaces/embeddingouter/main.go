package main

import "fmt"

// TYPES OMIT
type Base struct{}

func (Base) Show() { fmt.Println("base") }

type T struct{ Base }

func (T) Show() { fmt.Println("outer") }

// END TYPES OMIT

func main() {
	var t T
	t.Show()
	t.Base.Show()
}
