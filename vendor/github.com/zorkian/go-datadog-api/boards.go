/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2019 by authors and contributors.
 */

package datadog

import (
	"fmt"
)

// Board represents a user created dashboard. This is the full dashboard
// struct when we load a dashboard in detail.
type Board struct {
	Title             *string            `json:"title"`
	Widgets           []BoardWidget      `json:"widgets"`
	LayoutType        *string            `json:"layout_type"`
	Id                *string            `json:"id,omitempty"`
	Description       *string            `json:"description,omitempty"`
	TemplateVariables []TemplateVariable `json:"template_variables,omitempty"`
	IsReadOnly        *bool              `json:"is_read_only,omitempty"`
	NotifyList        []string           `json:"notify_list,omitempty"`
	AuthorHandle      *string            `json:"author_handle,omitempty"`
	Url               *string            `json:"url,omitempty"`
	CreatedAt         *string            `json:"created_at,omitempty"`
	ModifiedAt        *string            `json:"modified_at,omitempty"`
}

// GetBoard returns a single dashboard created on this account.
func (client *Client) GetBoard(id string) (*Board, error) {
	var board Board
	if err := client.doJsonRequest("GET", fmt.Sprintf("/v1/dashboard/%s", id), nil, &board); err != nil {
		return nil, err
	}
	return &board, nil
}

// DeleteBoard deletes a dashboard by the identifier.
func (client *Client) DeleteBoard(id string) error {
	return client.doJsonRequest("DELETE", fmt.Sprintf("/v1/dashboard/%s", id), nil, nil)
}

// CreateBoard creates a new dashboard when given a Board struct.
func (client *Client) CreateBoard(board *Board) (*Board, error) {
	var createdBoard Board
	if err := client.doJsonRequest("POST", "/v1/dashboard", board, &createdBoard); err != nil {
		return nil, err
	}
	return &createdBoard, nil
}

// UpdateBoard takes a Board struct and persists it back to the server.
// Use this if you've updated your local and need to push it back.
func (client *Client) UpdateBoard(board *Board) error {
	return client.doJsonRequest("PUT", fmt.Sprintf("/v1/dashboard/%s", *board.Id), board, nil)
}
