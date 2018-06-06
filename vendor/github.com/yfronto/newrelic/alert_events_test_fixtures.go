package newrelic

type getAlertEventsTestsInput struct {
	options *AlertEventOptions
	data    string
}

type getAlertEventsTestsOutput struct {
	data []AlertEvent
	err  error
}

const (
	testAlertEventJSON = `
  {
    "id": 123,
    "event_type": "VIOLATION_OPEN",
    "product": "APM",
    "entity_type": "Application",
    "entity_group_id": 1234,
    "entity_id": 12,
    "priority": "Warning",
    "description": "Test Alert",
    "timestamp": 1472355451353,
    "incident_id": 23
  }
`
)

var (
	testAlertEvent = AlertEvent{
		ID:            123,
		EventType:     "VIOLATION_OPEN",
		Product:       "APM",
		EntityType:    "Application",
		EntityGroupID: 1234,
		EntityID:      12,
		Priority:      "Warning",
		Description:   "Test Alert",
		Timestamp:     1472355451353,
		IncidentID:    23,
	}
	getAlertEventsTests = []struct {
		in  getAlertEventsTestsInput
		out getAlertEventsTestsOutput
	}{
		{
			getAlertEventsTestsInput{
				options: nil,
				data:    `{"recent_events": [` + testAlertEventJSON + `]}`,
			},
			getAlertEventsTestsOutput{
				data: []AlertEvent{
					testAlertEvent,
				},
				err: nil,
			},
		},
	}
	alertEventOptionsStringerTests = []struct {
		in  *AlertEventOptions
		out string
	}{
		{
			&AlertEventOptions{},
			"",
		},
		{
			nil,
			"",
		},
		{
			&AlertEventOptions{
				Filter: AlertEventFilter{
					Product:       "testProduct",
					EntityType:    "testEntityType",
					EntityGroupID: 123,
					EntityID:      1234,
					EventType:     "testEventType",
				},
				Page: 1,
			},
			"filter%5Bentity_group_id%5D=123" +
				"&filter%5Bentity_id%5D=1234" +
				"&filter%5Bentity_type%5D=testEntityType" +
				"&filter%5Bevent_type%5D=testEventType" +
				"&filter%5Bproduct%5D=testProduct" +
				"&page=1",
		},
	}
)
