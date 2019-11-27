package newrelic

type getApplicationMetricsTestsInput struct {
	id      int
	options *MetricsOptions
	data    string
}

type getApplicationMetricsTestsOutput struct {
	data []Metric
	err  error
}

type getApplicationMetricDataTestsInput struct {
	id      int
	Names   []string
	options *MetricDataOptions
	data    string
}

type getApplicationMetricDataTestsOutput struct {
	data *MetricDataResponse
	err  error
}

const (
	testApplicationMetricJSON = `
  {
    "name": "testMetric",
    "values": [
      "testValue1",
      "testValue2"
    ]
  }
`
)

var (
	testApplicationMetric = Metric{
		Name:   "testMetric",
		Values: []string{"testValue1", "testValue2"},
	}
	getApplicaitonMetricsTests = []struct {
		in  getApplicationMetricsTestsInput
		out getApplicationMetricsTestsOutput
	}{
		{
			getApplicationMetricsTestsInput{
				id:      123,
				options: nil,
				data: `{"metrics": [` +
					testApplicationMetricJSON + `,` +
					testApplicationMetricJSON +
					`]}`,
			},
			getApplicationMetricsTestsOutput{
				data: []Metric{
					testApplicationMetric,
					testApplicationMetric,
				},
				err: nil,
			},
		},
	}
	applicationMetricOptionsStringerTests = []struct {
		in  *MetricsOptions
		out string
	}{
		{
			&MetricsOptions{},
			"",
		},
		{
			nil,
			"",
		},
		{
			&MetricsOptions{
				Name: "testName",
				Page: 5,
			},
			"name=testName" +
				"&page=5",
		},
	}
	applicationMetricDataOptionsStringerTests = []struct {
		in  *MetricDataOptions
		out string
	}{
		{
			&MetricDataOptions{},
			"",
		},
		{
			nil,
			"",
		},
		{
			&MetricDataOptions{
				Names:     Array{[]string{"test1", "test2"}},
				Values:    Array{[]string{"value1", "value2"}},
				From:      testTime,
				To:        testTime,
				Period:    123,
				Summarize: true,
				Raw:       true,
			},
			"from=" + testTimeStringEscaped +
				"&names%5B%5D=test1" +
				"&names%5B%5D=test2" +
				"&period=123" +
				"&raw=true" +
				"&summarize=true" +
				"&to=" + testTimeStringEscaped +
				"&values%5B%5D=value1&values%5B%5D=value2",
		},
	}
	testApplicationMetricDataJSON = `
  {
    "metric_data": {
      "from": "` + testTimeRawString + `",
      "to": "` + testTimeRawString + `",
      "metrics_found": ["name1"],
      "metrics_not_found": ["name2"],
      "metrics": [
        {
          "name": "testName",
          "timeslices": [
            {
              "from": "` + testTimeRawString + `",
              "to": "` + testTimeRawString + `",
              "values": {"testVal": 1.234}
            }
          ]
        }
      ]
    }
  }
`
	testApplicationMetricData = MetricData{
		Name: "testName",
		Timeslices: []MetricTimeslice{
			MetricTimeslice{
				From: testTime,
				To:   testTime,
				Values: map[string]float64{
					"testVal": 1.234,
				},
			},
		},
	}
	getApplicaitonMetricDataTests = []struct {
		in  getApplicationMetricDataTestsInput
		out getApplicationMetricDataTestsOutput
	}{
		{
			getApplicationMetricDataTestsInput{
				id:      1234,
				Names:   []string{"name1", "name2"},
				options: &MetricDataOptions{},
				data:    testApplicationMetricDataJSON,
			},
			getApplicationMetricDataTestsOutput{
				data: &MetricDataResponse{
					From:            testTime,
					To:              testTime,
					MetricsFound:    []string{"name1"},
					MetricsNotFound: []string{"name2"},
					Metrics:         []MetricData{testApplicationMetricData},
				},
				err: nil,
			},
		},
	}
)
