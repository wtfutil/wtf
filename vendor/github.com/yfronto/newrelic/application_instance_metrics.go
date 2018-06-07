package newrelic

import (
	"fmt"
)

// GetApplicationInstanceMetrics will return a slice of Metric items for a
// particular Application ID's instance ID, optionally filtering by
// MetricsOptions.
func (c *Client) GetApplicationInstanceMetrics(appID, instanceID int, options *MetricsOptions) ([]Metric, error) {
	mc := NewMetricClient(c)

	return mc.GetMetrics(
		fmt.Sprintf(
			"applications/%d/instances/%d/metrics.json",
			appID,
			instanceID,
		),
		options,
	)
}

// GetApplicationInstanceMetricData will return all metric data for a
// particular application's instance and slice of metric names, optionally
// filtered by MetricDataOptions.
func (c *Client) GetApplicationInstanceMetricData(appID, instanceID int, names []string, options *MetricDataOptions) (*MetricDataResponse, error) {
	mc := NewMetricClient(c)

	return mc.GetMetricData(
		fmt.Sprintf(
			"applications/%d/instances/%d/metrics/data.json",
			appID,
			instanceID,
		),
		names,
		options,
	)
}
