package speller_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"hse-go/03-Collections/speller"
)

func TestSpell(t *testing.T) {
	testCases := []struct {
		number int64
		want   string
	}{
		{number: 0, want: "zero"},
		{number: 7, want: "seven"},
		{number: 19, want: "nineteen"},
		{number: 20, want: "twenty"},
		{number: 42, want: "forty-two"},
		{number: 100, want: "one hundred"},
		{number: 110, want: "one hundred ten"},
		{number: 1_234, want: "one thousand two hundred thirty-four"},
		{number: 1_002_345, want: "one million two thousand three hundred forty-five"},
		{number: -301, want: "minus three hundred one"},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprint(testCase.number), func(t *testing.T) {
			require.Equal(t, testCase.want, speller.Spell(testCase.number))
		})
	}
}
