package main

import (
	"fmt"

	"example.com/moreints"
	"example.com/morestrings"
)

func main() {
	if diff := morestrings.CheckReverse("!oG ,olleH", "Hello, Go!"); diff == "" {
		fmt.Println("morestrings: OK")
	} else {
		fmt.Print(diff)
	}

	if diff := moreints.CheckReverse([]int{3, 2, 1}, []int{1, 2, 3}); diff == "" {
		fmt.Println("moreints: OK")
	} else {
		fmt.Print(diff)
	}
}
