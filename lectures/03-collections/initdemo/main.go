package main

import "fmt"

var squares [4]int

func init() {
	for index := range squares {
		squares[index] = index * index
	}
}

func main() {
	fmt.Println(squares)
}
