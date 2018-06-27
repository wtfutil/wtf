package gerrit

import (
	glb "github.com/andygrunwald/go-gerrit"
)

type GerritProject struct {
	gerrit *glb.Client
	Path   string

	Changes *[]glb.ChangeInfo
}

func NewGerritProject(path string, gerrit *glb.Client) *GerritProject {
	project := GerritProject{
		gerrit: gerrit,
		Path:   path,
	}

	return &project
}

// Refresh reloads the gerrit data via the Gerrit API
func (project *GerritProject) Refresh() {
	project.Changes, _ = project.loadChanges()
}

/* -------------------- Counts -------------------- */

func (project *GerritProject) IssueCount() int {
	if project.Changes == nil {
		return 0
	}

	return len(*project.Changes)
}

func (project *GerritProject) ReviewCount() int {
	if project.Changes == nil {
		return 0
	}

	return len(*project.Changes)
}

/* -------------------- Unexported Functions -------------------- */

// myOutgoingReviews returns a list of my outgoing reviews created by username on this project
func (project *GerritProject) myOutgoingReviews(username string) []glb.ChangeInfo {
	changes := []glb.ChangeInfo{}

	if project.Changes == nil {
		return changes
	}

	for _, change := range *project.Changes {
		user := change.Owner

		if user.Username == username {
			changes = append(changes, change)
		}
	}

	return changes
}

// myIncomingReviews returns a list of merge requests for which username has been requested to ChangeInfo
func (project *GerritProject) myIncomingReviews(username string) []glb.ChangeInfo {
	changes := []glb.ChangeInfo{}

	if project.Changes == nil {
		return changes
	}

	for _, change := range *project.Changes {
		reviewers := change.Reviewers

		for _, reviewer := range reviewers["REVIEWER"] {
			if reviewer.Username == username {
				changes = append(changes, change)
			}
		}
	}

	return changes
}

func (project *GerritProject) loadChanges() (*[]glb.ChangeInfo, error) {
	opt := &glb.QueryChangeOptions{}
	opt.Query = []string{"(projects:" + project.Path + "+ is:open + owner:self) " + " OR " +
		"(projects:" + project.Path + " + is:open + ((reviewer:self + -owner:self + -star:ignore) + OR + assignee:self))"}
	opt.AdditionalFields = []string{"DETAILED_LABELS", "DETAILED_ACCOUNTS"}
	changes, _, err := project.gerrit.Changes.QueryChanges(opt)

	if err != nil {
		return nil, err
	}

	return changes, err
}
