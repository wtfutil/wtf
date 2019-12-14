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

// DashboardListItemV2 represents a single dashboard in a dashboard list.
type DashboardListItemV2 struct {
	ID   *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
}

type reqDashboardListItemsV2 struct {
	Dashboards []DashboardListItemV2 `json:"dashboards,omitempty"`
}

type reqAddedDashboardListItemsV2 struct {
	Dashboards []DashboardListItemV2 `json:"added_dashboards_to_list,omitempty"`
}

type reqDeletedDashboardListItemsV2 struct {
	Dashboards []DashboardListItemV2 `json:"deleted_dashboards_from_list,omitempty"`
}

// GetDashboardListItemsV2 fetches the dashboard list's dashboard definitions.
func (client *Client) GetDashboardListItemsV2(id int) ([]DashboardListItemV2, error) {
	var out reqDashboardListItemsV2
	if err := client.doJsonRequest("GET", fmt.Sprintf("/v2/dashboard/lists/manual/%d/dashboards", id), nil, &out); err != nil {
		return nil, err
	}
	return out.Dashboards, nil
}

// AddDashboardListItemsV2 adds dashboards to an existing dashboard list.
//
// Any items already in the list are ignored (not added twice).
func (client *Client) AddDashboardListItemsV2(dashboardListID int, items []DashboardListItemV2) ([]DashboardListItemV2, error) {
	req := reqDashboardListItemsV2{items}
	var out reqAddedDashboardListItemsV2
	if err := client.doJsonRequest("POST", fmt.Sprintf("/v2/dashboard/lists/manual/%d/dashboards", dashboardListID), req, &out); err != nil {
		return nil, err
	}
	return out.Dashboards, nil
}

// UpdateDashboardListItemsV2 updates dashboards of an existing dashboard list.
//
// This will set the list of dashboards to contain only the items in items.
func (client *Client) UpdateDashboardListItemsV2(dashboardListID int, items []DashboardListItemV2) ([]DashboardListItemV2, error) {
	req := reqDashboardListItemsV2{items}
	var out reqDashboardListItemsV2
	if err := client.doJsonRequest("PUT", fmt.Sprintf("/v2/dashboard/lists/manual/%d/dashboards", dashboardListID), req, &out); err != nil {
		return nil, err
	}
	return out.Dashboards, nil
}

// DeleteDashboardListItemsV2 deletes dashboards from an existing dashboard list.
//
// Deletes any dashboards in the list of items from the dashboard list.
func (client *Client) DeleteDashboardListItemsV2(dashboardListID int, items []DashboardListItemV2) ([]DashboardListItemV2, error) {
	req := reqDashboardListItemsV2{items}
	var out reqDeletedDashboardListItemsV2
	if err := client.doJsonRequest("DELETE", fmt.Sprintf("/v2/dashboard/lists/manual/%d/dashboards", dashboardListID), req, &out); err != nil {
		return nil, err
	}
	return out.Dashboards, nil
}
