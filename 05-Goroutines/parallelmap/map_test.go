package parallelmap

import (
	"fmt"
	"reflect"
	"testing"

	"hse-go/tools/testtool"
)

func TestMapValues(t *testing.T) {
	values := []int{3, 1, 3, -2}

	want := []int{9, 1, 9, 4}
	if got := Map(values, 2, func(x int) int { return x * x }); !reflect.DeepEqual(got, want) {
		t.Fatalf("Map = %v, want %v", got, want)
	}

	if !reflect.DeepEqual(values, []int{3, 1, 3, -2}) {
		t.Fatalf("input was changed: %v", values)
	}
}

func TestMapEmpty(t *testing.T) {
	for _, values := range [][]int{nil, {}} {
		got := Map(values, 3, func(int) int {
			t.Error("f called for empty input")
			return 0
		})
		if len(got) != 0 {
			t.Fatalf("Map(empty) = %v", got)
		}
	}
}

func TestMapRepeatedValuesAreSeparateJobs(t *testing.T) {
	calls := make(chan int, 3)

	got := Map([]int{4, 4, 4}, 2, func(value int) int {
		calls <- value
		return value + 1
	})
	if !reflect.DeepEqual(got, []int{5, 5, 5}) || len(calls) != 3 {
		t.Fatalf("Map = %v, callback called %d times, want 3", got, len(calls))
	}
}

func TestMapInvalidWorkers(t *testing.T) {
	for _, workers := range []int{0, -1} {
		t.Run(fmt.Sprint(workers), func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("expected panic")
				}
			}()

			Map(nil, workers, nil)
		})
	}
}

func TestNoSyncPackageImported(t *testing.T) {
	testtool.CheckForbiddenImport(t, "sync")
}
