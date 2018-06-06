package openweathermap

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// DateTimeAliases holds the alias the pollution API supports in lieu
// of an ISO 8601 timestamp
var DateTimeAliases = []string{"current"}

// ValidAlias checks to make sure the given alias is a valid one
func ValidAlias(alias string) bool {
	for _, i := range DateTimeAliases {
		if i == alias {
			return true
		}
	}
	return false
}

// PollutionData holds the pollution specific data from the call
type PollutionData struct {
	Precision float64 `json:"precision"`
	Pressure  float64 `json:"pressure"`
	Value     float64 `json:"value"`
}

// PollutionParameters holds the parameters needed to make
// a call to the pollution API
type PollutionParameters struct {
	Location Coordinates
	Datetime string // this should be either ISO 8601 or an alias
}

// Pollution holds the data returnd from the pollution API
type Pollution struct {
	Time     string          `json:"time"`
	Location Coordinates     `json:"location"`
	Data     []PollutionData `json:"data"`
	Key      string
	*Settings
}

// NewPollution creates a new reference to Pollution
func NewPollution(key string, options ...Option) (*Pollution, error) {
	k, err := setKey(key)
	if err != nil {
		return nil, err
	}
	p := &Pollution{
		Key:      k,
		Settings: NewSettings(),
	}

	if err := setOptions(p.Settings, options); err != nil {
		return nil, err
	}
	return p, nil
}

// PollutionByParams gets the pollution data based on the given parameters
func (p *Pollution) PollutionByParams(params *PollutionParameters) error {
	url := fmt.Sprintf("%s%s,%s/%s.json?appid=%s",
		pollutionURL,
		strconv.FormatFloat(params.Location.Latitude, 'f', -1, 64),
		strconv.FormatFloat(params.Location.Longitude, 'f', -1, 64),
		params.Datetime,
		p.Key)
	response, err := p.client.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&p); err != nil {
		return err
	}

	return nil
}
