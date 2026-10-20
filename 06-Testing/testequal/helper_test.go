package testequal

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelper(t *testing.T) {
	if mode := os.Getenv("FAIL_ASSERTIONS"); mode != "" {
		_, _, line, _ := runtime.Caller(0)
		fmt.Printf("expected-line=%d\n", line+2)
		AssertEqual(t, 1, 2, "%d must be equal to %d", 1, 2)

		_, _, line, _ = runtime.Caller(0)
		fmt.Printf("expected-line=%d\n", line+2)
		AssertNotEqual(t, 1, 1, "1 != 1")

		if mode == "equal" {
			_, _, line, _ = runtime.Caller(0)
			fmt.Printf("expected-line=%d\n", line+2)
			RequireEqual(t, 1, 2)
		} else {
			_, _, line, _ = runtime.Caller(0)
			fmt.Printf("expected-line=%d\n", line+2)
			RequireNotEqual(t, 1, 1)
		}

		fmt.Println("unexpected-return-after-require")

		return
	}

	for _, mode := range []string{"equal", "not-equal"} {
		t.Run(mode, func(t *testing.T) {
			cmd := exec.Command(os.Args[0], "-test.v", "-test.run=^TestHelper$", "-test.timeout=5s")

			cmd.Env = append(os.Environ(), "FAIL_ASSERTIONS="+mode)

			var buf bytes.Buffer

			cmd.Stdout = &buf
			cmd.Stderr = &buf
			err := cmd.Run()

			var exitError *exec.ExitError

			require.ErrorAs(t, err, &exitError)
			require.Equal(t, 1, exitError.ExitCode())

			lines := regexp.MustCompile(`expected-line=(\d+)`).
				FindAllStringSubmatch(buf.String(), -1)
			require.Len(t, lines, 3, "Assert must continue execution")

			for _, line := range lines {
				require.Contains(t, buf.String(), "helper_test.go:"+line[1]+":")
			}

			require.NotContains(t, buf.String(), "unexpected-return-after-require")
		})
	}
}
