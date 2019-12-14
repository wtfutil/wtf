package datadog

import (
	"fmt"
)

const logsIndexPath = "/v1/logs/config/indexes"

// LogsIndex represents the Logs index object from config API.
type LogsIndex struct {
	Name             *string              `json:"name"`
	NumRetentionDays *int64               `json:"num_retention_days,omitempty"`
	DailyLimit       *int64               `json:"daily_limit,omitempty"`
	IsRateLimited    *bool                `json:"is_rate_limited,omitempty"`
	Filter           *FilterConfiguration `json:"filter"`
	ExclusionFilters []ExclusionFilter    `json:"exclusion_filters"`
}

// ExclusionFilter represents the index exclusion filter object from config API.
type ExclusionFilter struct {
	Name      *string `json:"name"`
	IsEnabled *bool   `json:"is_enabled,omitempty"`
	Filter    *Filter `json:"filter"`
}

// Filter represents the index filter object from config API.
type Filter struct {
	Query      *string  `json:"query,omitempty"`
	SampleRate *float64 `json:"sample_rate,omitempty"`
}

// GetLogsIndex gets the specific logs index by specific name.
func (client *Client) GetLogsIndex(name string) (*LogsIndex, error) {
	var index LogsIndex
	if err := client.doJsonRequest("GET", fmt.Sprintf("%s/%s", logsIndexPath, name), nil, &index); err != nil {
		return nil, err
	}
	return &index, nil
}

// UpdateLogsIndex updates the specific index by it's name.
func (client *Client) UpdateLogsIndex(name string, index *LogsIndex) (*LogsIndex, error) {
	var updatedIndex = &LogsIndex{}
	if err := client.doJsonRequest("PUT", fmt.Sprintf("%s/%s", logsIndexPath, name), index, updatedIndex); err != nil {
		return nil, err
	}
	return updatedIndex, nil
}
