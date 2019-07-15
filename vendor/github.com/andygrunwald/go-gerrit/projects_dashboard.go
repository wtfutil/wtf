package gerrit

import (
	"fmt"
	"net/url"
)

// DashboardSectionInfo entity contains information about a section in a dashboard.
type DashboardSectionInfo struct {
	Name  string `json:"name"`
	Query string `json:"query"`
}

// DashboardInput entity contains information to create/update a project dashboard.
type DashboardInput struct {
	ID            string `json:"id,omitempty"`
	CommitMessage string `json:"commit_message,omitempty"`
}

// DashboardInfo entity contains information about a project dashboard.
type DashboardInfo struct {
	ID              string                 `json:"id"`
	Project         string                 `json:"project"`
	DefiningProject string                 `json:"defining_project"`
	Ref             string                 `json:"ref"`
	Path            string                 `json:"path"`
	Description     string                 `json:"description,omitempty"`
	Foreach         string                 `json:"foreach,omitempty"`
	URL             string                 `json:"url"`
	Default         bool                   `json:"default"`
	Title           string                 `json:"title,omitempty"`
	Sections        []DashboardSectionInfo `json:"sections"`
}

// ListDashboards list custom dashboards for a project.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#list-dashboards
func (s *ProjectsService) ListDashboards(projectName string) (*[]DashboardInfo, *Response, error) {
	u := fmt.Sprintf("projects/%s/dashboards/", url.QueryEscape(projectName))

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]DashboardInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetDashboard list custom dashboards for a project.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#get-dashboard
func (s *ProjectsService) GetDashboard(projectName, dashboardName string) (*DashboardInfo, *Response, error) {
	u := fmt.Sprintf("projects/%s/dashboards/%s", url.QueryEscape(projectName), url.QueryEscape(dashboardName))

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(DashboardInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// SetDashboard updates/Creates a project dashboard.
// Currently only supported for the default dashboard.
//
// The creation/update information for the dashboard must be provided in the request body as a DashboardInput entity.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#set-dashboard
func (s *ProjectsService) SetDashboard(projectName, dashboardID string, input *DashboardInput) (*DashboardInfo, *Response, error) {
	u := fmt.Sprintf("projects/%s/dashboards/%s", url.QueryEscape(projectName), url.QueryEscape(dashboardID))

	req, err := s.client.NewRequest("PUT", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(DashboardInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// DeleteDashboard deletes a project dashboard.
// Currently only supported for the default dashboard.
//
// The request body does not need to include a DashboardInput entity if no commit message is specified.
// Please note that some proxies prohibit request bodies for DELETE requests.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#delete-dashboard
func (s *ProjectsService) DeleteDashboard(projectName, dashboardID string, input *DashboardInput) (*Response, error) {
	u := fmt.Sprintf("projects/%s/dashboards/%s", url.QueryEscape(projectName), url.QueryEscape(dashboardID))
	return s.client.DeleteRequest(u, input)
}
