package parallel

import "testing"

func TestOrder(t *testing.T) {
	t.Log("parent starts")
	t.Run("child", func(t *testing.T) {
		t.Log("child pauses")
		t.Parallel()
		t.Log("child continues")
	})
	t.Log("parent returns")
}

func TestTwoChildren(t *testing.T) {
	t.Log("parent starts")
	t.Run("first", func(t *testing.T) {
		t.Log("first pauses")
		t.Parallel()
		t.Log("first continues")
	})
	t.Run("second", func(t *testing.T) {
		t.Log("second pauses")
		t.Parallel()
		t.Log("second continues")
	})
	t.Log("parent returns")
}

func TestCleanupOrder(t *testing.T) {
	t.Cleanup(func() { t.Log("cleanup first") })
	t.Cleanup(func() { t.Log("cleanup second") })

	t.Run("child", func(t *testing.T) {
		t.Parallel()
		t.Log("child uses resource")
	})
	defer t.Log("defer runs")
}
