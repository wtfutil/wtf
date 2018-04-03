package opsgenie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Data struct {
	OnCallRecipients []string `json:"onCallRecipients"`
	Parent           Parent   `json:"_parent"`
}
type OnCallData struct {
	Data      Data    `json:"data"`
	Message   string  `json:"message"`
	RequestID string  `json:"requestId"`
	Took      float32 `json:"took"`
}

type Parent struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

func Fetch() *OnCallData {
	apiKey := os.Getenv("WTF_OPS_GENIE_API_KEY")
	scheduleName := "Oversight"

	url := fmt.Sprintf("https://api.opsgenie.com/v2/schedules/%s/on-calls?scheduleIdentifierType=name&flat=true", scheduleName)

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

	var onCallData OnCallData

	if err := json.NewDecoder(resp.Body).Decode(&onCallData); err != nil {
		panic(err)
	}

	return &onCallData
}
