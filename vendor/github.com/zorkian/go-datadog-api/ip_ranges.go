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

// IP ranges US: https://ip-ranges.datadoghq.com
// EU: https://ip-ranges.datadoghq.eu
// Same structure
type IPRangesResp struct {
	Agents     map[string][]string `json:"agents"`
	API        map[string][]string `json:"api"`
	Apm        map[string][]string `json:"apm"`
	Logs       map[string][]string `json:"logs"`
	Process    map[string][]string `json:"process"`
	Synthetics map[string][]string `json:"synthetics"`
	Webhooks   map[string][]string `json:"webhooks"`
}

// GetIPRanges returns all IP addresses by section: agents, api, apm, logs, process, synthetics, webhooks
func (client *Client) GetIPRanges() (*IPRangesResp, error) {
	var out IPRangesResp
	urlIPRanges, err := client.URLIPRanges()
	if err != nil {
		return nil, fmt.Errorf("Error getting IP Ranges URL: %s", err)
	}
	if err := client.doJsonRequest("GET", urlIPRanges, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
