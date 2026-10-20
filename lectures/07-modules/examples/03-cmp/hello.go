package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"

	"example.com/hello/morestrings"
)

func main() {
	got := morestrings.ReverseRunes("!oG ,olleH")
	want := "Hello, world!"

	if got != want {
		fmt.Println("ReverseRunes: mismatch (-want +got)")
		fmt.Print(cmp.Diff(want, got))
	} else {
		fmt.Println("OK")
	}
}
