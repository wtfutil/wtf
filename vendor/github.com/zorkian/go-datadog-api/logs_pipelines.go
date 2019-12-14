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

const (
	logsPipelinesPath = "/v1/logs/config/pipelines"
)

// LogsPipeline struct to represent the json object received from Logs Public Config API.
type LogsPipeline struct {
	Id         *string              `json:"id,omitempty"`
	Type       *string              `json:"type,omitempty"`
	Name       *string              `json:"name"`
	IsEnabled  *bool                `json:"is_enabled,omitempty"`
	IsReadOnly *bool                `json:"is_read_only,omitempty"`
	Filter     *FilterConfiguration `json:"filter"`
	Processors []LogsProcessor      `json:"processors,omitempty"`
}

// FilterConfiguration struct to represent the json object of filter configuration.
type FilterConfiguration struct {
	Query *string `json:"query"`
}

// GetLogsPipeline queries Logs Public Config API with given a pipeline id for the complete pipeline object.
func (client *Client) GetLogsPipeline(id string) (*LogsPipeline, error) {
	var pipeline LogsPipeline
	if err := client.doJsonRequest("GET", fmt.Sprintf("%s/%s", logsPipelinesPath, id), nil, &pipeline); err != nil {
		return nil, err
	}
	return &pipeline, nil
}

// CreateLogsPipeline sends pipeline creation request to Config API
func (client *Client) CreateLogsPipeline(pipeline *LogsPipeline) (*LogsPipeline, error) {
	var createdPipeline = &LogsPipeline{}
	if err := client.doJsonRequest("POST", logsPipelinesPath, pipeline, createdPipeline); err != nil {
		return nil, err
	}
	return createdPipeline, nil
}

// UpdateLogsPipeline updates the pipeline object of a given pipeline id.
func (client *Client) UpdateLogsPipeline(id string, pipeline *LogsPipeline) (*LogsPipeline, error) {
	var updatedPipeline = &LogsPipeline{}
	if err := client.doJsonRequest("PUT", fmt.Sprintf("%s/%s", logsPipelinesPath, id), pipeline, updatedPipeline); err != nil {
		return nil, err
	}
	return updatedPipeline, nil
}

// DeleteLogsPipeline deletes the pipeline for a given id, returns 200 OK when operation succeed
func (client *Client) DeleteLogsPipeline(id string) error {
	return client.doJsonRequest("DELETE", fmt.Sprintf("%s/%s", logsPipelinesPath, id), nil, nil)
}
