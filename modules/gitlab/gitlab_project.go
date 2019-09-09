package gitlab

import (
	glb "github.com/xanzy/go-gitlab"
)

type GitlabProject struct {
	client *glb.Client
	path   string

	MergeRequests []*glb.MergeRequest
	RemoteProject *glb.Project
}

func NewGitlabProject(projectPath string, client *glb.Client) *GitlabProject {
	project := GitlabProject{
		client: client,
		path:   projectPath,
	}

	return &project
}

// Refresh reloads the gitlab data via the Gitlab API
func (project *GitlabProject) Refresh() {
	project.MergeRequests, _ = project.loadMergeRequests()
	project.RemoteProject, _ = project.loadRemoteProject()
}

/* -------------------- Counts -------------------- */

func (project *GitlabProject) IssueCount() int {
	if project.RemoteProject == nil {
		return 0
	}

	return project.RemoteProject.OpenIssuesCount
}

func (project *GitlabProject) MergeRequestCount() int {
	return len(project.MergeRequests)
}

func (project *GitlabProject) StarCount() int {
	if project.RemoteProject == nil {
		return 0
	}

	return project.RemoteProject.StarCount
}

/* -------------------- Unexported Functions -------------------- */

// myMergeRequests returns a list of merge requests created by username on this project
func (project *GitlabProject) myMergeRequests(username string) []*glb.MergeRequest {
	mrs := []*glb.MergeRequest{}

	for _, mr := range project.MergeRequests {
		user := mr.Author

		if user.Username == username {
			mrs = append(mrs, mr)
		}
	}

	return mrs
}

// myApprovalRequests returns a list of merge requests for which username has been
// requested to approve
func (project *GitlabProject) myApprovalRequests(username string) []*glb.MergeRequest {
	mrs := []*glb.MergeRequest{}

	for _, mr := range project.MergeRequests {
		approvers, _, err := project.client.MergeRequests.GetMergeRequestApprovals(project.path, mr.IID)
		if err != nil {
			continue
		}
		for _, approver := range approvers.Approvers {
			if approver.User.Username == username {
				mrs = append(mrs, mr)
			}
		}
	}

	return mrs
}

func (project *GitlabProject) loadMergeRequests() ([]*glb.MergeRequest, error) {
	state := "opened"
	opts := glb.ListProjectMergeRequestsOptions{
		State: &state,
	}

	mrs, _, err := project.client.MergeRequests.ListProjectMergeRequests(project.path, &opts)

	if err != nil {
		return nil, err
	}

	return mrs, nil
}

func (project *GitlabProject) loadRemoteProject() (*glb.Project, error) {
	projectsitory, _, err := project.client.Projects.GetProject(project.path, nil)

	if err != nil {
		return nil, err
	}

	return projectsitory, nil
}
