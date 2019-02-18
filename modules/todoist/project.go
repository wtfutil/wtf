package todoist

import (
	"fmt"

	"github.com/darkSasori/todoist"
)

type Project struct {
	todoist.Project

	index int
	tasks []todoist.Task
}

func NewProject(id int) *Project {
  // Todoist seems to experience a lot of network issues on their side
  // If we can't connect, handle it with an empty project until we can
	project, err := todoist.GetProject(id)
  if err != nil {
    return &Project{}
  }

	proj := &Project{
		Project: project,

		index: -1,
	}

	proj.loadTasks()

	return proj
}

func (proj *Project) isFirst() bool {
	return proj.index == 0
}

func (proj *Project) isLast() bool {
	return proj.index >= len(proj.tasks)-1
}

func (proj *Project) up() {
	proj.index = proj.index - 1

	if proj.index < 0 {
		proj.index = len(proj.tasks) - 1
	}
}

func (proj *Project) down() {
	if proj.index == -1 {
		proj.index = 0
		return
	}

	proj.index = proj.index + 1
	if proj.index >= len(proj.tasks) {
		proj.index = 0
	}
}

func (proj *Project) loadTasks() {
	tasks, err := todoist.ListTask(todoist.QueryParam{"project_id": fmt.Sprintf("%d", proj.ID)})
	if err == nil {
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
