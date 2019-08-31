package todoist

import (
	"fmt"

	"github.com/darkSasori/todoist"
)

type Project struct {
	todoist.Project

	index int
	tasks []todoist.Task
	err   error
}

func NewProject(id int) *Project {
	// Todoist seems to experience a lot of network issues on their side
	// If we can't connect, handle it with an empty project until we can
	project, err := todoist.GetProject(id)
	proj := &Project{
		index: -1,
	}
	if err != nil {
		proj.err = err
		return proj
	}

	proj.Project = project

	proj.loadTasks()

	return proj
}

func (proj *Project) isLast() bool {
	return proj.index >= len(proj.tasks)-1
}

func (proj *Project) loadTasks() {
	tasks, err := todoist.ListTask(todoist.QueryParam{"project_id": fmt.Sprintf("%d", proj.ID)})
	if err != nil {
		proj.err = err
		proj.tasks = nil
	} else {
		proj.err = nil
		proj.tasks = tasks
	}
}

func (proj *Project) LongestLine() int {
	maxLen := 0

	for _, task := range proj.tasks {
		if len(task.Content) > maxLen {
			maxLen = len(task.Content)
		}
	}

	return maxLen
}

func (proj *Project) currentTask() *todoist.Task {
	if proj.index < 0 {
		return nil
	}

	return &proj.tasks[proj.index]
}

func (proj *Project) closeSelectedTask() {
	currTask := proj.currentTask()

	if currTask != nil {
		if err := currTask.Close(); err != nil {
			return
		}

		proj.loadTasks()
	}
}

func (proj *Project) deleteSelectedTask() {
	currTask := proj.currentTask()

	if currTask != nil {
		if err := currTask.Delete(); err != nil {
			return
		}

		proj.loadTasks()
	}
}
