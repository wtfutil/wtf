package backend

import (
	"strconv"

	"github.com/darksasori/todoist"
	"github.com/olebedev/config"
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
		i := strconv.Itoa(id.(int))
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
	i, _ := strconv.Atoi(id)
	project, err := todoist.GetProject(i)
	if err != nil {
		proj.Err = err
		return proj
	}

	proj.ID = strconv.Itoa(project.ID)
	proj.Name = project.Name

	tasks, err := todo.LoadTasks(proj.ID)
	proj.Err = err
	proj.Tasks = tasks

	return proj
}

func toTask(task todoist.Task) Task {
	id := strconv.Itoa(task.ID)
	return Task{
		ID:        id,
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
		i, _ := strconv.Atoi(task.ID)
		internal := todoist.Task{ID: i}
		return internal.Close()
	}
	return nil
}

func (todo *Todoist) DeleteTask(task *Task) error {
	if task != nil {
		i, _ := strconv.Atoi(task.ID)
		internal := todoist.Task{ID: i}
		return internal.Delete()
	}
	return nil
}

func (todo *Todoist) Sources() []string {
	var result []string
	for _, id := range todo.projects {
		i := strconv.Itoa(id.(int))
		result = append(result, i)
	}
	return result
}
