package main

import (
	"fmt"

	"hse-go/lectures/04-interfaces/geometry"
)

type Point = geometry.Point

func main() {
	p := Point{X: 1, Y: 2}
	q := Point{X: 4, Y: 6}
	fmt.Println(p.Distance(q))
	bound := p.Distance
	p.ScaleBy(2)
	fmt.Printf("bound: %.2f\n", bound(q))
	distance := Point.Distance
	fmt.Printf("expression: %.2f\n", distance(p, q))
}
