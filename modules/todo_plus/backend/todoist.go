package backend

import (
	"fmt"

	"github.com/olebedev/config"
	"github.com/wtfutil/todoist"
)

type Todoist struct {
	projects []interface{}
}

func (todo *Todoist) Title() string {
	return "Todoist"
}

func (todo *Todoist) Setup(config *config.Config) {
	todoist.Token = config.UString("apiKey")
	todo.projects = config.UList("projects")
}

func (todo *Todoist) BuildProjects() []*Project {
	projects := []*Project{}

	for _, id := range todo.projects {
		i := fmt.Sprintf("%v", id)
		proj := todo.GetProject(i)
		projects = append(projects, proj)
	}
	return projects
}

func (todo *Todoist) GetProject(id string) *Project {
	// Todoist seems to experience a lot of network issues on their side
	// If we can't connect, handle it with an empty project until we can
	proj := &Project{
		Index:   -1,
		backend: todo,
	}
	project, err := todoist.GetProject(id)
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
		Name:      task.Content,
	}
}

func (todo *Todoist) LoadTasks(id string) ([]Task, error) {
	tasks, err := todoist.ListTask(todoist.QueryParam{"project_id": id})

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

func (todo *Todoist) Sources() []string {
	var result []string
	for _, id := range todo.projects {
		i := fmt.Sprintf("%v", id)
		result = append(result, i)
	}
	return result
}
