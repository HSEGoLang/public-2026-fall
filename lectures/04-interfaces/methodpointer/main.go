package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{X: 1, Y: 2}
	q := Point{X: 4, Y: 6}
	fmt.Println(p.Distance(q))
	bound := p.Distance
	p.ScaleBy(2)
	fmt.Printf("bound: %.2f\n", bound(q))
	distance := (*Point).Distance
	fmt.Printf("expression: %.2f\n", distance(&p, q))
}
