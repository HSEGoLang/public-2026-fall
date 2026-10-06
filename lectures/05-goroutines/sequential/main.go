package main

import "fmt"

func calculate(value int) int {
	return value * value
}

func main() {
	for _, value := range []int{1, 2, 3} {
		fmt.Println(calculate(value))
	}
}
