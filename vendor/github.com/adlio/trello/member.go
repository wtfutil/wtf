// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"fmt"
)

// Member represents a Trello member.
// https://developers.trello.com/reference/#member-object
type Member struct {
	client     *Client
	ID         string `json:"id"`
	Username   string `json:"username"`
	FullName   string `json:"fullName"`
	Initials   string `json:"initials"`
	AvatarHash string `json:"avatarHash"`
	Email      string `json:"email"`
}

// GetMember takes a member id and Arguments and returns a Member or an error.
func (c *Client) GetMember(memberID string, args Arguments) (member *Member, err error) {
	path := fmt.Sprintf("members/%s", memberID)
	err = c.Get(path, args, &member)
	if err == nil {
		member.client = c
	}
	return
}

// GetMembers takes Arguments and returns a slice of all members of the organization or an error.
func (o *Organization) GetMembers(args Arguments) (members []*Member, err error) {
	path := fmt.Sprintf("organizations/%s/members", o.ID)
	err = o.client.Get(path, args, &members)
	for i := range members {
		members[i].client = o.client
	}
	return
}

// GetMembers takes Arguments and returns a slice of all members of the Board or an error.
func (b *Board) GetMembers(args Arguments) (members []*Member, err error) {
	path := fmt.Sprintf("boards/%s/members", b.ID)
	err = b.client.Get(path, args, &members)
	for i := range members {
		members[i].client = b.client
	}
	return
}

// GetMembers takes Arguments and returns a slice of all members of the Card or an error.
func (c *Card) GetMembers(args Arguments) (members []*Member, err error) {
	path := fmt.Sprintf("cards/%s/members", c.ID)
	err = c.client.Get(path, args, &members)
	for i := range members {
		members[i].client = c.client
	}
	return
}
