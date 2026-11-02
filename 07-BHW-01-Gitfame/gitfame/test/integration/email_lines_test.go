package integration

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmailAndLineMode(t *testing.T) {
	binary, err := binCache.GetBinary(importPath)
	require.NoError(t, err)
	repo := t.TempDir()
	fixtureGit(t, repo, nil, "init", "-q")

	write := func(name, content string) {
		t.Helper()
		require.NoError(t, os.WriteFile(filepath.Join(repo, name), []byte(content), 0o600))
	}
	commit := func(name, email, committerName, committerEmail string) {
		t.Helper()
		fixtureGit(t, repo, nil, "add", ".")
		fixtureGit(t, repo, []string{
			"GIT_AUTHOR_NAME=" + name,
			"GIT_AUTHOR_EMAIL=" + email,
			"GIT_COMMITTER_NAME=" + committerName,
			"GIT_COMMITTER_EMAIL=" + committerEmail,
			"GIT_AUTHOR_DATE=2026-01-01T00:00:00+00:00",
			"GIT_COMMITTER_DATE=2026-01-01T00:00:00+00:00",
		}, "-c", "commit.gpgsign=false", "commit", "-qm", "fixture")
	}

	write("code.go", "alpha\nbeta\n")
	write("second.go", "gamma\n")
	commit("Same Name", "alice@example.test", "Bot", "bot@example.test")

	// This contributor owns only whitespace and an empty file.
	write("code.go", "alpha\nbeta\n \n\t\n\u00a0\n")
	write("blank.go", " \n\t\n\u2003\n")
	write("empty.txt", "")
	commit("Same Name", "bob@example.test", "CI", "ci@example.test")

	// Renaming a person does not change their email identity. One commit
	// contributes to two files and must still count as one commit.
	write("second.go", "gamma\ndelta\n")
	write("new.go", "epsilon\n")
	commit("Renamed Author", "alice@example.test", "Renamed Bot", "bot@example.test")

	oldRevision := string(fixtureGit(t, repo, nil, "rev-parse", "HEAD"))
	oldRevision = oldRevision[:len(oldRevision)-1]

	// Email comparison is case-sensitive; a final newline is optional.
	write("case.go", "zeta")
	commit("Same Name", "ALICE@example.test", "Other", "other@example.test")

	const (
		header = "Email,Lines,Commits,Files\n"
		alice  = "alice@example.test,5,2,3\n"
		upper  = "ALICE@example.test,1,1,1\n"
	)

	tests := []struct {
		name string
		args []string
		want string
	}{
		{"default nonblank", nil, header + alice + upper},
		{"explicit nonblank", []string{"--line-mode", "nonblank"}, header + alice + upper},
		{
			"all",
			[]string{"--line-mode", "all"},
			header + "bob@example.test,6,1,3\n" + alice + upper,
		},
		{
			"committer nonblank",
			[]string{"--use-committer"},
			header + "bot@example.test,5,2,3\nother@example.test,1,1,1\n",
		},
		{
			"committer all",
			[]string{"--use-committer", "--line-mode", "all"},
			header + "ci@example.test,6,1,3\nbot@example.test,5,2,3\nother@example.test,1,1,1\n",
		},
		{"empty nonblank", []string{"--restrict-to", "empty.txt"}, header},
		{
			"empty all",
			[]string{"--restrict-to", "empty.txt", "--line-mode", "all"},
			header + "bob@example.test,0,1,1\n",
		},
		{"whitespace nonblank", []string{"--restrict-to", "blank.go"}, header},
		{
			"whitespace all",
			[]string{"--restrict-to", "blank.go", "--line-mode", "all"},
			header + "bob@example.test,3,1,1\n",
		},
		{"old revision", []string{"--revision", oldRevision}, header + alice},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			args := append([]string{"--repository", repo, "--format", "csv"}, tc.args...)
			cmd := exec.Command(binary, args...)
			out, err := cmd.Output()
			require.NoError(t, err)
			require.Equal(t, tc.want, string(out))
		})
	}

	t.Run("dirty working tree and old revision", func(t *testing.T) {
		write("code.go", "unstaged\n")
		write("new.go", "staged\n")
		fixtureGit(t, repo, nil, "add", "new.go")
		write("untracked.go", "untracked\n")

		before := fixtureGit(t, repo, nil, "status", "--porcelain=v1")
		head := GetHEADRef(t, repo)

		cmd := exec.Command(
			binary,
			"--repository",
			repo,
			"--revision",
			oldRevision,
			"--format",
			"csv",
		)
		out, err := cmd.Output()
		require.NoError(t, err)
		require.Equal(t, header+alice, string(out))
		require.Equal(t, head, GetHEADRef(t, repo))
		require.Equal(t, before, fixtureGit(t, repo, nil, "status", "--porcelain=v1"))

		for name, want := range map[string]string{
			"code.go": "unstaged\n", "new.go": "staged\n", "untracked.go": "untracked\n",
		} {
			data, err := os.ReadFile(filepath.Join(repo, name))
			require.NoError(t, err)
			require.Equal(t, want, string(data))
		}
	})

	t.Run("invalid line mode", func(t *testing.T) {
		cmd := exec.Command(binary, "--repository", repo, "--line-mode", "unknown")
		out, err := cmd.Output()
		require.Error(t, err)
		require.IsType(t, &exec.ExitError{}, err)
		require.Empty(t, out)
	})
}

func fixtureGit(t *testing.T, repo string, env []string, args ...string) []byte {
	t.Helper()

	cmd := exec.Command("git", args...)
	cmd.Dir = repo

	cmd.Env = append(os.Environ(), "GIT_CONFIG_GLOBAL="+os.DevNull, "GIT_CONFIG_NOSYSTEM=1")
	cmd.Env = append(cmd.Env, env...)
	out, err := cmd.CombinedOutput()
	require.NoError(t, err, "%s", out)

	return out
}
