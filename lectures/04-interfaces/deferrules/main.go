package main

import "fmt"

func order() {
	x := 1
	defer fmt.Println("argument:", x)
	defer func() { fmt.Println("closure:", x) }()
	x = 2
}

func main() {
	order()
}
