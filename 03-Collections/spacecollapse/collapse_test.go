package spacecollapse_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"hse-go/03-Collections/spacecollapse"
)

func TestCollapseSpaces(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  string
	}{
		{name: "empty", input: "", want: ""},
		{name: "no spaces", input: "Hello, мир!", want: "Hello, мир!"},
		{name: "ASCII spaces", input: "Hello,   World!", want: "Hello, World!"},
		{name: "different spaces", input: "\r\n\t", want: " "},
		{name: "leading and trailing", input: " \t x\n ", want: " x "},
		{name: "non-breaking space", input: "a\u00a0\u00a0b", want: "a b"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			require.Equal(t, testCase.want, spacecollapse.CollapseSpaces(testCase.input))
		})
	}
}
