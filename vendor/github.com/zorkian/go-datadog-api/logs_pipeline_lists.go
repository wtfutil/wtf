package datadog

const (
	logsPipelineListPath = "/v1/logs/config/pipeline-order"
)

// LogsPipelineList struct represents the pipeline order from Logs Public Config API.
type LogsPipelineList struct {
	PipelineIds []string `json:"pipeline_ids"`
}

// GetLogsPipelineList get the full list of created pipelines.
func (client *Client) GetLogsPipelineList() (*LogsPipelineList, error) {
	var pipelineList LogsPipelineList
	if err := client.doJsonRequest("GET", logsPipelineListPath, nil, &pipelineList); err != nil {
		return nil, err
	}
	return &pipelineList, nil
}

// UpdateLogsPipelineList updates the pipeline list order, it returns error (422 Unprocessable Entity)
// if one tries to delete or add pipeline.
func (client *Client) UpdateLogsPipelineList(pipelineList *LogsPipelineList) (*LogsPipelineList, error) {
	var updatedPipelineList = &LogsPipelineList{}
	if err := client.doJsonRequest("PUT", logsPipelineListPath, pipelineList, updatedPipelineList); err != nil {
		return nil, err
	}
	return updatedPipelineList, nil
}
