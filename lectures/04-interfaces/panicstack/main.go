package main

import "fmt"

// CHAIN OMIT
func inner() {
	defer fmt.Println("inner: cleanup")
	panic("boom")
}

func outer() {
	defer fmt.Println("outer: cleanup")
	inner()
	fmt.Println("outer: after")
}

func main() {
	defer func() { fmt.Println("main: recover:", recover()) }()
	outer()
	fmt.Println("main: after")
}

// END CHAIN OMIT
