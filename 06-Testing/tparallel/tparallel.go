//go:build !solution

package tparallel

type T struct{}

func (t *T) Parallel() {
	panic("implement me")
}

func (t *T) Run(subtest func(t *T)) {
	_ = subtest

	panic("implement me")
}

func Run(topTests []func(t *T)) {
	_ = topTests

	panic("implement me")
}
