package newrelic

import (
	"fmt"
)

// GetComponentMetrics will return a slice of Metric items for a
// particular Component ID, optionally filtered by MetricsOptions.
func (c *Client) GetComponentMetrics(id int, options *MetricsOptions) ([]Metric, error) {
	mc := NewMetricClient(c)

	return mc.GetMetrics(
		fmt.Sprintf(
			"components/%d/metrics.json",
			id,
		),
		options,
	)
}

// GetComponentMetricData will return all metric data for a particular
// component, optionally filtered by MetricDataOptions.
func (c *Client) GetComponentMetricData(id int, names []string, options *MetricDataOptions) (*MetricDataResponse, error) {
	mc := NewMetricClient(c)

	return mc.GetMetricData(
		fmt.Sprintf(
			"components/%d/metrics/data.json",
			id,
		),
		names,
		options,
	)
}
