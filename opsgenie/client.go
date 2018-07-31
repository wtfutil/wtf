package opsgenie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/senorprogrammer/wtf/wtf"
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

/* -------------------- Exported Functions -------------------- */

func Fetch() (*OnCallResponse, error) {
	scheduleUrl := "https://api.opsgenie.com/v2/schedules/on-calls?flat=true"

	response, err := opsGenieRequest(scheduleUrl, apiKey())

	return response, err
}

/* -------------------- Unexported Functions -------------------- */

func apiKey() string {
	return wtf.Config.UString(
		"wtf.mods.opsgenie.apiKey",
		os.Getenv("WTF_OPS_GENIE_API_KEY"),
	)
}

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
