/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

import (
	"fmt"
)

// Screenboard represents a user created screenboard. This is the full screenboard
// struct when we load a screenboard in detail.
type Screenboard struct {
	Id                *int               `json:"id,omitempty"`
	NewId             *string            `json:"new_id,omitempty"`
	Title             *string            `json:"board_title,omitempty"`
	Height            *int               `json:"height,omitempty"`
	Width             *int               `json:"width,omitempty"`
	Shared            *bool              `json:"shared,omitempty"`
	TemplateVariables []TemplateVariable `json:"template_variables,omitempty"`
	Widgets           []Widget           `json:"widgets"`
	ReadOnly          *bool              `json:"read_only,omitempty"`
}

// ScreenboardLite represents a user created screenboard. This is the mini
// struct when we load the summaries.
type ScreenboardLite struct {
	Id       *int    `json:"id,omitempty"`
	Resource *string `json:"resource,omitempty"`
	Title    *string `json:"title,omitempty"`
}

// reqGetScreenboards from /api/v1/screen
type reqGetScreenboards struct {
	Screenboards []*ScreenboardLite `json:"screenboards,omitempty"`
}

// GetScreenboard returns a single screenboard created on this account.
func (client *Client) GetScreenboard(id interface{}) (*Screenboard, error) {
	stringId, err := GetStringId(id)
	if err != nil {
		return nil, err
	}

	out := &Screenboard{}
	if err := client.doJsonRequest("GET", fmt.Sprintf("/v1/screen/%s", stringId), nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetScreenboards returns a list of all screenboards created on this account.
func (client *Client) GetScreenboards() ([]*ScreenboardLite, error) {
	var out reqGetScreenboards
	if err := client.doJsonRequest("GET", "/v1/screen", nil, &out); err != nil {
		return nil, err
	}
	return out.Screenboards, nil
}

// DeleteScreenboard deletes a screenboard by the identifier.
func (client *Client) DeleteScreenboard(id int) error {
	return client.doJsonRequest("DELETE", fmt.Sprintf("/v1/screen/%d", id), nil, nil)
}

// CreateScreenboard creates a new screenboard when given a Screenboard struct. Note
// that the Id, Resource, Url and similar elements are not used in creation.
func (client *Client) CreateScreenboard(board *Screenboard) (*Screenboard, error) {
	out := &Screenboard{}
	if err := client.doJsonRequest("POST", "/v1/screen", board, out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateScreenboard in essence takes a Screenboard struct and persists it back to
// the server. Use this if you've updated your local and need to push it back.
func (client *Client) UpdateScreenboard(board *Screenboard) error {
	return client.doJsonRequest("PUT", fmt.Sprintf("/v1/screen/%d", *board.Id), board, nil)
}

type ScreenShareResponse struct {
	BoardId   int    `json:"board_id"`
	PublicUrl string `json:"public_url"`
}

// ShareScreenboard shares an existing screenboard, it takes and updates ScreenShareResponse
func (client *Client) ShareScreenboard(id int, response *ScreenShareResponse) error {
	return client.doJsonRequest("POST", fmt.Sprintf("/v1/screen/share/%d", id), nil, response)
}

// RevokeScreenboard revokes a currently shared screenboard
func (client *Client) RevokeScreenboard(id int) error {
	return client.doJsonRequest("DELETE", fmt.Sprintf("/v1/screen/share/%d", id), nil, nil)
}
