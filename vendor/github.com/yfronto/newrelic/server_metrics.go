package newrelic

import (
	"fmt"
)

// GetServerMetrics will return a slice of Metric items for a particular
// Server ID, optionally filtering by MetricsOptions.
func (c *Client) GetServerMetrics(id int, options *MetricsOptions) ([]Metric, error) {
	mc := NewMetricClient(c)

	return mc.GetMetrics(
		fmt.Sprintf(
			"servers/%d/metrics.json",
			id,
		),
		options,
	)
}

// GetServerMetricData will return all metric data for a particular Server and
// slice of metric names, optionally filtered by MetricDataOptions.
func (c *Client) GetServerMetricData(id int, names []string, options *MetricDataOptions) (*MetricDataResponse, error) {
	mc := NewMetricClient(c)

	return mc.GetMetricData(
		fmt.Sprintf(
			"servers/%d/metrics/data.json",
			id,
		),
		names,
		options,
	)
}
