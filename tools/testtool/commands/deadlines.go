package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

type (
	Task struct {
		Name  string   `yaml:"task"`
		Watch []string `yaml:"watch"`
	}

	Group struct {
		Name  string `yaml:"group"`
		Tasks []Task `yaml:"tasks"`
	}

	Deadlines []Group
)

func (d Deadlines) Tasks() []*Task {
	var tasks []*Task

	for _, g := range d {
		for i := range g.Tasks {
			tasks = append(tasks, &g.Tasks[i])
		}
	}

	return tasks
}

func (d Deadlines) FindTask(name string) (*Group, *Task) {
	for _, g := range d {
		for _, t := range g.Tasks {
			if t.Name == name {
				return &g, &t
			}
		}
	}

	return nil, nil
}

func (d Deadlines) FindGroup(name string) *Group {
	for _, g := range d {
		if g.Name == name {
			return &g
		}
	}

	return nil
}

func loadDeadlines(filename string) (Deadlines, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var m struct {
		Deadlines struct {
			Schedule Deadlines `yaml:"schedule"`
		} `yaml:"deadlines"`
	}

	if err := yaml.Unmarshal(b, &m); err != nil {
		return nil, fmt.Errorf("error reading deadlines: %w", err)
	}

	d := m.Deadlines.Schedule

	return d, nil
}

func updateChangedTasksFilepathsByFilename(
	d Deadlines,
	filename string,
	tasks map[string]struct{},
) map[string]struct{} {
	stringSeparator := string(filepath.Separator)

	components := strings.Split(filename, stringSeparator)
	if len(components) < 2 {
		return tasks
	}

	taskName := components[0]
	taskFullName := taskName

	group := d.FindGroup(components[0])
	if group != nil {
		taskName = components[1]
		taskFullName = group.Name + stringSeparator + taskName
	}

	_, task := d.FindTask(taskName)
	if task != nil {
		tasks[taskFullName] = struct{}{}
	}

	for _, task := range d.Tasks() {
		for _, path := range task.Watch {
			if strings.HasPrefix(filename, path) {
				groupWatch, taskWatch := d.FindTask(task.Name)
				tasks[groupWatch.Name+stringSeparator+taskWatch.Name] = struct{}{}
			}
		}
	}

	return tasks
}

func findChangedTasksFilepaths(d Deadlines, files []string) []string {
	tasks := map[string]struct{}{}

	for _, f := range files {
		tasks = updateChangedTasksFilepathsByFilename(d, f, tasks)
	}

	var l []string
	for t := range tasks {
		l = append(l, t)
	}

	sort.Strings(l)

	return l
}
