package opsgenie

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OnCallResponse struct {
	OnCallData OnCallData `json:"data"`
	Message    string     `json:"message"`
	RequestID  string     `json:"requestId"`
	Took       float32    `json:"took"`
}

type OnCallData struct {
	Recipients []string `json:"onCallRecipients"`
	Parent     Parent   `json:"_parent"`
}

type Parent struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

var opsGenieAPIUrl = map[string]string{
	"us": "https://api.opsgenie.com",
	"eu": "https://api.eu.opsgenie.com",
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Fetch(scheduleIdentifierType string, schedules []string) ([]*OnCallResponse, error) {
	agregatedResponses := []*OnCallResponse{}

	if regionURL, regionErr := opsGenieAPIUrl[widget.settings.region]; regionErr {
		for _, sched := range schedules {
			scheduleURL := fmt.Sprintf("%s/v2/schedules/%s/on-calls?scheduleIdentifierType=%s&flat=true", regionURL, sched, scheduleIdentifierType)
			response, err := opsGenieRequest(scheduleURL, widget.settings.apiKey)
			agregatedResponses = append(agregatedResponses, response)
			if err != nil {
				return nil, err
			}
		}
		return agregatedResponses, nil
	} else {
		return nil, fmt.Errorf("you specified wrong region. Possible options are only 'us' and 'eu'")
	}
}

/* -------------------- Unexported Functions -------------------- */

func opsGenieRequest(url string, apiKey string) (*OnCallResponse, error) {
	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("GenieKey %s", apiKey))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	response := &OnCallResponse{}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}
