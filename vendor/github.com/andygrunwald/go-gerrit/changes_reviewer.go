package gerrit

import (
	"fmt"
)

// ReviewerInfo entity contains information about a reviewer and its votes on a change.
type ReviewerInfo struct {
	AccountInfo
	Approvals map[string]string `json:"approvals"`
}

// SuggestedReviewerInfo entity contains information about a reviewer that can be added to a change (an account or a group).
type SuggestedReviewerInfo struct {
	Account AccountInfo   `json:"account,omitempty"`
	Group   GroupBaseInfo `json:"group,omitempty"`
}

// AddReviewerResult entity describes the result of adding a reviewer to a change.
type AddReviewerResult struct {
	Input     string         `json:"input,omitempty"`
	Reviewers []ReviewerInfo `json:"reviewers,omitempty"`
	CCS       []ReviewerInfo `json:"ccs,omitempty"`
	Error     string         `json:"error,omitempty"`
	Confirm   bool           `json:"confirm,omitempty"`
}

// DeleteVoteInput entity contains options for the deletion of a vote.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#delete-vote-input
type DeleteVoteInput struct {
	Label         string                `json:"label,omitempty"`
	Notify        string                `json:"notify,omitempty"`
	NotifyDetails map[string]NotifyInfo `json:"notify_details"`
}

// ListReviewers lists the reviewers of a change.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-reviewers
func (s *ChangesService) ListReviewers(changeID string) (*[]ReviewerInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/reviewers/", changeID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]ReviewerInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// SuggestReviewers suggest the reviewers for a given query q and result limit n.
// If result limit is not passed, then the default 10 is used.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#suggest-reviewers
func (s *ChangesService) SuggestReviewers(changeID string, opt *QueryOptions) (*[]SuggestedReviewerInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/suggest_reviewers", changeID)

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]SuggestedReviewerInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetReviewer retrieves a reviewer of a change.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-reviewer
func (s *ChangesService) GetReviewer(changeID, accountID string) (*ReviewerInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/reviewers/%s", changeID, accountID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(ReviewerInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// AddReviewer adds one user or all members of one group as reviewer to the change.
// The reviewer to be added to the change must be provided in the request body as a ReviewerInput entity.
//
// As response an AddReviewerResult entity is returned that describes the newly added reviewers.
// If a group is specified, adding the group members as reviewers is an atomic operation.
// This means if an error is returned, none of the members are added as reviewer.
// If a group with many members is added as reviewer a confirmation may be required.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#add-reviewer
func (s *ChangesService) AddReviewer(changeID string, input *ReviewerInput) (*AddReviewerResult, *Response, error) {
	u := fmt.Sprintf("changes/%s/reviewers", changeID)

	req, err := s.client.NewRequest("POST", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(AddReviewerResult)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// DeleteReviewer deletes a reviewer from a change.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#delete-reviewer
func (s *ChangesService) DeleteReviewer(changeID, accountID string) (*Response, error) {
	u := fmt.Sprintf("changes/%s/reviewers/%s", changeID, accountID)
	return s.client.DeleteRequest(u, nil)
}

// ListVotes lists the votes for a specific reviewer of the change.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-votes
func (s *ChangesService) ListVotes(changeID string, accountID string) (map[string]int, *Response, error) {
	u := fmt.Sprintf("changes/%s/reviewers/%s/votes/", changeID, accountID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var v map[string]int
	resp, err := s.client.Do(req, &v)
	if err != nil {
		return nil, resp, err
	}
	return v, resp, err
}

// DeleteVote deletes a single vote from a change. Note, that even when the
// last vote of a reviewer is removed the reviewer itself is still listed on
// the change.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#delete-vote
func (s *ChangesService) DeleteVote(changeID string, accountID string, label string, input *DeleteVoteInput) (*Response, error) {
	u := fmt.Sprintf("changes/%s/reviewers/%s/votes/%s", changeID, accountID, label)
	return s.client.DeleteRequest(u, input)
}
