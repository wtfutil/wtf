package gitlab

import (
	glab "gitlab.com/gitlab-org/api/client-go"
)

type context struct {
	client *glab.Client
	user   *glab.User
}

func newContext(settings *Settings) (*context, error) {
	baseURL := settings.domain
	gitlabClient, _ := glab.NewClient(settings.apiKey, glab.WithBaseURL(baseURL))

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

	MergeRequests         []*glab.MergeRequest
	AssignedMergeRequests []*glab.MergeRequest
	AuthoredMergeRequests []*glab.MergeRequest
	AssignedIssues        []*glab.Issue
	AuthoredIssues        []*glab.Issue
	RemoteProject         *glab.Project
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
func (project *GitlabProject) myMergeRequests() []*glab.MergeRequest {
	return project.AuthoredMergeRequests
}

// myAssignedMergeRequests returns a list of merge requests
// assigned
func (project *GitlabProject) myAssignedMergeRequests() []*glab.MergeRequest {
	return project.AssignedMergeRequests
}

// myAssignedIssues returns a list of issues
func (project *GitlabProject) myAssignedIssues() []*glab.Issue {
	return project.AssignedIssues
}

// myIssues returns a list of issues
func (project *GitlabProject) myIssues() []*glab.Issue {
	return project.AuthoredIssues
}

func (project *GitlabProject) loadMergeRequests() ([]*glab.MergeRequest, error) {
	state := "opened"
	opts := glab.ListProjectMergeRequestsOptions{
		State: &state,
	}

	mrs, _, err := project.context.client.MergeRequests.ListProjectMergeRequests(project.path, &opts)

	if err != nil {
		return nil, err
	}

	return mrs, nil
}

func (project *GitlabProject) loadAssignedMergeRequests() ([]*glab.MergeRequest, error) {
	state := "opened"
	opts := glab.ListProjectMergeRequestsOptions{
		State:      &state,
		AssigneeID: glab.AssigneeID(project.context.user.ID),
	}

	mrs, _, err := project.context.client.MergeRequests.ListProjectMergeRequests(project.path, &opts)

	if err != nil {
		return nil, err
	}

	return mrs, nil
}

func (project *GitlabProject) loadAuthoredMergeRequests() ([]*glab.MergeRequest, error) {
	state := "opened"
	opts := glab.ListProjectMergeRequestsOptions{
		State:    &state,
		AuthorID: &project.context.user.ID,
	}

	mrs, _, err := project.context.client.MergeRequests.ListProjectMergeRequests(project.path, &opts)

	if err != nil {
		return nil, err
	}

	return mrs, nil
}

func (project *GitlabProject) loadAssignedIssues() ([]*glab.Issue, error) {
	state := "opened"
	opts := glab.ListProjectIssuesOptions{
		State:      &state,
		AssigneeID: glab.AssigneeID(project.context.user.ID),
	}

	issues, _, err := project.context.client.Issues.ListProjectIssues(project.path, &opts)

	if err != nil {
		return nil, err
	}

	return issues, nil
}

func (project *GitlabProject) loadAuthoredIssues() ([]*glab.Issue, interface{}) {
	state := "opened"
	opts := glab.ListProjectIssuesOptions{
		State:    &state,
		AuthorID: &project.context.user.ID,
	}

	issues, _, err := project.context.client.Issues.ListProjectIssues(project.path, &opts)

	if err != nil {
		return nil, err
	}

	return issues, nil
}

func (project *GitlabProject) loadRemoteProject() (*glab.Project, error) {
	projectsitory, _, err := project.context.client.Projects.GetProject(project.path, nil)

	if err != nil {
		return nil, err
	}

	return projectsitory, nil
}
