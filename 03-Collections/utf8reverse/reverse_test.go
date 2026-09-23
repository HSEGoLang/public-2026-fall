package utf8reverse_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"hse-go/03-Collections/utf8reverse"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  string
	}{
		{name: "empty", input: "", want: ""},
		{name: "ASCII", input: "Hello!", want: "!olleH"},
		{name: "Cyrillic", input: "Привет", want: "тевирП"},
		{name: "mixed", input: "Hello, 世界", want: "界世 ,olleH"},
		{name: "newlines", input: "\r\n", want: "\n\r"},
		{name: "emoji", input: "🙂x", want: "x🙂"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			require.Equal(t, testCase.want, utf8reverse.Reverse(testCase.input))
		})
	}
}
