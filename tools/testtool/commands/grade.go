package commands

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const (
	privateRepoRoot = "/opt/shad"
	manytaskYML     = ".manytask.yml"
)

func grade() error {
	userID := os.Getenv("GITLAB_USER_ID")
	testerToken := os.Getenv("TESTER_TOKEN")
	submitRoot := os.Getenv("CI_PROJECT_DIR")

	changedFiles, err := listChangedFiles(submitRoot)
	if err != nil {
		return err
	}

	if len(changedFiles) != 0 {
		log.Printf("changed files:")

		for _, file := range changedFiles {
			log.Printf("~ %s", file)
		}
	}

	deadlines, err := loadDeadlines(filepath.Join(privateRepoRoot, manytaskYML))
	if err != nil {
		return err
	}

	changedTasksFilepaths := findChangedTasksFilepaths(deadlines, changedFiles)
	log.Printf("detected change in tasks %v", changedTasksFilepaths)

	var failed bool

	for _, taskFilepath := range changedTasksFilepaths {
		components := strings.Split(taskFilepath, string(filepath.Separator))
		task := components[len(components)-1]

		log.Printf("testing task %s", task)

		var testFailed bool

		err := testSubmission(submitRoot, privateRepoRoot, taskFilepath, task)
		if err != nil {
			log.Printf("task %s failed: %s", task, err)

			failed = true

			var testFailedErr *TestFailedError

			testFailed = errors.As(err, &testFailedErr)

			if !testFailed {
				continue
			}
		} else {
			log.Printf("task %s passed", task)
		}

		if err := reportTestResults(testerToken, task, userID, testFailed); err != nil {
			log.Fatal(err)
		}
	}

	if failed {
		return fmt.Errorf("some tasks failed")
	}

	return nil
}

var gradeCmd = &cobra.Command{
	Use:   "grade",
	Short: "test all tasks in the last commit",
	Run: func(_ *cobra.Command, _ []string) {
		if err := grade(); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(gradeCmd)
}
