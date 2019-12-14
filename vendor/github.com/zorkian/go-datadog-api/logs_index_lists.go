package datadog

const logsIndexListPath = "/v1/logs/config/index-order"

// LogsIndexList represents the index list object from config API.
type LogsIndexList struct {
	IndexNames []string `json:"index_names"`
}

// GetLogsIndexList gets the full list of available indexes by their names.
func (client *Client) GetLogsIndexList() (*LogsIndexList, error) {
	var indexList LogsIndexList
	if err := client.doJsonRequest("GET", logsIndexListPath, nil, &indexList); err != nil {
		return nil, err
	}
	return &indexList, nil
}

// UpdateLogsIndexList updates the order of indexes.
func (client *Client) UpdateLogsIndexList(indexList *LogsIndexList) (*LogsIndexList, error) {
	var updatedIndexList = &LogsIndexList{}
	if err := client.doJsonRequest("PUT", logsIndexListPath, indexList, updatedIndexList); err != nil {
		return nil, err
	}
	return updatedIndexList, nil
}
