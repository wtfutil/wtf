package gitlab

import (
	glb "github.com/xanzy/go-gitlab"
)

type context struct {
	client *glb.Client
	user   *glb.User
}

func newContext(settings *Settings) (*context, error) {
	baseURL := settings.domain
	gitlabClient, _ := glb.NewClient(settings.apiKey, glb.WithBaseURL(baseURL))

	user, _, err := gitlabClient.Users.CurrentUser()

	if err != nil {
		return nil, err
	}

	ctx := &context{
		client: gitlabClient,
		user:   user,
	}

	return ctx, nil
}

type GitlabProject struct {
	context *context
	path    string

	MergeRequests         []*glb.MergeRequest
	AssignedMergeRequests []*glb.MergeRequest
	AuthoredMergeRequests []*glb.MergeRequest
	AssignedIssues        []*glb.Issue
	AuthoredIssues        []*glb.Issue
	RemoteProject         *glb.Project
}

func NewGitlabProject(context *context, projectPath string) *GitlabProject {
	project := GitlabProject{
		context: context,
		path:    projectPath,
	}

	return &project
}

// Refresh reloads the gitlab data via the Gitlab API
func (project *GitlabProject) Refresh() {
	project.MergeRequests, _ = project.loadMergeRequests()
	project.AssignedMergeRequests, _ = project.loadAssignedMergeRequests()
	project.AuthoredMergeRequests, _ = project.loadAuthoredMergeRequests()
	project.AssignedIssues, _ = project.loadAssignedIssues()
	project.AuthoredIssues, _ = project.loadAuthoredIssues()
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

// myMergeRequests returns a list of merge requests
func (project *GitlabProject) myMergeRequests() []*glb.MergeRequest {
	return project.AuthoredMergeRequests
}

// myAssignedMergeRequests returns a list of merge requests
// assigned
func (project *GitlabProject) myAssignedMergeRequests() []*glb.MergeRequest {
	return project.AssignedMergeRequests
}

// myAssignedIssues returns a list of issues
func (project *GitlabProject) myAssignedIssues() []*glb.Issue {
	return project.AssignedIssues
}

// myIssues returns a list of issues
func (project *GitlabProject) myIssues() []*glb.Issue {
	return project.AuthoredIssues
}

func (project *GitlabProject) loadMergeRequests() ([]*glb.MergeRequest, error) {
	state := "opened"
	opts := glb.ListProjectMergeRequestsOptions{
		State: &state,
	}

	mrs, _, err := project.context.client.MergeRequests.ListProjectMergeRequests(project.path, &opts)

	if err != nil {
		return nil, err
	}

	return mrs, nil
}

func (project *GitlabProject) loadAssignedMergeRequests() ([]*glb.MergeRequest, error) {
	state := "opened"
	opts := glb.ListProjectMergeRequestsOptions{
		State:      &state,
		AssigneeID: glb.AssigneeID(project.context.user.ID),
	}

	mrs, _, err := project.context.client.MergeRequests.ListProjectMergeRequests(project.path, &opts)

	if err != nil {
		return nil, err
	}

	return mrs, nil
}

func (project *GitlabProject) loadAuthoredMergeRequests() ([]*glb.MergeRequest, error) {
	state := "opened"
	opts := glb.ListProjectMergeRequestsOptions{
		State:    &state,
		AuthorID: &project.context.user.ID,
	}

	mrs, _, err := project.context.client.MergeRequests.ListProjectMergeRequests(project.path, &opts)

	if err != nil {
		return nil, err
	}

	return mrs, nil
}

func (project *GitlabProject) loadAssignedIssues() ([]*glb.Issue, error) {
	state := "opened"
	opts := glb.ListProjectIssuesOptions{
		State:      &state,
		AssigneeID: glb.AssigneeID(project.context.user.ID),
	}

	issues, _, err := project.context.client.Issues.ListProjectIssues(project.path, &opts)

	if err != nil {
		return nil, err
	}

	return issues, nil
}

func (project *GitlabProject) loadAuthoredIssues() ([]*glb.Issue, interface{}) {
	state := "opened"
	opts := glb.ListProjectIssuesOptions{
		State:    &state,
		AuthorID: &project.context.user.ID,
	}

	issues, _, err := project.context.client.Issues.ListProjectIssues(project.path, &opts)

	if err != nil {
		return nil, err
	}

	return issues, nil
}

func (project *GitlabProject) loadRemoteProject() (*glb.Project, error) {
	projectsitory, _, err := project.context.client.Projects.GetProject(project.path, nil)

	if err != nil {
		return nil, err
	}

	return projectsitory, nil
}
