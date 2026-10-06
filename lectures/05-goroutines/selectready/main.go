package main

import "fmt"

func main() {
	left, right := make(chan int, 1), make(chan int, 1)
	left <- 4
	right <- 9
	select {
	case value := <-left:
		fmt.Println(value)
	case value := <-right:
		fmt.Println(value)
	}
}
