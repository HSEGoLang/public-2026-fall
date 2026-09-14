package main

import "fmt"

func main() {
	word := "hello"
	for _, word := range word {
		word := word + 'A' - 'a'
		fmt.Printf("%c", word)
	}

	fmt.Println()
	fmt.Println(word)
}
