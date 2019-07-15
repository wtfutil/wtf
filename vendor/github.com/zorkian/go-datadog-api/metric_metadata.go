/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

import "fmt"

// MetricMetadata allows you to edit fields of a metric's metadata.
type MetricMetadata struct {
	Type           *string `json:"type,omitempty"`
	Description    *string `json:"description,omitempty"`
	ShortName      *string `json:"short_name,omitempty"`
	Unit           *string `json:"unit,omitempty"`
	PerUnit        *string `json:"per_unit,omitempty"`
	StatsdInterval *int    `json:"statsd_interval,omitempty"`
}

// ViewMetricMetadata allows you to get metadata about a specific metric.
func (client *Client) ViewMetricMetadata(mn string) (*MetricMetadata, error) {
	var out MetricMetadata
	if err := client.doJsonRequest("GET", fmt.Sprintf("/v1/metrics/%s", mn), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// EditMetricMetadata edits the metadata for the given metric.
func (client *Client) EditMetricMetadata(mn string, mm *MetricMetadata) (*MetricMetadata, error) {
	var out MetricMetadata
	if err := client.doJsonRequest("PUT", fmt.Sprintf("/v1/metrics/%s", mn), mm, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
