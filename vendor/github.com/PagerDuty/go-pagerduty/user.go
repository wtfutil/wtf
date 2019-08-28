package pagerduty

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// NotificationRule is a rule for notifying the user.
type NotificationRule struct {
	ID                  string        `json:"id"`
	StartDelayInMinutes uint          `json:"start_delay_in_minutes"`
	CreatedAt           string        `json:"created_at"`
	ContactMethod       ContactMethod `json:"contact_method"`
	Urgency             string        `json:"urgency"`
	Type                string        `json:"type"`
}

// User is a member of a PagerDuty account that has the ability to interact with incidents and other data on the account.
type User struct {
	APIObject
	Type              string             `json:"type"`
	Name              string             `json:"name"`
	Summary           string             `json:"summary"`
	Email             string             `json:"email"`
	Timezone          string             `json:"time_zone,omitempty"`
	Color             string             `json:"color,omitempty"`
	Role              string             `json:"role,omitempty"`
	AvatarURL         string             `json:"avatar_url,omitempty"`
	Description       string             `json:"description,omitempty"`
	InvitationSent    bool               `json:"invitation_sent,omitempty"`
	ContactMethods    []ContactMethod    `json:"contact_methods"`
	NotificationRules []NotificationRule `json:"notification_rules"`
	JobTitle          string             `json:"job_title,omitempty"`
	Teams             []Team
}

// ContactMethod is a way of contacting the user.
type ContactMethod struct {
	ID             string `json:"id"`
	Type           string `json:"type"`
	Summary        string `json:"summary"`
	Self           string `json:"self"`
	Label          string `json:"label"`
	Address        string `json:"address"`
	SendShortEmail bool   `json:"send_short_email,omitempty"`
	SendHTMLEmail  bool   `json:"send_html_email,omitempty"`
	Blacklisted    bool   `json:"blacklisted,omitempty"`
	CountryCode    int    `json:"country_code,omitempty"`
	Enabled        bool   `json:"enabled,omitempty"`
	HTMLUrl        string `json:"html_url"`
}

// ListUsersResponse is the data structure returned from calling the ListUsers API endpoint.
type ListUsersResponse struct {
	APIListObject
	Users []User
}

// ListUsersOptions is the data structure used when calling the ListUsers API endpoint.
type ListUsersOptions struct {
	APIListObject
	Query    string   `url:"query,omitempty"`
	TeamIDs  []string `url:"team_ids,omitempty,brackets"`
	Includes []string `url:"include,omitempty,brackets"`
}

// ListContactMethodResponse is the data structure returned from calling the GetUserContactMethod API endpoint.
type ListContactMethodsResponse struct {
	APIListObject
	ContactMethods []ContactMethod `json:"contact_methods"`
}

// GetUserOptions is the data structure used when calling the GetUser API endpoint.
type GetUserOptions struct {
	Includes []string `url:"include,omitempty,brackets"`
}

// ListUsers lists users of your PagerDuty account, optionally filtered by a search query.
func (c *Client) ListUsers(o ListUsersOptions) (*ListUsersResponse, error) {
	v, err := query.Values(o)
	if err != nil {
		return nil, err
	}
	resp, err := c.get("/users?" + v.Encode())
	if err != nil {
		return nil, err
	}
	var result ListUsersResponse
	return &result, c.decodeJSON(resp, &result)
}

// CreateUser creates a new user.
func (c *Client) CreateUser(u User) (*User, error) {
	data := make(map[string]User)
	data["user"] = u
	resp, err := c.post("/users", data, nil)
	return getUserFromResponse(c, resp, err)
}

// DeleteUser deletes a user.
func (c *Client) DeleteUser(id string) error {
	_, err := c.delete("/users/" + id)
	return err
}

// GetUser gets details about an existing user.
func (c *Client) GetUser(id string, o GetUserOptions) (*User, error) {
	v, err := query.Values(o)
	if err != nil {
		return nil, err
	}
	resp, err := c.get("/users/" + id + "?" + v.Encode())
	return getUserFromResponse(c, resp, err)
}

// UpdateUser updates an existing user.
func (c *Client) UpdateUser(u User) (*User, error) {
	v := make(map[string]User)
	v["user"] = u
	resp, err := c.put("/users/"+u.ID, v, nil)
	return getUserFromResponse(c, resp, err)
}

func getUserFromResponse(c *Client, resp *http.Response, err error) (*User, error) {
	if err != nil {
		return nil, err
	}
	var target map[string]User
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	rootNode := "user"
	t, nodeOK := target[rootNode]
	if !nodeOK {
		return nil, fmt.Errorf("JSON response does not have %s field", rootNode)
	}
	return &t, nil
}

// ListUserContactMethod fetches contact methods of the existing user.
func (c *Client) ListUserContactMethods(userId string) (*ListContactMethodsResponse, error) {
	resp, err := c.get("/users/" + userId + "/contact_methods")
	if err != nil {
		return nil, err
	}
	var result ListContactMethodsResponse
	return &result, c.decodeJSON(resp, &result)
}

// GetContactMethod gets details about a contact method.
func (c *Client) GetUserContactMethod(userID, id string) (*ContactMethod, error) {
	resp, err := c.get("/users/" + userID + "/contact_methods/" + id)
	return getContactMethodFromResponse(c, resp, err)
}

func getContactMethodFromResponse(c *Client, resp *http.Response, err error) (*ContactMethod, error) {
	if err != nil {
		return nil, err
	}
	var target map[string]ContactMethod
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	rootNode := "contact_method"
	t, nodeOK := target[rootNode]
	if !nodeOK {
		return nil, fmt.Errorf("JSON response does not have %s field", rootNode)
	}
	return &t, nil
}
