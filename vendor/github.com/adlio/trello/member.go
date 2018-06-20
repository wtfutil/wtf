// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"fmt"
)

type Member struct {
	client     *Client
	ID         string `json:"id"`
	Username   string `json:"username"`
	FullName   string `json:"fullName"`
	Initials   string `json:"initials"`
	AvatarHash string `json:"avatarHash"`
	Email      string `json:"email"`
}

func (c *Client) GetMember(memberID string, args Arguments) (member *Member, err error) {
	path := fmt.Sprintf("members/%s", memberID)
	err = c.Get(path, args, &member)
	if err == nil {
		member.client = c
	}
	return
}

func (o *Organization) GetMembers(args Arguments) (members []*Member, err error) {
	path := fmt.Sprintf("organizations/%s/members", o.ID)
	err = o.client.Get(path, args, &members)
	for i := range members {
		members[i].client = o.client
	}
	return
}

func (b *Board) GetMembers(args Arguments) (members []*Member, err error) {
	path := fmt.Sprintf("boards/%s/members", b.ID)
	err = b.client.Get(path, args, &members)
	for i := range members {
		members[i].client = b.client
	}
	return
}

func (c *Card) GetMembers(args Arguments) (members []*Member, err error) {
	path := fmt.Sprintf("cards/%s/members", c.ID)
	err = c.client.Get(path, args, &members)
	for i := range members {
		members[i].client = c.client
	}
	return
}
