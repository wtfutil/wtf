package todoist

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Project is a model of todoist project entity
type Project struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	CommentCount int    `json:"comment_count"`
	Order        int    `json:"order"`
	Indent       int    `json:"indent"`
}

func decodeProject(body io.ReadCloser) (Project, error) {
	defer body.Close()
	decoder := json.NewDecoder(body)
	var project Project

	if err := decoder.Decode(&project); err != nil {
		return Project{}, err
	}
	return project, nil
}

// ListProject return all projects
//
// Example:
//		todoist.Token = "your token"
//		projects, err := todoist.ListProject()
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println(projects)
func ListProject() ([]Project, error) {
	res, err := makeRequest(http.MethodGet, "projects", nil)
	if err != nil {
		return []Project{}, err
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var projects []Project

	if err := decoder.Decode(&projects); err != nil {
		return []Project{}, err
	}

	return projects, nil
}

// GetProject return a project by id
//
// Example:
//		todoist.Token = "your token"
//		project, err := todoist.GetProject(1)
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println(project)
func GetProject(id int) (Project, error) {
	path := fmt.Sprintf("projects/%d", id)
	res, err := makeRequest(http.MethodGet, path, nil)
	if err != nil {
		return Project{}, err
	}

	return decodeProject(res.Body)
}

// CreateProject create a new project with a name
//
// Example:
//		todoist.Token = "your token"
//		project, err := todoist.CreateProject("New Project")
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println(project)
func CreateProject(name string) (Project, error) {
	project := struct {
		Name string `json:"name"`
	}{
		name,
	}

	res, err := makeRequest(http.MethodPost, "projects", project)
	if err != nil {
		return Project{}, err
	}

	return decodeProject(res.Body)
}

// Delete project
//
// Example:
//		todoist.Token = "your token"
//		project, err := todoist.GetProject(1)
//		if err != nil {
//			panic(err)
//		}
//		err = project.Delete()
//		if err != nil {
//			panic(err)
//		}
func (p Project) Delete() error {
	path := fmt.Sprintf("projects/%d", p.ID)
	_, err := makeRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	return nil
}

// Update project
//
// Example:
//		todoist.Token = "your token"
//		project, err := todoist.GetProject(1)
//		if err != nil {
//			panic(err)
//		}
//		project.Name = "updated"
//		err = project.Update()
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println(project)
func (p Project) Update() error {
	path := fmt.Sprintf("projects/%d", p.ID)
	project := struct {
		Name string `json:"name"`
	}{
		p.Name,
	}

	_, err := makeRequest(http.MethodPost, path, project)
	if err != nil {
		return err
	}

	return nil
}
