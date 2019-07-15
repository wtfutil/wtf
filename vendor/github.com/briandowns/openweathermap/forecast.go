// Copyright 2015 Brian J. Downs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openweathermap

import (
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
)

// ForecastSys area population
type ForecastSys struct {
	Population int `json:"population"`
}

// Temperature holds returned termperate sure stats
type Temperature struct {
	Day   float64 `json:"day"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

// City data for given location
type City struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Coord      Coordinates `json:"coord"`
	Country    string      `json:"country"`
	Population int         `json:"population"`
	Sys        ForecastSys `json:"sys"`
}

type ForecastWeather interface {
	DailyByName(location string, days int) error
	DailyByCoordinates(location *Coordinates, days int) error
	DailyByID(id, days int) error
}

// json served by OWM API can take different forms, so all of them must be matched
// by corresponding data type and unmarshall method
type ForecastWeatherJson interface {
	Decode(r io.Reader) error
}

type ForecastWeatherData struct {
	Unit    string
	Lang    string
	Key     string
	baseURL string
	*Settings
	ForecastWeatherJson
}

// NewForecast returns a new HistoricalWeatherData pointer with
// the supplied arguments.
func NewForecast(forecastType, unit, lang, key string, options ...Option) (*ForecastWeatherData, error) {
	unitChoice := strings.ToUpper(unit)
	langChoice := strings.ToUpper(lang)

	if forecastType != "16" && forecastType != "5" {
		return nil, errForecastUnavailable
	}

	if !ValidDataUnit(unitChoice) {
		return nil, errUnitUnavailable
	}

	if !ValidLangCode(langChoice) {
		return nil, errLangUnavailable
	}

	settings := NewSettings()
	if err := setOptions(settings, options); err != nil {
		return nil, err
	}

	var err error
	k, err := setKey(key)
	if err != nil {
		return nil, err
	}
	forecastData := ForecastWeatherData{
		Unit:     DataUnits[unitChoice],
		Lang:     langChoice,
		Key:      k,
		Settings: settings,
	}

	if forecastType == "16" {
		forecastData.baseURL = forecast16Base
		forecastData.ForecastWeatherJson = &Forecast16WeatherData{}
	} else {
		forecastData.baseURL = forecast5Base
		forecastData.ForecastWeatherJson = &Forecast5WeatherData{}
	}

	return &forecastData, nil
}

// DailyByName will provide a forecast for the location given for the
// number of days given.
func (f *ForecastWeatherData) DailyByName(location string, days int) error {
	response, err := f.client.Get(fmt.Sprintf(f.baseURL, f.Key, fmt.Sprintf("%s=%s", "q", url.QueryEscape(location)), f.Unit, f.Lang, days))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return f.ForecastWeatherJson.Decode(response.Body)
}

// DailyByCoordinates will provide a forecast for the coordinates ID give
// for the number of days given.
func (f *ForecastWeatherData) DailyByCoordinates(location *Coordinates, days int) error {
	response, err := f.client.Get(fmt.Sprintf(f.baseURL, f.Key, fmt.Sprintf("lat=%f&lon=%f", location.Latitude, location.Longitude), f.Unit, f.Lang, days))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return f.ForecastWeatherJson.Decode(response.Body)
}

// DailyByID will provide a forecast for the location ID give for the
// number of days given.
func (f *ForecastWeatherData) DailyByID(id, days int) error {
	response, err := f.client.Get(fmt.Sprintf(f.baseURL, f.Key, fmt.Sprintf("%s=%s", "id", strconv.Itoa(id)), f.Unit, f.Lang, days))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return f.ForecastWeatherJson.Decode(response.Body)
}
