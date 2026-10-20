package testmain

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("setup")

	code := m.Run()

	fmt.Println("teardown")
	os.Exit(code)
}

func TestExample(t *testing.T) {
	t.Log("test uses prepared environment")
}
