package main

import "fmt"

func counter(start int) func() int {
	value := start

	return func() int {
		value++
		return value
	}
}

func main() {
	first := counter(10)
	second := counter(100)

	fmt.Println(first())
	fmt.Println(first())
	fmt.Println(second())
}
