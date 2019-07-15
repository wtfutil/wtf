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
	"strings"
)

// DowntimeType are a classification of a given downtime scope
type DowntimeType int

// The three downtime type classifications.
const (
	StarDowntimeType  DowntimeType = 0
	HostDowntimeType  DowntimeType = 1
	OtherDowntimeType DowntimeType = 2
)

type Recurrence struct {
	Period           *int     `json:"period,omitempty"`
	Type             *string  `json:"type,omitempty"`
	UntilDate        *int     `json:"until_date,omitempty"`
	UntilOccurrences *int     `json:"until_occurrences,omitempty"`
	WeekDays         []string `json:"week_days,omitempty"`
}

type Downtime struct {
	Active      *bool       `json:"active,omitempty"`
	Canceled    *int        `json:"canceled,omitempty"`
	Disabled    *bool       `json:"disabled,omitempty"`
	End         *int        `json:"end,omitempty"`
	Id          *int        `json:"id,omitempty"`
	Message     *string     `json:"message,omitempty"`
	MonitorId   *int        `json:"monitor_id,omitempty"`
	MonitorTags []string    `json:"monitor_tags,omitempty"`
	ParentId    *int        `json:"parent_id,omitempty"`
	Timezone    *string     `json:"timezone,omitempty"`
	Recurrence  *Recurrence `json:"recurrence,omitempty"`
	Scope       []string    `json:"scope,omitempty"`
	Start       *int        `json:"start,omitempty"`
	CreatorID   *int        `json:"creator_id,omitempty"`
	UpdaterID   *int        `json:"updater_id,omitempty"`
	Type        *int        `json:"downtime_type,omitempty"`
}

// DowntimeType returns the canonical downtime type classification.
// This is calculated based on the provided server response, but the logic is copied down here to calculate locally.
func (d *Downtime) DowntimeType() DowntimeType {
	if d.Type != nil {
		switch *d.Type {
		case 0:
			return StarDowntimeType
		case 1:
			return HostDowntimeType
		default:
			return OtherDowntimeType
		}
	}
	if len(d.Scope) == 1 {
		if d.Scope[0] == "*" {
			return StarDowntimeType
		}
		if strings.HasPrefix(d.Scope[0], "host:") {
			return HostDowntimeType
		}
	}
	return OtherDowntimeType
}

// reqDowntimes retrieves a slice of all Downtimes.
type reqDowntimes struct {
	Downtimes []Downtime `json:"downtimes,omitempty"`
}

// CreateDowntime adds a new downtme to the system. This returns a pointer
// to a Downtime so you can pass that to UpdateDowntime or CancelDowntime
// later if needed.
func (client *Client) CreateDowntime(downtime *Downtime) (*Downtime, error) {
	var out Downtime
	if err := client.doJsonRequest("POST", "/v1/downtime", downtime, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateDowntime takes a downtime that was previously retrieved through some method
// and sends it back to the server.
func (client *Client) UpdateDowntime(downtime *Downtime) error {
	return client.doJsonRequest("PUT", fmt.Sprintf("/v1/downtime/%d", *downtime.Id),
		downtime, downtime)
}

// Getdowntime retrieves an downtime by identifier.
func (client *Client) GetDowntime(id int) (*Downtime, error) {
	var out Downtime
	if err := client.doJsonRequest("GET", fmt.Sprintf("/v1/downtime/%d", id), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteDowntime removes an downtime from the system.
func (client *Client) DeleteDowntime(id int) error {
	return client.doJsonRequest("DELETE", fmt.Sprintf("/v1/downtime/%d", id),
		nil, nil)
}

// GetDowntimes returns a slice of all downtimes.
func (client *Client) GetDowntimes() ([]Downtime, error) {
	var out reqDowntimes
	if err := client.doJsonRequest("GET", "/v1/downtime", nil, &out.Downtimes); err != nil {
		return nil, err
	}
	return out.Downtimes, nil
}
