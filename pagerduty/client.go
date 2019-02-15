package pagerduty

import (
	"time"

	"github.com/PagerDuty/go-pagerduty"
)

type PagerDutyClient struct {
	Client *pagerduty.Client
}

func NewPDClient(apiKey string) *PagerDutyClient {
	return &PagerDutyClient{
		Client: pagerduty.NewClient(apiKey),
	}
}

// GetOnCalls returns a list of people currently on call
func (pd *PagerDutyClient) GetOnCalls() ([]pagerduty.OnCall, error) {
	client := pd.Client

	var results []pagerduty.OnCall

	var queryOpts pagerduty.ListOnCallOptions
	queryOpts.Since = time.Now().Format("2006-01-02T15:04:05Z07:00")
	queryOpts.Until = time.Now().Format("2006-01-02T15:04:05Z07:00")

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
func (pd *PagerDutyClient) GetIncidents() ([]pagerduty.Incident, error) {
	client := pd.Client

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
