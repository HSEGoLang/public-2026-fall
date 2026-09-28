package main

import "fmt"

type Rocket struct{ Name string }

func (r *Rocket) Launch() {
	fmt.Println("launch:", r.Name)
}

func run(action func()) { action() }

func main() {
	r := &Rocket{Name: "Vostok"}
	run(r.Launch)
	run(func() { r.Launch() })
}
