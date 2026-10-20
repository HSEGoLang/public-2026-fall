//go:build !race

package commands

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// Exercise the real CLI with a grouped task, including the private test overlay.
// A group/task path must remain intact for package lookup while the temporary
// directory's name must be a leaf name rather than contain a path separator.
func TestCheckTaskGroupedDirectory(t *testing.T) {
	root := t.TempDir()
	studentRepo := filepath.Join(root, "student")
	privateRepo := filepath.Join(root, "private")

	write := func(repo, name, contents string) {
		t.Helper()

		target := filepath.Join(repo, name)
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			t.Fatal(err)
		}

		if err := os.WriteFile(target, []byte(contents), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	for _, repo := range []string{studentRepo, privateRepo} {
		write(repo, "go.mod", "module hse-go\n\ngo 1.27.0\n")
		write(repo, "go.sum", "")
		write(
			repo,
			".golangci.yml",
			"version: \"2\"\nlinters:\n  default: none\n  enable: [govet]\n",
		)
		write(repo, "group/sum/sum.go", "package sum\n\nfunc Sum(a, b int) int { return a + b }\n")
	}

	write(privateRepo, "group/sum/sum_private_test.go", `//go:build private

package sum_test

import (
    "testing"
    "hse-go/group/sum"
)

func TestPrivateGroupedSum(t *testing.T) {
    if got := sum.Sum(2, 3); got != 5 { t.Fatalf("Sum = %d", got) }
}
`)
	cmd := exec.Command("go", "run", "../cmd/testtool/main.go", "check-task",
		"--problem", "group/sum", "--student-repo", studentRepo, "--private-repo", privateRepo)

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("grouped task CLI failed: %v\n%s", err, output)
	}
}
