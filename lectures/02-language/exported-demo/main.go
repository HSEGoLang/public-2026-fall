package main

import (
	"fmt"

	"hse-go/lectures/02-language/exported"
)

func main() {
	exported.Name = "geometry"
	point := exported.Point{X: 7}

	fmt.Printf("%s: X=%d gravity=%.1f distance=%.1f\n",
		exported.Name, point.X, exported.Gravity, exported.Distance())
}
