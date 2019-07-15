package gerrit

import (
	"fmt"
	"net/url"
)

// ChildProjectOptions specifies the parameters to the Child Project API endpoints.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#list-child-projects
type ChildProjectOptions struct {
	// Recursive resolve the child projects of a project recursively.
	// Child projects that are not visible to the calling user are ignored and are not resolved further.
	Recursive int `url:"recursive,omitempty"`
}

// ListChildProjects lists the direct child projects of a project.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#list-child-projects
func (s *ProjectsService) ListChildProjects(projectName string, opt *ChildProjectOptions) (*[]ProjectInfo, *Response, error) {
	u := fmt.Sprintf("projects/%s/children/", url.QueryEscape(projectName))

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]ProjectInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetChildProject retrieves a child project.
// If a non-direct child project should be retrieved the parameter recursive must be set.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#get-child-project
func (s *ProjectsService) GetChildProject(projectName, childProjectName string, opt *ChildProjectOptions) (*ProjectInfo, *Response, error) {
	u := fmt.Sprintf("projects/%s/children/%s", url.QueryEscape(projectName), url.QueryEscape(childProjectName))

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(ProjectInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}
