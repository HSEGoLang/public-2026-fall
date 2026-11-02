// Package pitfalls contains intentionally incorrect snippets. Do not call them.
package pitfalls

// START UNSAFE_BANK OMIT
type Account struct{ Balance int }

func (a *Account) Deposit(amount int) {
	a.Balance = a.Balance + amount
}

// END UNSAFE_BANK OMIT

// START SLICE_RACE OMIT
func SliceRace() {
	var x []int
	go func() { x = make([]int, 10) }()
	go func() { x = make([]int, 1000000) }()

	x[999999] = 1
}

// END SLICE_RACE OMIT

// START MEMORY_ORDER OMIT
func MemoryOrder() (int, int) {
	var x, y int

	reads := make(chan int, 2)

	go func() { x = 1; reads <- y }()
	go func() { y = 1; reads <- x }()

	return <-reads, <-reads
}

// END MEMORY_ORDER OMIT

// START UNSAFE_ICON OMIT
var icons map[string]string

func Icon(name string) string {
	if icons == nil {
		icons = map[string]string{"spades.png": "♠"}
	}

	return icons[name]
}

// END UNSAFE_ICON OMIT
