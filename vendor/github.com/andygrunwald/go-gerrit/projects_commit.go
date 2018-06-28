package gerrit

import (
	"fmt"
	"net/url"
)

// GetCommit retrieves a commit of a project.
// The commit must be visible to the caller.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#get-commit
func (s *ProjectsService) GetCommit(projectName, commitID string) (*CommitInfo, *Response, error) {
	u := fmt.Sprintf("projects/%s/commits/%s", url.QueryEscape(projectName), commitID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(CommitInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetCommitContent gets the content of a file from the HEAD revision of a certain branch.
// The content is returned as base64 encoded string.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#get-content
func (s *ProjectsService) GetCommitContent(projectName, branchID, fileID string) (string, *Response, error) {
	u := fmt.Sprintf("projects/%s/branches/%s/files/%s/content", url.QueryEscape(projectName), branchID, fileID)
	return getStringResponseWithoutOptions(s.client, u)
}
