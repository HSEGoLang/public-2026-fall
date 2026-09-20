package main

import (
	"fmt"
	"slices"
)

func main() {
	source := make([]int, 2, 4)
	fmt.Printf("before copy: %v len=%d cap=%d\n", source, len(source), cap(source))
	copy(source, []int{1, 2})
	fmt.Printf("after copy: %v len=%d cap=%d\n", source, len(source), cap(source))

	shared := append(source, 3)
	shared[0] = 9
	fmt.Println(source, shared)

	isolated := append(source[:len(source):len(source)], 4)
	isolated[0] = 7
	fmt.Println(source, isolated)

	clone := slices.Clone(source)
	clone[0] = 5
	fmt.Println(source, clone)
}
