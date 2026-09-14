package main

import "fmt"

func localPointer() *int {
	value := 42
	return &value
}

func localValue() int {
	value := new(int)
	*value = 42
	return *value
}

func main() {
	fmt.Println(*localPointer())
	fmt.Println(localValue())
}
