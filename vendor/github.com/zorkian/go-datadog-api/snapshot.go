/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2016 by authors and contributors.
 */

package datadog

import (
	"fmt"
	"net/url"
	"time"
)

func (client *Client) doSnapshotRequest(values url.Values) (string, error) {
	out := struct {
		SnapshotURL string `json:"snapshot_url,omitempty"`
	}{}
	if err := client.doJsonRequest("GET", "/v1/graph/snapshot?"+values.Encode(), nil, &out); err != nil {
		return "", err
	}
	return out.SnapshotURL, nil
}

// Snapshot creates an image from a graph and returns the URL of the image.
func (client *Client) Snapshot(query string, start, end time.Time, eventQuery string) (string, error) {
	options := map[string]string{"metric_query": query, "event_query": eventQuery}

	return client.SnapshotGeneric(options, start, end)
}

// Generic function for snapshots, use map[string]string to create url.Values() instead of pre-defined params
func (client *Client) SnapshotGeneric(options map[string]string, start, end time.Time) (string, error) {
	v := url.Values{}
	v.Add("start", fmt.Sprintf("%d", start.Unix()))
	v.Add("end", fmt.Sprintf("%d", end.Unix()))

	for opt, val := range options {
		v.Add(opt, val)
	}

	return client.doSnapshotRequest(v)
}
