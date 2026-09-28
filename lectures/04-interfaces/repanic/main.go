package main

import "fmt"

// PANICS OMIT
func inner() {
	defer func() {
		fmt.Println("inner:", recover())
		panic("second")
	}()
	panic("first")
}

func main() {
	defer func() { fmt.Println("outer:", recover()) }()
	inner()
	fmt.Println("after inner")
}

// END PANICS OMIT
