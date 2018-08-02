package gerrit

import (
	glb "github.com/andygrunwald/go-gerrit"
	"github.com/senorprogrammer/wtf/wtf"
)

type GerritProject struct {
	gerrit *glb.Client
	Path   string

	Changes         *[]glb.ChangeInfo
	ReviewCount     int
	IncomingReviews []glb.ChangeInfo
	OutgoingReviews []glb.ChangeInfo
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
	username := wtf.Config.UString("wtf.mods.gerrit.username")
	project.Changes, _ = project.loadChanges()

	project.ReviewCount = project.countReviews(project.Changes)
	project.IncomingReviews = project.myIncomingReviews(project.Changes, username)
	project.OutgoingReviews = project.myOutgoingReviews(project.Changes, username)

}

/* -------------------- Counts -------------------- */

func (project *GerritProject) countReviews(changes *[]glb.ChangeInfo) int {
	if changes == nil {
		return 0
	}

	return len(*changes)
}

/* -------------------- Unexported Functions -------------------- */

// myOutgoingReviews returns a list of my outgoing reviews created by username on this project
func (project *GerritProject) myOutgoingReviews(changes *[]glb.ChangeInfo, username string) []glb.ChangeInfo {
	var ors []glb.ChangeInfo

	if changes == nil {
		return ors
	}

	for _, change := range *changes {
		user := change.Owner

		if user.Username == username {
			ors = append(ors, change)
		}
	}

	return ors
}

// myIncomingReviews returns a list of merge requests for which username has been requested to ChangeInfo
func (project *GerritProject) myIncomingReviews(changes *[]glb.ChangeInfo, username string) []glb.ChangeInfo {
	var irs []glb.ChangeInfo

	if changes == nil {
		return irs
	}

	for _, change := range *changes {
		reviewers := change.Reviewers

		for _, reviewer := range reviewers["REVIEWER"] {
			if reviewer.Username == username {
				irs = append(irs, change)
			}
		}
	}

	return irs
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
