package gerrit

import (
	"fmt"
)

// GroupsService contains Group related REST endpoints
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html
type GroupsService struct {
	client *Client
}

// GroupAuditEventInfo entity contains information about an audit event of a group.
type GroupAuditEventInfo struct {
	// TODO Member AccountInfo OR GroupInfo `json:"member"`
	Type string      `json:"type"`
	User AccountInfo `json:"user"`
	Date Timestamp   `json:"date"`
}

// GroupInfo entity contains information about a group.
// This can be a Gerrit internal group, or an external group that is known to Gerrit.
type GroupInfo struct {
	ID          string           `json:"id"`
	Name        string           `json:"name,omitempty"`
	URL         string           `json:"url,omitempty"`
	Options     GroupOptionsInfo `json:"options"`
	Description string           `json:"description,omitempty"`
	GroupID     int              `json:"group_id,omitempty"`
	Owner       string           `json:"owner,omitempty"`
	OwnerID     string           `json:"owner_id,omitempty"`
	CreatedOn   *Timestamp       `json:"created_on,omitempty"`
	Members     []AccountInfo    `json:"members,omitempty"`
	Includes    []GroupInfo      `json:"includes,omitempty"`
}

// GroupInput entity contains information for the creation of a new internal group.
type GroupInput struct {
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	VisibleToAll bool   `json:"visible_to_all,omitempty"`
	OwnerID      string `json:"owner_id,omitempty"`
}

// GroupOptionsInfo entity contains options of the group.
type GroupOptionsInfo struct {
	VisibleToAll bool `json:"visible_to_all,omitempty"`
}

// GroupOptionsInput entity contains new options for a group.
type GroupOptionsInput struct {
	VisibleToAll bool `json:"visible_to_all,omitempty"`
}

// GroupsInput entity contains information about groups that should be included into a group or that should be deleted from a group.
type GroupsInput struct {
	OneGroup string   `json:"_one_group,omitempty"`
	Groups   []string `json:"groups,omitempty"`
}

// ListGroupsOptions specifies the different options for the ListGroups call.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#list-groups
type ListGroupsOptions struct {
	// Group Options
	// Options fields can be obtained by adding o parameters, each option requires more lookups and slows down the query response time to the client so they are generally disabled by default.
	// Optional fields are:
	//	INCLUDES: include list of directly included groups.
	//	MEMBERS: include list of direct group members.
	Options []string `url:"o,omitempty"`

	// Check if a group is owned by the calling user
	// By setting the option owned and specifying a group to inspect with the option q, it is possible to find out, if this group is owned by the calling user.
	// If the group is owned by the calling user, the returned map contains this group. If the calling user doesn’t own this group an empty map is returned.
	Owned string `url:"owned,omitempty"`
	Group string `url:"q,omitempty"`

	// Group Limit
	// The /groups/ URL also accepts a limit integer in the n parameter. This limits the results to show n groups.
	Limit int `url:"n,omitempty"`
	// The /groups/ URL also accepts a start integer in the S parameter. The results will skip S groups from group list.
	Skip int `url:"S,omitempty"`
}

