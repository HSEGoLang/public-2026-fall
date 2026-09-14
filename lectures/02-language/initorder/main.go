package main

import "fmt"

var (
	third  = second
	second = twice(first)
	first  = 1
)

func twice(value int) int {
	return value * 2
}

func init() {
	fmt.Println("init", first, second)
}

func main() {
	fmt.Println("main", third)
}
