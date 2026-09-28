package main

import (
	"fmt"
	"io"
)

func main() {
	var rc io.ReadCloser
	var assigned io.Reader = rc
	fmt.Println("assignment:", assigned == nil)
	asserted, ok := rc.(io.Reader)
	fmt.Println("assertion:", asserted == nil, ok)
	// _ = rc.(io.Reader)
}
