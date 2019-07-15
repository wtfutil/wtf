// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"fmt"
)

// Organization represents a Trello organization or team, i.e. a collection of members and boards.
// https://developers.trello.com/reference/#organizations
type Organization struct {
	client      *Client
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	DisplayName string   `json:"displayName"`
	Desc        string   `json:"desc"`
	URL         string   `json:"url"`
	Website     string   `json:"website"`
	Products    []string `json:"products"`
	PowerUps    []string `json:"powerUps"`
}

// GetOrganization takes an organization id and Arguments and either
// GETs returns an Organization, or an error.
func (c *Client) GetOrganization(orgID string, args Arguments) (organization *Organization, err error) {
	path := fmt.Sprintf("organizations/%s", orgID)
	err = c.Get(path, args, &organization)
	if organization != nil {
		organization.client = c
	}
	return
}
