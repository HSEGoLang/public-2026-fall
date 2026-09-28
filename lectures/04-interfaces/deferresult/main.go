package main

import "fmt"

func result() (n int) {
	defer func() { n++ }()
	return 1
}

func main() {
	fmt.Println(result())
}
