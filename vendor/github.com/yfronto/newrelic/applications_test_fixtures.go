package newrelic

type getApplicationsTestsInput struct {
	options *ApplicationOptions
	data    string
}

type getApplicationsTestsOutput struct {
	data []Application
	err  error
}

type getApplicationTestsInput struct {
	id   int
	data string
}

type getApplicationTestsOutput struct {
	data *Application
	err  error
}

var (
	testApplicationJSON = `
  {
    "application_summary": {
      "apdex_score": 1.0,
      "apdex_target": 0.5,
      "error_rate": 0.0,
      "host_count": 1,
      "instance_count": 1,
      "response_time": 0.263,
      "throughput": 12.3
    },
    "end_user_summary": {
      "response_time": 0.263,
      "throughput": 12.3,
      "apdex_target": 0.5,
      "apdex_score": 1
    },
    "health_status": "green",
    "id": 12345,
    "language": "java",
    "last_reported_at": "` + testTimeRawString + `",
    "links": {
      "alert_policy": 123,
      "application_hosts": [
        1234567
      ],
      "application_instances": [
        1234568
      ],
      "servers": [
        54321
      ]
    },
    "name": "test.example.com",
    "reporting": true,
    "settings": {
      "app_apdex_threshold": 0.5,
      "enable_real_user_monitoring": true,
      "end_user_apdex_threshold": 1.0,
      "use_server_side_config": false
    }
  }
`
	testApplication = Application{
		ID:             12345,
		Name:           "test.example.com",
		Language:       "java",
		HealthStatus:   "green",
		Reporting:      true,
		LastReportedAt: testTime,
		ApplicationSummary: ApplicationSummary{
			ResponseTime:            0.263,
			Throughput:              12.3,
			ErrorRate:               0,
			ApdexTarget:             0.5,
			ApdexScore:              1,
			HostCount:               1,
			InstanceCount:           1,
			ConcurrentInstanceCount: 0,
		},
		EndUserSummary: EndUserSummary{
			ResponseTime: 0.263,
			Throughput:   12.3,
			ApdexTarget:  0.5,
			ApdexScore:   1,
		},
		Settings: Settings{
			AppApdexThreshold:        0.5,
			EndUserApdexThreshold:    1,
			EnableRealUserMonitoring: true,
			UseServerSideConfig:      false,
		},
		Links: Links{
			Servers:              []int{54321},
			ApplicationHosts:     []int{1234567},
			ApplicationInstances: []int{1234568},
			AlertPolicy:          123,
		},
	}
	testApplications = []Application{
		testApplication,
		testApplication,
	}
	getApplicationTests = []struct {
		in  getApplicationTestsInput
		out getApplicationTestsOutput
	}{
		{
			getApplicationTestsInput{
				id:   12345,
				data: `{ "application":` + testApplicationJSON + `}`,
			},
			getApplicationTestsOutput{
				data: &testApplication,
			},
		},
	}
	getApplicationsTests = []struct {
		in  getApplicationsTestsInput
		out getApplicationsTestsOutput
	}{
		{
			getApplicationsTestsInput{
				options: nil,
				data:    `{"applications":[` + testApplicationJSON + "," + testApplicationJSON + "]}",
			},
			getApplicationsTestsOutput{
				data: testApplications,
				err:  nil,
			},
		},
	}
	applicationOptionsStringerTests = []struct {
		in  *ApplicationOptions
		out string
	}{
		{
			&ApplicationOptions{},
			"",
		},
		{
			&ApplicationOptions{
				Filter: ApplicationFilter{
					Name:     "testName",
					Host:     "testHost",
					IDs:      []int{123, 456},
					Language: "java",
				},
				Page: 5,
			},
			`filter%5Bhost%5D=testHost` +
				`&filter%5Bids%5D=123%2C456` +
				`&filter%5Blanguage%5D=java` +
				`&filter%5Bname%5D=testName` +
				`&page=5`,
		},
		{
			nil,
			"",
		},
	}
)
