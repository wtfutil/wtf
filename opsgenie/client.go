package opsgenie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type OnCallResponse struct {
	OnCallData []OnCallData `json:"data"`
	Message    string       `json:"message"`
	RequestID  string       `json:"requestId"`
	Took       float32      `json:"took"`
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

func Fetch() *OnCallResponse {
	apiKey := os.Getenv("WTF_OPS_GENIE_API_KEY")

	url := "https://api.opsgenie.com/v2/schedules/on-calls?flat=true"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("GenieKey %s", apiKey))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var onCallResponse OnCallResponse
	if err := json.NewDecoder(resp.Body).Decode(&onCallResponse); err != nil {
		panic(err)
	}

	return &onCallResponse
}
