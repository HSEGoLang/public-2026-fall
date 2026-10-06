package main

import "fmt"

func main() {
	go fmt.Println("квадрат:", 2*2)
	fmt.Println("запрос принят")
}
