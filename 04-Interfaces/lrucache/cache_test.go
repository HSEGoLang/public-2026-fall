package lrucache

import (
	"reflect"
	"testing"
)

func entries(c Cache) [][2]int {
	out := make([][2]int, 0)

	c.Range(func(k, v int) bool { out = append(out, [2]int{k, v}); return true })

	return out
}

func TestRecencyAndEviction(t *testing.T) {
	c := New(2)
	c.Set(1, 10)
	c.Set(2, 20)

	if v, ok := c.Get(1); !ok || v != 10 {
		t.Fatalf("Get(1) = %d, %v", v, ok)
	}

	c.Set(3, 30)

	if _, ok := c.Get(2); ok {
		t.Fatal("2 must be evicted")
	}

	want := [][2]int{{1, 10}, {3, 30}}
	if got := entries(c); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	if got := entries(c); !reflect.DeepEqual(got, want) {
		t.Fatal("Range changed recency")
	}
}

func TestUpdateClearAndZero(t *testing.T) {
	c := New(2)
	c.Set(1, 1)
	c.Set(2, 2)
	c.Set(1, 10)

	if got := entries(c); !reflect.DeepEqual(got, [][2]int{{2, 2}, {1, 10}}) {
		t.Fatal(got)
	}

	calls := 0

	c.Range(func(int, int) bool { calls++; return false })

	if calls != 1 {
		t.Fatal(calls)
	}

	c.Clear()
	c.Set(0, 0)

	if v, ok := c.Get(0); !ok || v != 0 {
		t.Fatal("zero is a valid key and value")
	}

	disabled := New(0)
	if disabled == nil {
		t.Fatal("New(0) returned nil")
	}

	disabled.Set(1, 1)

	if _, ok := disabled.Get(1); ok || len(entries(disabled)) != 0 {
		t.Fatal("disabled cache stores data")
	}
}
