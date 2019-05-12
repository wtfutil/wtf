package pagerduty

import (
	"time"

	"github.com/PagerDuty/go-pagerduty"
)

// GetOnCalls returns a list of people currently on call
func GetOnCalls(apiKey string, scheduleIDs []string) ([]pagerduty.OnCall, error) {
	client := pagerduty.NewClient(apiKey)

	var results []pagerduty.OnCall
	var queryOpts pagerduty.ListOnCallOptions

	queryOpts.ScheduleIDs = scheduleIDs

	timeFmt := "2006-01-02T15:04:05Z07:00"
	queryOpts.Since = time.Now().Format(timeFmt)
	queryOpts.Until = time.Now().Format(timeFmt)

	oncalls, err := client.ListOnCalls(queryOpts)
	if err != nil {
		return nil, err
	}

	results = append(results, oncalls.OnCalls...)

	for oncalls.APIListObject.More == true {
		queryOpts.APIListObject.Offset = oncalls.APIListObject.Offset
		oncalls, err = client.ListOnCalls(queryOpts)
		if err != nil {
			return nil, err
		}
		results = append(results, oncalls.OnCalls...)
	}

	return results, nil
}

// GetIncidents returns a list of people currently on call
func GetIncidents(apiKey string) ([]pagerduty.Incident, error) {
	client := pagerduty.NewClient(apiKey)

	var results []pagerduty.Incident

	var queryOpts pagerduty.ListIncidentsOptions
	queryOpts.DateRange = "all"
	queryOpts.Statuses = []string{"triggered", "acknowledged"}

	items, err := client.ListIncidents(queryOpts)
	if err != nil {
		return nil, err
	}
	results = append(results, items.Incidents...)

	for items.APIListObject.More == true {
		queryOpts.APIListObject.Offset = items.APIListObject.Offset
		items, err = client.ListIncidents(queryOpts)
		if err != nil {
			return nil, err
		}
		results = append(results, items.Incidents...)
	}

	return results, nil
}
