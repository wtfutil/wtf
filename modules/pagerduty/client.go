package pagerduty

import (
	"context"
	"time"

	"github.com/PagerDuty/go-pagerduty"
)

const (
	queryTimeFmt = "2006-01-02T15:04:05Z07:00"
)

// GetOnCalls returns a list of people currently on call
func GetOnCalls(apiKey string, scheduleIDs []string) ([]pagerduty.OnCall, error) {
	client := pagerduty.NewClient(apiKey)

	var results []pagerduty.OnCall
	var queryOpts pagerduty.ListOnCallOptions

	queryOpts.ScheduleIDs = scheduleIDs
	queryOpts.Since = time.Now().Format(queryTimeFmt)
	queryOpts.Until = time.Now().Format(queryTimeFmt)

	oncalls, err := client.ListOnCallsWithContext(context.Background(), queryOpts)
	if err != nil {
		return nil, err
	}

	results = append(results, oncalls.OnCalls...)

	for oncalls.APIListObject.More {
		queryOpts.Offset = oncalls.APIListObject.Offset
		oncalls, err = client.ListOnCallsWithContext(context.Background(), queryOpts)
		if err != nil {
			return nil, err
		}
		results = append(results, oncalls.OnCalls...)
	}

	return results, nil
}

// GetIncidents returns a list of unresolved incidents
func GetIncidents(apiKey string, teamIDs []string, userIDs []string) ([]pagerduty.Incident, error) {
	client := pagerduty.NewClient(apiKey)

	var results []pagerduty.Incident

	var queryOpts pagerduty.ListIncidentsOptions
	queryOpts.DateRange = "all"
	queryOpts.Statuses = []string{"triggered", "acknowledged"}
	queryOpts.TeamIDs = teamIDs
	queryOpts.UserIDs = userIDs

	items, err := client.ListIncidentsWithContext(context.Background(), queryOpts)
	if err != nil {
		return nil, err
	}
	results = append(results, items.Incidents...)

	for items.APIListObject.More {
		queryOpts.Offset = items.APIListObject.Offset
		items, err = client.ListIncidentsWithContext(context.Background(), queryOpts)
		if err != nil {
			return nil, err
		}
		results = append(results, items.Incidents...)
	}

	return results, nil
}
