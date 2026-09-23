//go:build !change

package speller

var smallNumbers = [...]string{
	"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen",
	"seventeen", "eighteen", "nineteen",
}

var tens = [...]string{
	"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}

type scale struct {
	value int64
	word  string
}

var scales = [...]scale{
	{value: 1_000_000_000, word: "billion"},
	{value: 1_000_000, word: "million"},
	{value: 1_000, word: "thousand"},
}
