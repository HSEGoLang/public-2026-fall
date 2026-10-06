package main

import "fmt"

func main() {
	ch := make(chan int)
	select {
	case value := <-ch:
		fmt.Println(value)
	default:
		fmt.Println("сейчас данных нет")
	}
}
