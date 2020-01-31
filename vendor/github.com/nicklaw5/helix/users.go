package helix

import "time"

// User ...
type User struct {
	ID              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	Type            string `json:"type"`
	BroadcasterType string `json:"broadcaster_type"`
	Description     string `json:"description"`
	ProfileImageURL string `json:"profile_image_url"`
	OfflineImageURL string `json:"offline_image_url"`
	ViewCount       int    `json:"view_count"`
	Email           string `json:"email"`
}

// ManyUsers ...
type ManyUsers struct {
	Users []User `json:"data"`
}

// UsersResponse ...
type UsersResponse struct {
	ResponseCommon
	Data ManyUsers
}

// UsersParams ...
type UsersParams struct {
	IDs    []string `query:"id"`    // Limit 100
	Logins []string `query:"login"` // Limit 100
}

// GetUsers gets information about one or more specified Twitch users.
// Users are identified by optional user IDs and/or login name. If neither
// a user ID nor a login name is specified, the user is looked up by Bearer token.
//
// Optional scope: user:read:email
func (c *Client) GetUsers(params *UsersParams) (*UsersResponse, error) {
	resp, err := c.get("/users", &ManyUsers{}, params)
	if err != nil {
		return nil, err
	}

	users := &UsersResponse{}
	users.StatusCode = resp.StatusCode
	users.Header = resp.Header
	users.Error = resp.Error
	users.ErrorStatus = resp.ErrorStatus
	users.ErrorMessage = resp.ErrorMessage
	users.Data.Users = resp.Data.(*ManyUsers).Users

	return users, nil
}

// UpdateUserParams ...
type UpdateUserParams struct {
	Description string `query:"description"`
}

// UpdateUser updates the description of a user specified
// by a Bearer token.
//
// Required scope: user:edit
func (c *Client) UpdateUser(params *UpdateUserParams) (*UsersResponse, error) {
	resp, err := c.put("/users", &ManyUsers{}, params)
	if err != nil {
		return nil, err
	}

	users := &UsersResponse{}
	users.StatusCode = resp.StatusCode
	users.Header = resp.Header
	users.Error = resp.Error
	users.ErrorStatus = resp.ErrorStatus
	users.ErrorMessage = resp.ErrorMessage
	users.Data.Users = resp.Data.(*ManyUsers).Users

	return users, nil
}

// UserFollow ...
type UserFollow struct {
	FromID     string    `json:"from_id"`
	FromName   string    `json:"from_name"`
	ToID       string    `json:"to_id"`
	ToName     string    `json:"to_name"`
	FollowedAt time.Time `json:"followed_at"`
}

// ManyFollows ...
type ManyFollows struct {
	Total      int          `json:"total"`
	Follows    []UserFollow `json:"data"`
	Pagination Pagination   `json:"pagination"`
}

// UsersFollowsResponse ...
type UsersFollowsResponse struct {
	ResponseCommon
	Data ManyFollows
}

// UsersFollowsParams ...
type UsersFollowsParams struct {
	After  string `query:"after"`
	First  int    `query:"first,20"` // Limit 100
	FromID string `query:"from_id"`
	ToID   string `query:"to_id"`
}

// GetUsersFollows gets information on follow relationships between two Twitch users.
// Information returned is sorted in order, most recent follow first. This can return
// information like “who is lirik following,” “who is following lirik,” or “is user X
// following user Y.”
func (c *Client) GetUsersFollows(params *UsersFollowsParams) (*UsersFollowsResponse, error) {
	resp, err := c.get("/users/follows", &ManyFollows{}, params)
	if err != nil {
		return nil, err
	}

	users := &UsersFollowsResponse{}
	users.StatusCode = resp.StatusCode
	users.Header = resp.Header
	users.Error = resp.Error
	users.ErrorStatus = resp.ErrorStatus
	users.ErrorMessage = resp.ErrorMessage
	users.Data.Total = resp.Data.(*ManyFollows).Total
	users.Data.Follows = resp.Data.(*ManyFollows).Follows
	users.Data.Pagination = resp.Data.(*ManyFollows).Pagination

	return users, nil
}
