/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2018 by authors and contributors.
 */

package datadog

import (
	"fmt"
)

const (
	DashboardListItemCustomTimeboard        = "custom_timeboard"
	DashboardListItemCustomScreenboard      = "custom_screenboard"
	DashboardListItemIntegerationTimeboard  = "integration_timeboard"
	DashboardListItemIntegrationScreenboard = "integration_screenboard"
	DashboardListItemHostTimeboard          = "host_timeboard"
)

// DashboardList represents a dashboard list.
type DashboardList struct {
	Id             *int    `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	DashboardCount *int    `json:"dashboard_count,omitempty"`
}

// DashboardListItem represents a single dashboard in a dashboard list.
type DashboardListItem struct {
	Id   *int    `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
}

type reqDashboardListItems struct {
	Dashboards []DashboardListItem `json:"dashboards,omitempty"`
}

type reqAddedDashboardListItems struct {
	Dashboards []DashboardListItem `json:"added_dashboards_to_list,omitempty"`
}

type reqDeletedDashboardListItems struct {
	Dashboards []DashboardListItem `json:"deleted_dashboards_from_list,omitempty"`
}

type reqUpdateDashboardList struct {
	Name string `json:"name,omitempty"`
}

type reqGetDashboardLists struct {
	DashboardLists []DashboardList `json:"dashboard_lists,omitempty"`
}

// GetDashboardList returns a single dashboard list created on this account.
func (client *Client) GetDashboardList(id int) (*DashboardList, error) {
	var out DashboardList
	if err := client.doJsonRequest("GET", fmt.Sprintf("/v1/dashboard/lists/manual/%d", id), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetDashboardLists returns a list of all dashboard lists created on this account.
func (client *Client) GetDashboardLists() ([]DashboardList, error) {
	var out reqGetDashboardLists
	if err := client.doJsonRequest("GET", "/v1/dashboard/lists/manual", nil, &out); err != nil {
		return nil, err
	}
	return out.DashboardLists, nil
}

// CreateDashboardList returns a single dashboard list created on this account.
func (client *Client) CreateDashboardList(list *DashboardList) (*DashboardList, error) {
	var out DashboardList
	if err := client.doJsonRequest("POST", "/v1/dashboard/lists/manual", list, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateDashboardList returns a single dashboard list created on this account.
func (client *Client) UpdateDashboardList(list *DashboardList) error {
	req := reqUpdateDashboardList{list.GetName()}
	return client.doJsonRequest("PUT", fmt.Sprintf("/v1/dashboard/lists/manual/%d", *list.Id), req, nil)
}

// DeleteDashboardList deletes a dashboard list by the identifier.
func (client *Client) DeleteDashboardList(id int) error {
	return client.doJsonRequest("DELETE", fmt.Sprintf("/v1/dashboard/lists/manual/%d", id), nil, nil)
}

// GetDashboardListItems fetches the dashboard list's dashboard definitions.
func (client *Client) GetDashboardListItems(id int) ([]DashboardListItem, error) {
	var out reqDashboardListItems
	if err := client.doJsonRequest("GET", fmt.Sprintf("/v1/dashboard/lists/manual/%d/dashboards", id), nil, &out); err != nil {
		return nil, err
	}
	return out.Dashboards, nil
}

// AddDashboardListItems adds dashboards to an existing dashboard list.
//
// Any items already in the list are ignored (not added twice).
func (client *Client) AddDashboardListItems(dashboardListId int, items []DashboardListItem) ([]DashboardListItem, error) {
	req := reqDashboardListItems{items}
	var out reqAddedDashboardListItems
	if err := client.doJsonRequest("POST", fmt.Sprintf("/v1/dashboard/lists/manual/%d/dashboards", dashboardListId), req, &out); err != nil {
		return nil, err
	}
	return out.Dashboards, nil
}

// UpdateDashboardListItems updates dashboards of an existing dashboard list.
//
// This will set the list of dashboards to contain only the items in items.
func (client *Client) UpdateDashboardListItems(dashboardListId int, items []DashboardListItem) ([]DashboardListItem, error) {
	req := reqDashboardListItems{items}
	var out reqDashboardListItems
	if err := client.doJsonRequest("PUT", fmt.Sprintf("/v1/dashboard/lists/manual/%d/dashboards", dashboardListId), req, &out); err != nil {
		return nil, err
	}
	return out.Dashboards, nil
}

// DeleteDashboardListItems deletes dashboards from an existing dashboard list.
//
// Deletes any dashboards in the list of items from the dashboard list.
func (client *Client) DeleteDashboardListItems(dashboardListId int, items []DashboardListItem) ([]DashboardListItem, error) {
	req := reqDashboardListItems{items}
	var out reqDeletedDashboardListItems
	if err := client.doJsonRequest("DELETE", fmt.Sprintf("/v1/dashboard/lists/manual/%d/dashboards", dashboardListId), req, &out); err != nil {
		return nil, err
	}
	return out.Dashboards, nil
}
