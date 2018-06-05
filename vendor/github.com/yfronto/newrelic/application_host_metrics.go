package newrelic

import (
	"fmt"
)

// GetApplicationHostMetrics will return a slice of Metric items for a
// particular Application ID's Host ID, optionally filtering by
// MetricsOptions.
func (c *Client) GetApplicationHostMetrics(appID, hostID int, options *MetricsOptions) ([]Metric, error) {
	mc := NewMetricClient(c)

	return mc.GetMetrics(
		fmt.Sprintf(
			"applications/%d/hosts/%d/metrics.json",
			appID,
			hostID,
		),
		options,
	)
}

// GetApplicationHostMetricData will return all metric data for a particular
// application's host and slice of metric names, optionally filtered by
// MetricDataOptions.
func (c *Client) GetApplicationHostMetricData(appID, hostID int, names []string, options *MetricDataOptions) (*MetricDataResponse, error) {
	mc := NewMetricClient(c)

	return mc.GetMetricData(
		fmt.Sprintf(
			"applications/%d/hosts/%d/metrics/data.json",
			appID,
			hostID,
		),
		names,
		options,
	)
}
