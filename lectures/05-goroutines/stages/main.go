package main

import "fmt"

func numbers(out chan<- int) {
	defer close(out)
	for value := 1; value <= 3; value++ {
		out <- value
	}
}

func squares(in <-chan int, out chan<- int) {
	defer close(out)
	for value := range in {
		out <- value * value
	}
}

func main() {
	naturals, squared := make(chan int), make(chan int)
	go numbers(naturals)
	go squares(naturals, squared)
	for value := range squared {
		fmt.Println(value)
	}
}
