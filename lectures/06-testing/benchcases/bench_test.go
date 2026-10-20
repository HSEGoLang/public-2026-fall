package benchcases

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkSizes(b *testing.B) {
	for _, n := range []int{10, 100} {
		b.Run(fmt.Sprint(n), func(b *testing.B) {
			input := strings.Repeat("a", n)

			b.ReportAllocs()

			for b.Loop() {
				strings.ToUpper(input)
			}
		})
	}
}
