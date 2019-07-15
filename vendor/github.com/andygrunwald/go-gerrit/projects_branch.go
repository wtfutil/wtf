package gerrit

import (
	"fmt"
	"net/url"
)

// BranchInfo entity contains information about a branch.
type BranchInfo struct {
	Ref       string        `json:"ref"`
	Revision  string        `json:"revision"`
	CanDelete bool          `json:"can_delete"`
	WebLinks  []WebLinkInfo `json:"web_links,omitempty"`
}

// BranchInput entity contains information for the creation of a new branch.
type BranchInput struct {
	Ref      string `json:"ref,omitempty"`
	Revision string `json:"revision,omitempty"`
}

// DeleteBranchesInput entity contains information about branches that should be deleted.
type DeleteBranchesInput struct {
	Branches []string `json:"DeleteBranchesInput"`
}

// BranchOptions specifies the parameters to the branch API endpoints.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#branch-options
type BranchOptions struct {
	// Limit the number of branches to be included in the results.
	Limit int `url:"n,omitempty"`

	// Skip the given number of branches from the beginning of the list.
	Skip string `url:"s,omitempty"`

	// Substring limits the results to those projects that match the specified substring.
	Substring string `url:"m,omitempty"`

	// Limit the results to those branches that match the specified regex.
	// Boundary matchers '^' and '$' are implicit.
	// For example: the regex 't*' will match any branches that start with 'test' and regex '*t' will match any branches that end with 'test'.
	Regex string `url:"r,omitempty"`
}

// ListBranches list the branches of a project.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#list-branches
func (s *ProjectsService) ListBranches(projectName string, opt *BranchOptions) (*[]BranchInfo, *Response, error) {
	u := fmt.Sprintf("projects/%s/branches/", url.QueryEscape(projectName))

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]BranchInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetBranch retrieves a branch of a project.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#get-branch
func (s *ProjectsService) GetBranch(projectName, branchID string) (*BranchInfo, *Response, error) {
	u := fmt.Sprintf("projects/%s/branches/%s", url.QueryEscape(projectName), url.QueryEscape(branchID))

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(BranchInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetReflog gets the reflog of a certain branch.
// The caller must be project owner.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#get-reflog
func (s *ProjectsService) GetReflog(projectName, branchID string) (*[]ReflogEntryInfo, *Response, error) {
	u := fmt.Sprintf("projects/%s/branches/%s/reflog", url.QueryEscape(projectName), url.QueryEscape(branchID))

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]ReflogEntryInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// CreateBranch creates a new branch.
// In the request body additional data for the branch can be provided as BranchInput.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#create-branch
func (s *ProjectsService) CreateBranch(projectName, branchID string, input *BranchInput) (*BranchInfo, *Response, error) {
	u := fmt.Sprintf("projects/%s/branches/%s", url.QueryEscape(projectName), url.QueryEscape(branchID))

	req, err := s.client.NewRequest("PUT", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(BranchInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// DeleteBranch deletes a branch.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#delete-branch
func (s *ProjectsService) DeleteBranch(projectName, branchID string) (*Response, error) {
	u := fmt.Sprintf("projects/%s/branches/%s", url.QueryEscape(projectName), url.QueryEscape(branchID))
	return s.client.DeleteRequest(u, nil)
}

// DeleteBranches delete one or more branches.
// If some branches could not be deleted, the response is “409 Conflict” and the error message is contained in the response body.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#delete-branches
func (s *ProjectsService) DeleteBranches(projectName string, input *DeleteBranchesInput) (*Response, error) {
	u := fmt.Sprintf("projects/%s/branches:delete", url.QueryEscape(projectName))
	return s.client.DeleteRequest(u, input)
}

// GetBranchContent gets the content of a file from the HEAD revision of a certain branch.
// The content is returned as base64 encoded string.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#get-content
func (s *ProjectsService) GetBranchContent(projectName, branchID, fileID string) (string, *Response, error) {
	u := fmt.Sprintf("projects/%s/branches/%s/files/%s/content", url.QueryEscape(projectName), url.QueryEscape(branchID), fileID)
	return getStringResponseWithoutOptions(s.client, u)
}
