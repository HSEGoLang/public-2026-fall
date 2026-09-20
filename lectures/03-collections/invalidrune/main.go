package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	input := string([]byte{0xff, 0xfe, 'A'})
	normalized := string([]rune(input))

	fmt.Printf("input: len=%d valid=%v\n", len(input), utf8.ValidString(input))
	fmt.Printf("normalized: len=%d valid=%v\n", len(normalized), utf8.ValidString(normalized))
	fmt.Printf("%q -> %q\n", input, normalized)
}
