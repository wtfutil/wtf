package grafana

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"

	"github.com/wtfutil/wtf/utils"
)

type AlertState int

const (
	Alerting AlertState = iota
	Pending
	NoData
	Paused
	Ok
)

var toString = map[AlertState]string{
	Alerting: "alerting",
	Pending:  "pending",
	NoData:   "no_data",
	Paused:   "paused",
	Ok:       "ok",
}

var toID = map[string]AlertState{
	"alerting": Alerting,
	"pending":  Pending,
	"no_data":  NoData,
	"paused":   Paused,
	"ok":       Ok,
}

// MarshalJSON marshals the enum as a quoted json string
func (s AlertState) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *AlertState) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// if we somehow get an invalid value we'll end up in the alerting state
	*s = toID[j]
	return nil
}

type Alert struct {
	Name  string     `json:"name"`
	State AlertState `json:"state"`
	URL   string     `json:"url"`
}

type Client struct {
	apiKey  string
	baseURI string
}

func NewClient(settings *Settings) *Client {
	return &Client{
		apiKey:  settings.apiKey,
		baseURI: settings.baseURI,
	}
}

func (client *Client) Alerts() ([]Alert, error) {
	// query the alerts API of Grafana https://grafana.com/docs/grafana/latest/http_api/alerting/
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/alerts", client.baseURI), http.NoBody)
	if err != nil {
		return nil, err
	}

	if client.apiKey != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.apiKey))
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != 200 {
		msg := struct {
			Msg string `json:"message"`
		}{}
		err = utils.ParseJSON(&msg, res.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(msg.Msg)
	}

	var out []Alert
	err = utils.ParseJSON(&out, res.Body)
	if err != nil {
		return nil, err
	}

	sort.SliceStable(out, func(i, j int) bool {
		return out[i].State < out[j].State
	})

	return out, nil
}
