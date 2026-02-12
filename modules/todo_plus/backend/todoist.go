package backend

import (
	"fmt"

	"github.com/gopherlibs/todoist/api"
	"github.com/olebedev/config"
)

type Todoist struct {
	client   *api.Client
	projects []interface{}
}

func (todo *Todoist) Title() string {
	return "Todoist"
}

func (todo *Todoist) Setup(config *config.Config) {

	todo.client = api.New(config.UString("apiKey"))
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

	proj.ID = id
	proj.Name = "Error"

	p, err := todo.client.Project(id)
	if err != nil {
		return proj
	}

	proj.Name = p.Name

	tasks, err := todo.LoadTasks(proj.ID)
	proj.Err = err
	proj.Tasks = tasks

	return proj
}

func toTask(task api.Task) Task {
	return Task{
		ID:        task.ID,
		Completed: task.Checked,
		Name:      task.Content,
	}
}

func (todo *Todoist) LoadTasks(id string) ([]Task, error) {

	tasks, err := todo.client.Tasks(id)
	if err != nil {
		return nil, err
	}

	var finalTasks []Task
	for _, item := range tasks.Results {
		finalTasks = append(finalTasks, toTask(item))
	}
	return finalTasks, nil
}

func (todo *Todoist) CloseTask(task *Task) error {
	if task != nil {
		_, err := todo.client.TaskClose(task.ID)
		return err
	}
	return nil
}

func (todo *Todoist) DeleteTask(task *Task) error {
	if task != nil {
		_, err := todo.client.TaskDelete(task.ID)
		return err
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
