package gitlab

import (
	glb "github.com/xanzy/go-gitlab"
)

type GitlabProject struct {
	gitlab *glb.Client
	Path   string

	MergeRequests []*glb.MergeRequest
	RemoteProject *glb.Project
}

func NewGitlabProject(name string, namespace string, gitlab *glb.Client) *GitlabProject {
	path := namespace + "/" + name
	project := GitlabProject{
		gitlab: gitlab,
		Path:   path,
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
		approvers, _, err := project.gitlab.MergeRequests.GetMergeRequestApprovals(project.Path, mr.IID)
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

	mrs, _, err := project.gitlab.MergeRequests.ListProjectMergeRequests(project.Path, &opts)

	if err != nil {
		return nil, err
	}

	return mrs, nil
}

func (project *GitlabProject) loadRemoteProject() (*glb.Project, error) {
	projectsitory, _, err := project.gitlab.Projects.GetProject(project.Path)

	if err != nil {
		return nil, err
	}

	return projectsitory, nil
}
