package main

import "fmt"

func firstAnswer() int {
	responses := make(chan int, 3)
	for range 3 {
		go func() { responses <- 2 * 2 }()
	}
	return <-responses
}

func main() {
	fmt.Println(firstAnswer())
}
