package main

import (
	"fmt"

	"hse-go/lectures/02-language/tempconv"
)

func main() {
	cold := tempconv.Celsius(-40)
	hot := tempconv.Fahrenheit(212)

	fmt.Printf("%.0f °C = %.0f °F\n", cold, tempconv.CToF(cold))
	fmt.Printf("%.0f °F = %.0f °C\n", hot, tempconv.FToC(hot))
}
