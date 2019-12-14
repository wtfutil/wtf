// Copyright 2019 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import "context"

// createOrgRequest is a subset of Organization and is used internally
// by CreateOrg to pass only the known fields for the endpoint.
type createOrgRequest struct {
	Login *string `json:"login,omitempty"`
	Admin *string `json:"admin,omitempty"`
}

// CreateOrg creates a new organization in GitHub Enterprise.
//
// Note that only a subset of the org fields are used and org must
// not be nil.
//
// GitHub Enterprise API docs: https://developer.github.com/enterprise/v3/enterprise-admin/orgs/#create-an-organization
func (s *AdminService) CreateOrg(ctx context.Context, org *Organization, admin string) (*Organization, *Response, error) {
	u := "admin/organizations"

	orgReq := &createOrgRequest{
		Login: org.Login,
		Admin: &admin,
	}

	req, err := s.client.NewRequest("POST", u, orgReq)
	if err != nil {
		return nil, nil, err
	}

	o := new(Organization)
	resp, err := s.client.Do(ctx, req, o)
	if err != nil {
		return nil, resp, err
	}

	return o, resp, nil
}