// ListGroups lists the groups accessible by the caller.
// This is the same as using the ls-groups command over SSH, and accepts the same options as query parameters.
// The entries in the map are sorted by group name.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#list-groups
func (s *GroupsService) ListGroups(opt *ListGroupsOptions) (*map[string]GroupInfo, *Response, error) {
	u := "groups/"

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(map[string]GroupInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetGroup retrieves a group.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#get-group
func (s *GroupsService) GetGroup(groupID string) (*GroupInfo, *Response, error) {
	u := fmt.Sprintf("groups/%s", groupID)
	return s.getGroupInfoResponse(u)
}

// GetGroupDetail retrieves a group with the direct members and the directly included groups.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#get-group-detail
func (s *GroupsService) GetGroupDetail(groupID string) (*GroupInfo, *Response, error) {
	u := fmt.Sprintf("groups/%s/detail", groupID)
	return s.getGroupInfoResponse(u)
}

// getGroupInfoResponse retrieved a single GroupInfo Response for a GET request
func (s *GroupsService) getGroupInfoResponse(u string) (*GroupInfo, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(GroupInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetGroupName retrieves the name of a group.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#get-group-name
func (s *GroupsService) GetGroupName(groupID string) (string, *Response, error) {
	u := fmt.Sprintf("groups/%s/name", groupID)
	return getStringResponseWithoutOptions(s.client, u)
}

// GetGroupDescription retrieves the description of a group.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#get-group-description
func (s *GroupsService) GetGroupDescription(groupID string) (string, *Response, error) {
	u := fmt.Sprintf("groups/%s/description", groupID)
	return getStringResponseWithoutOptions(s.client, u)
}

// GetGroupOptions retrieves the options of a group.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#get-group-options
func (s *GroupsService) GetGroupOptions(groupID string) (*GroupOptionsInfo, *Response, error) {
	u := fmt.Sprintf("groups/%s/options", groupID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(GroupOptionsInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetGroupOwner retrieves the owner group of a Gerrit internal group.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#get-group-owner
func (s *GroupsService) GetGroupOwner(groupID string) (*GroupInfo, *Response, error) {
	u := fmt.Sprintf("groups/%s/owner", groupID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(GroupInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetAuditLog gets the audit log of a Gerrit internal group.
// The returned audit events are sorted by date in reverse order so that the newest audit event comes first.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#get-audit-log
func (s *GroupsService) GetAuditLog(groupID string) (*[]GroupAuditEventInfo, *Response, error) {
	u := fmt.Sprintf("groups/%s/log.audit", groupID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]GroupAuditEventInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// CreateGroup creates a new Gerrit internal group.
// In the request body additional data for the group can be provided as GroupInput.
//
// As response the GroupInfo entity is returned that describes the created group.
// If the group creation fails because the name is already in use the response is “409 Conflict”.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#create-group
func (s *GroupsService) CreateGroup(groupID string, input *GroupInput) (*GroupInfo, *Response, error) {
	u := fmt.Sprintf("groups/%s", groupID)

	req, err := s.client.NewRequest("PUT", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(GroupInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// RenameGroup renames a Gerrit internal group.
// The new group name must be provided in the request body.
//
// As response the new group name is returned.
// If renaming the group fails because the new name is already in use the response is “409 Conflict”.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#rename-group
func (s *GroupsService) RenameGroup(groupID, name string) (*string, *Response, error) {
	u := fmt.Sprintf("groups/%s/name", groupID)
	input := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}

	req, err := s.client.NewRequest("PUT", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(string)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// SetGroupDescription sets the description of a Gerrit internal group.
// The new group description must be provided in the request body.
//
// As response the new group description is returned.
// If the description was deleted the response is “204 No Content”.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#set-group-description
func (s *GroupsService) SetGroupDescription(groupID, description string) (*string, *Response, error) {
	u := fmt.Sprintf("groups/%s/description", groupID)
	input := struct {
		Description string `json:"description"`
	}{
		Description: description,
	}

	req, err := s.client.NewRequest("PUT", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(string)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// DeleteGroupDescription deletes the description of a Gerrit internal group.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#delete-group-description
func (s *GroupsService) DeleteGroupDescription(groupID string) (*Response, error) {
	u := fmt.Sprintf("groups/%s/description'", groupID)
	return s.client.DeleteRequest(u, nil)
}

// SetGroupOptions sets the options of a Gerrit internal group.
// The new group options must be provided in the request body as a GroupOptionsInput entity.
//
// As response the new group options are returned as a GroupOptionsInfo entity.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#set-group-options
func (s *GroupsService) SetGroupOptions(groupID string, input *GroupOptionsInput) (*GroupOptionsInfo, *Response, error) {
	u := fmt.Sprintf("groups/%s/options", groupID)

	req, err := s.client.NewRequest("PUT", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(GroupOptionsInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// SetGroupOwner sets the owner group of a Gerrit internal group.
// The new owner group must be provided in the request body.
// The new owner can be specified by name, by group UUID or by the legacy numeric group ID.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#set-group-owner
func (s *GroupsService) SetGroupOwner(groupID, owner string) (*GroupInfo, *Response, error) {
	u := fmt.Sprintf("groups/%s/owner", groupID)
	input := struct {
		Owner string `json:"owner"`
	}{
		Owner: owner,
	}

	req, err := s.client.NewRequest("PUT", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(GroupInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}
