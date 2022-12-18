package flighty

// Attribution:
// Most of this code was copied from https://github.com/OroraTech/go-opensky-api/blob/main/opensky.go
//
// This uses the https://openskynetwork.github.io/opensky-api/rest.html API

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	OpenSkyURL = "https://opensky-network.org/api"

	defaultTimeout = 5
)

// Utility time wrapper struct, needed for unmarshaling unix
// timestamps from JSON objects.
type UnixTime struct {
	time.Time
}

func (t *UnixTime) UnmarshalJSON(s []byte) error {
	raw := string(s)
	if raw == "null" {
		*t = UnixTime{time.Time{}}
		return nil
	}

	unixTimestamp, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return err
	}

	*t = UnixTime{time.Unix(unixTimestamp, 0)}

	return nil
}

// Represents a single flight of an aircraft.
type Flight struct {
	ICAO24                           string   `json:"icao24"`                           // ICAO24 address of the transmitter in hex string representation.
	FirstSeen                        UnixTime `json:"firstSeen"`                        // Estimated time of departure for the flight.
	EstDepartureAirport              string   `json:"estDepartureAirport,omitempty"`    // ICAO code of the estimated departure airport. Can be nil if the airport could not be identified.
	LastSeen                         UnixTime `json:"lastSeen"`                         // Estimated time of arrival for the flight.
	EstArrivalAirport                string   `json:"estArrivalAirport,omitempty"`      // ICAO code of the estimated arrival airport. Can be nil if the airport could not be identified.
	CallSign                         string   `json:"callsign,omitempty"`               // CallSign of the vehicle. Can be nil if no callsign has been received.
	EstDepartureAirportHorizDistance int      `json:"estDepartureAirportHorizDistance"` // Horizontal distance of the last received airborne position to the estimated departure airport in meters.
	EstDepartureAirportVertDistance  int      `json:"estDepartureAirportVertDistance"`  // Vertical distance of the last received airborne position to the estimated departure airport in meters.
	EstArrivalAirportHorizDistance   int      `json:"estArrivalAirportHorizDistance"`   // Horizontal distance of the last received airborne position to the estimated arrival airport in meters.
	EstArrivalAirportVertDistance    int      `json:"estArrivalAirportVertDistance"`    // Vertical distance of the last received airborne position to the estimated arrival airport in meters.
	DepartureAirportCandidatesCount  int      `json:"departureAirportCandidatesCount"`  // Number of other possible departure airports. These are airports in short distance to EstDepartureAirport.
	ArrivalAirportCandidatesCount    int      `json:"arrivalAirportCandidatesCount"`    // Number of other possible departure airports. These are airports in short distance to EstArrivalAirport.
}

/* -------------------- OpenSky Client -------------------- */

type OpenSkyClient struct {
	username   string
	password   string
	httpClient http.Client
}

func NewOpenSkyClient(username string, password string) *OpenSkyClient {
	client := &OpenSkyClient{
		username: username,
		password: password,
		httpClient: http.Client{
			Timeout: time.Minute * defaultTimeout,
		},
	}

	return client
}

func (client *OpenSkyClient) Flight(icao24 string, begin time.Time, end time.Time) ([]*Flight, error) {
	request, err := client.newRequest("GET", fmt.Sprintf("%s/flights/aircraft", OpenSkyURL))
	if err != nil {
		return nil, err
	}

	q := request.URL.Query()

	if begin.IsZero() {
		return nil, errors.New("no begin time defined")
	}
	q.Set("begin", fmt.Sprintf("%v", begin.Unix()))

	if end.IsZero() {
		return nil, errors.New("no end time defined")
	}
	q.Set("end", fmt.Sprintf("%v", end.Unix()))

	if icao24 == "" {
		return nil, errors.New("no icao24 defined")
	}
	q.Set("icao24", icao24)

	request.URL.RawQuery = q.Encode()

	// panic(request.URL.String()) // The URL being called in the query

	flights := []*Flight{}
	err = client.doHTTP(request, &flights)

	return flights, err
}

/* -------------------- Unexported Functions -------------------- */

func (client *OpenSkyClient) newRequest(method string, url string) (*http.Request, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return request, err
	}

	if request != nil && client.username != "" && client.password != "" {
		request.SetBasicAuth(client.username, client.password)
	}

	return request, nil
}

func (client *OpenSkyClient) doHTTP(request *http.Request, responseObject interface{}) error {
	var resp *http.Response

	resp, err := client.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Parse response
	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%d: %v", resp.StatusCode, string(body))
	}

	// Parse JSON
	err = json.Unmarshal(body, responseObject)
	if err != nil {
		return err
	}

	return nil
}
