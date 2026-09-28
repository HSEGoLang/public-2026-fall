package main

import "fmt"

func boundary() {
	defer func() {
		if value := recover(); value != nil {
			fmt.Println("recovered:", value)
		}
	}()
	defer fmt.Println("cleanup")
	panic("broken invariant")
}

func main() {
	boundary()
	fmt.Println("caller continues")
}
