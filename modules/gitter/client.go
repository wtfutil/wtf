package gitter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wtfutil/wtf/logger"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetMessages(roomId string, numberOfMessages int, apiToken string) ([]Message, error) {
	var messages []Message

	resp, err := apiRequest("rooms/"+roomId+"/chatMessages?limit="+strconv.Itoa(numberOfMessages), apiToken)
	if err != nil {
		return nil, err
	}

	parseJson(&messages, resp.Body)

	return messages, nil
}

func GetRoom(roomUri, apiToken string) (*Room, error) {
	var rooms Rooms

	resp, err := apiRequest("rooms?q="+roomUri, apiToken)
	if err != nil {
		return nil, err
	}

	parseJson(&rooms, resp.Body)

	for _, room := range rooms.Results {
		logger.Log(fmt.Sprintf("room: %s", room))
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
	req, err := http.NewRequest("GET", apiBaseURL+path, nil)
	bearer := fmt.Sprintf("Bearer %s", apiToken)
	req.Header.Add("Authorization", bearer)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf(resp.Status)
	}

	return resp, nil
}

func parseJson(obj interface{}, text io.Reader) {
	jsonStream, err := ioutil.ReadAll(text)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonStream))

	for {
		if err := decoder.Decode(obj); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
}
