package testequal

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type recordingT struct {
	errors  []string
	fails   int
	helpers int
}

func (r *recordingT) Errorf(format string, args ...interface{}) {
	r.errors = append(r.errors, fmt.Sprintf(format, args...))
}

func (r *recordingT) FailNow() { r.fails++ }

func (r *recordingT) Helper() { r.helpers++ }

func TestAssertionContract(t *testing.T) {
	type age int

	for _, tc := range []struct {
		name             string
		expected, actual interface{}
		equal            bool
	}{
		{name: "same int", expected: 1, actual: 1, equal: true},
		{name: "same uint", expected: uint(1), actual: uint(1), equal: true},
		{name: "different int", expected: 1, actual: 2},
		{name: "different numeric types", expected: 1, actual: int64(1)},
		{name: "string", expected: "cat", actual: "cat", equal: true},
		{name: "nil ints", expected: []int(nil), actual: []int(nil), equal: true},
		{name: "empty ints", expected: []int{}, actual: []int{}, equal: true},
		{name: "nil vs empty ints", expected: []int(nil), actual: []int{}},
		{name: "nil bytes", expected: []byte(nil), actual: []byte(nil), equal: true},
		{name: "empty bytes", expected: []byte{}, actual: []byte{}, equal: true},
		{name: "nil vs empty bytes", expected: []byte(nil), actual: []byte{}},
		{name: "nil maps", expected: map[string]string(nil), actual: map[string]string(nil), equal: true},
		{name: "empty maps", expected: map[string]string{}, actual: map[string]string{}, equal: true},
		{name: "nil vs empty maps", expected: map[string]string(nil), actual: map[string]string{}},
		{name: "plain nil", expected: nil, actual: nil},
		{name: "unsupported bool", expected: true, actual: true},
		{name: "unsupported float", expected: 1.0, actual: 1.0},
		{name: "unsupported uintptr", expected: uintptr(1), actual: uintptr(1)},
		{name: "unsupported struct", expected: struct{}{}, actual: struct{}{}},
		{name: "unsupported named type", expected: age(1), actual: age(1)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, assertion := range []struct {
				name string
				want bool
				call func(T, interface{}, interface{}, ...interface{}) bool
			}{
				{name: "Equal", want: tc.equal, call: AssertEqual},
				{name: "NotEqual", want: !tc.equal, call: AssertNotEqual},
			} {
				t.Run("Assert"+assertion.name, func(t *testing.T) {
					r := &recordingT{}
					got := assertion.call(r, tc.expected, tc.actual)
					require.Equal(t, assertion.want, got)
					require.Zero(t, r.fails, "Assert must not call FailNow")
					require.Positive(t, r.helpers)

					if assertion.want {
						require.Empty(t, r.errors)
					} else {
						require.Len(t, r.errors, 1)
					}
				})
			}

			for _, assertion := range []struct {
				name string
				want bool
				call func(T, interface{}, interface{}, ...interface{})
			}{
				{name: "Equal", want: tc.equal, call: RequireEqual},
				{name: "NotEqual", want: !tc.equal, call: RequireNotEqual},
			} {
				t.Run("Require"+assertion.name, func(t *testing.T) {
					r := &recordingT{}
					assertion.call(r, tc.expected, tc.actual)
					require.Positive(t, r.helpers)

					if assertion.want {
						require.Empty(t, r.errors)
						require.Zero(t, r.fails)
					} else {
						require.Len(t, r.errors, 1)
						require.Equal(t, 1, r.fails)
					}
				})
			}
		})
	}
}

func TestMessagesForSupportedTypes(t *testing.T) {
	for _, tc := range []struct {
		name             string
		expected, actual interface{}
		contains         []string
	}{
		{name: "int", expected: 123, actual: 456, contains: []string{"123", "456"}},
		{name: "uint", expected: uint(1), actual: uint(2), contains: []string{"1", "2"}},
		{name: "int8", expected: int8(12), actual: int8(34), contains: []string{"12", "34"}},
		{name: "uint8", expected: uint8(1), actual: uint8(2), contains: []string{"1", "2"}},
		{name: "int16", expected: int16(12), actual: int16(34), contains: []string{"12", "34"}},
		{name: "uint16", expected: uint16(1), actual: uint16(2), contains: []string{"1", "2"}},
		{name: "int32", expected: int32(12), actual: int32(34), contains: []string{"12", "34"}},
		{name: "uint32", expected: uint32(1), actual: uint32(2), contains: []string{"1", "2"}},
		{name: "int64", expected: int64(12), actual: int64(34), contains: []string{"12", "34"}},
		{name: "uint64", expected: uint64(1), actual: uint64(2), contains: []string{"1", "2"}},
		{name: "string", expected: "cat", actual: "dog", contains: []string{"cat", "dog"}},
		{name: "ints", expected: []int{1, 2}, actual: []int{3, 4}, contains: []string{"1", "2", "3", "4"}},
		{name: "bytes", expected: []byte{1}, actual: []byte{2}, contains: []string{"expected:", "actual"}},
		{
			name: "map", expected: map[string]string{"cat": "meow"}, actual: map[string]string{"dog": "woof"},
			contains: []string{"cat", "meow", "dog", "woof"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := &recordingT{}
			require.False(t, AssertEqual(r, tc.expected, tc.actual, "%s #%d", "case", 7))
			require.Len(t, r.errors, 1)
			require.Contains(t, r.errors[0], "not equal")
			require.Contains(t, r.errors[0], "case #7")
			require.False(t, strings.Contains(r.errors[0], "%!"), "invalid fmt diagnostic")

			for _, value := range tc.contains {
				require.Contains(t, r.errors[0], value)
			}
		})
	}
}

func TestUserMessageForms(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []interface{}
		want string
	}{
		{name: "none"},
		{name: "literal", args: []interface{}{"100% ready"}, want: "100% ready"},
		{name: "value", args: []interface{}{42}, want: "42"},
		{name: "format", args: []interface{}{"%s %d", "case", 7}, want: "case 7"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := &recordingT{}
			RequireNotEqual(r, 1, 1, tc.args...)
			require.Len(t, r.errors, 1)
			require.Contains(t, r.errors[0], tc.want)
			require.NotContains(t, r.errors[0], "%!")
		})
	}
}
