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
	"default": "https://api.opsgenie.com",
	"europe":  "https://api.eu.opsgenie.com",
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Fetch(scheduleIdentifierType string, schedules []string) ([]*OnCallResponse, error) {
	agregatedResponses := []*OnCallResponse{}
	region := "default"

	for _, sched := range schedules {
		if widget.settings.isEurope {
			region = "europe"
		}
		scheduleUrl := fmt.Sprintf("%s/v2/schedules/%s/on-calls?scheduleIdentifierType=%s&flat=true", opsGenieAPIUrl[region], sched, scheduleIdentifierType)
		response, err := opsGenieRequest(scheduleUrl, widget.settings.apiKey)
		agregatedResponses = append(agregatedResponses, response)
		if err != nil {
			return nil, err
		}
	}
	return agregatedResponses, nil
}

/* -------------------- Unexported Functions -------------------- */

func opsGenieRequest(url string, apiKey string) (*OnCallResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("GenieKey %s", apiKey))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &OnCallResponse{}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}
