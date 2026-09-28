package main

import "fmt"

// TYPES OMIT
type Counter struct{ N int }

func (c *Counter) Inc() { c.N++ }

type ByValue struct{ Counter }
type ByPointer struct{ *Counter }

type Incrementer interface{ Inc() }

var _ Incrementer = (*ByValue)(nil)
var _ Incrementer = ByPointer{}

// END TYPES OMIT

func main() {
	var value ByValue
	value.Inc() // OK: value адресуемо, эквивалентно (&value).Inc()
	fmt.Println(value.N)
	pointer := ByPointer{Counter: &Counter{}}
	var inc Incrementer = pointer
	inc.Inc()
	fmt.Println(pointer.N)
	// var _ Incrementer = value // ошибка: Inc имеет receiver *Counter
}
