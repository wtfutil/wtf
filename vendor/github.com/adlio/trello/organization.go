// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"fmt"
)

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

func (c *Client) GetOrganization(orgID string, args Arguments) (organization *Organization, err error) {
	path := fmt.Sprintf("organizations/%s", orgID)
	err = c.Get(path, args, &organization)
	if organization != nil {
		organization.client = c
	}
	return
}
