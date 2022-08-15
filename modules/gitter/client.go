package gitter

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/wtfutil/wtf/utils"
)

func GetMessages(roomId string, numberOfMessages int, apiToken string) ([]Message, error) {
	var messages []Message

	resp, err := apiRequest("rooms/"+roomId+"/chatMessages?limit="+strconv.Itoa(numberOfMessages), apiToken)
	if err != nil {
		return nil, err
	}

	err = utils.ParseJSON(&messages, resp.Body)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func GetRoom(roomUri, apiToken string) (*Room, error) {
	var rooms Rooms

	resp, err := apiRequest("rooms?q="+roomUri, apiToken)
	if err != nil {
		return nil, err
	}

	err = utils.ParseJSON(&rooms, resp.Body)
	if err != nil {
		return nil, err
	}

	for _, room := range rooms.Results {
		if room.URI == roomUri {
			return &room, nil
		}
	}

	return nil, nil
}

/* -------------------- Unexported Functions -------------------- */

var (
	apiBaseURL = "https://api.gitter.im/v1/"
)

func apiRequest(path, apiToken string) (*http.Response, error) {
	req, err := http.NewRequest("GET", apiBaseURL+path, http.NoBody)
	if err != nil {
		return nil, err
	}

	bearer := fmt.Sprintf("Bearer %s", apiToken)
	req.Header.Add("Authorization", bearer)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf(resp.Status)
	}

	return resp, nil
}
