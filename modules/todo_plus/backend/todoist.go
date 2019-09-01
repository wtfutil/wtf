package backend

import (
	"fmt"

	"github.com/darkSasori/todoist"
	"github.com/olebedev/config"
)

type Todoist struct {
}

func (todo *Todoist) Title() string {
	return "Todoist"
}

func (todo *Todoist) Setup(config *config.Config) {
	todoist.Token = config.UString("apiKey")
}

func (todo *Todoist) NewProject(id int) *Project {
	// Todoist seems to experience a lot of network issues on their side
	// If we can't connect, handle it with an empty project until we can
	project, err := todoist.GetProject(id)
	proj := &Project{
		Index: -1,
	}
	if err != nil {
		proj.Err = err
		return proj
	}

	proj.ID = project.ID
	proj.Name = project.Name

	tasks, err := todo.LoadTasks(proj.ID)
	proj.Err = err
	proj.Tasks = tasks

	return proj
}

func toTask(task todoist.Task) Task {
	return Task{
		ID:        task.ID,
		Completed: task.Completed,
		Content:   task.Content,
	}
}

func (todo *Todoist) LoadTasks(id int) ([]Task, error) {
	tasks, err := todoist.ListTask(todoist.QueryParam{"project_id": fmt.Sprintf("%d", id)})

	if err != nil {
		return nil, err
	}
	var finalTasks []Task
	for _, item := range tasks {
		finalTasks = append(finalTasks, toTask(item))
	}
	return finalTasks, nil
}

func (todo *Todoist) CloseTask(task *Task) error {
	if task != nil {
		internal := todoist.Task{ID: task.ID}
		return internal.Close()
	}
	return nil
}

func (todo *Todoist) DeleteTask(task *Task) error {
	if task != nil {
		internal := todoist.Task{ID: task.ID}
		return internal.Delete()
	}
	return nil
}
