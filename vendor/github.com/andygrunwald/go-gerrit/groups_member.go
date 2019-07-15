package gerrit

import (
	"fmt"
)

// ListGroupMembersOptions specifies the different options for the ListGroupMembers call.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#group-members
type ListGroupMembersOptions struct {
	// To resolve the included groups of a group recursively and to list all members the parameter recursive can be set.
	// Members from included external groups and from included groups which are not visible to the calling user are ignored.
	Recursive bool `url:"recursive,omitempty"`
}

// MembersInput entity contains information about accounts that should be added as members to a group or that should be deleted from the group
type MembersInput struct {
	OneMember string   `json:"_one_member,omitempty"`
	Members   []string `json:"members,omitempty"`
}

// ListGroupMembers lists the direct members of a Gerrit internal group.
// The entries in the list are sorted by full name, preferred email and id.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#group-members
func (s *GroupsService) ListGroupMembers(groupID string, opt *ListGroupMembersOptions) (*[]AccountInfo, *Response, error) {
	u := fmt.Sprintf("groups/%s/members/", groupID)

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]AccountInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetGroupMember retrieves a group member.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#get-group-member
func (s *GroupsService) GetGroupMember(groupID, accountID string) (*AccountInfo, *Response, error) {
	u := fmt.Sprintf("groups/%s/members/%s", groupID, accountID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(AccountInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// AddGroupMember adds a user as member to a Gerrit internal group.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#add-group-member
func (s *GroupsService) AddGroupMember(groupID, accountID string) (*AccountInfo, *Response, error) {
	u := fmt.Sprintf("groups/%s/members/%s", groupID, accountID)

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(AccountInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// AddGroupMembers adds one or several users to a Gerrit internal group.
// The users to be added to the group must be provided in the request body as a MembersInput entity.
//
// As response a list of detailed AccountInfo entities is returned that describes the group members that were specified in the MembersInput.
// An AccountInfo entity is returned for each user specified in the input, independently of whether the user was newly added to the group or whether the user was already a member of the group.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#_add_group_members
func (s *GroupsService) AddGroupMembers(groupID string, input *MembersInput) (*[]AccountInfo, *Response, error) {
	u := fmt.Sprintf("groups/%s/members", groupID)

	req, err := s.client.NewRequest("POST", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new([]AccountInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// DeleteGroupMember deletes a user from a Gerrit internal group.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#delete-group-member
func (s *GroupsService) DeleteGroupMember(groupID, accountID string) (*Response, error) {
	u := fmt.Sprintf("groups/%s/members/%s'", groupID, accountID)
	return s.client.DeleteRequest(u, nil)
}

// DeleteGroupMembers delete one or several users from a Gerrit internal group.
// The users to be deleted from the group must be provided in the request body as a MembersInput entity.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#delete-group-members
func (s *GroupsService) DeleteGroupMembers(groupID string, input *MembersInput) (*Response, error) {
	u := fmt.Sprintf("groups/%s/members.delete'", groupID)

	req, err := s.client.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
