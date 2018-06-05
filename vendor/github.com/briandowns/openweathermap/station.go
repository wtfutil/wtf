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
	"net/http"
	"net/url"
)

// Slice of type string of the valid parameters to be sent from a station.
// The API refers to this data as the "Weather station data transmission protocol"
var StationDataParameters = []string{
	"wind_dir",   // Wind direction
	"wind_speed", // Wind speed
	"wind_gust",  // Wind gust speed
	"temp",       // Temperature
	"humidity",   // Relative humidty
	"pressure",   // Atmospheric pressure
	"rain_1h",    // Rain in the last hour
	"rain_24h",   // Rain in the last 24 hours
	"rain_today", // Rain since midnight
	"snow",       // Snow in the last 24 hours
	"lum",        // Brightness
	"lat",        // Latitude
	"long",       // Longitude
	"alt",        // Altitude
	"radiation",  // Radiation
	"dewpoint",   // Dew point
	"uv",         // UV index
	"name",       // Weather station name
}

// ValidateStationDataParameter will make sure that whatever parameter
// supplied is one that can actually be used in the POST request.
func ValidateStationDataParameter(param string) bool {
	for _, p := range StationDataParameters {
		if param == p {
			return true
		}
	}

	return false
}

// ConvertToURLValues will convert a map to a url.Values instance. We're
// taking a map[string]string instead of something more type specific since
// the url.Values instance only takes strings to create the URL values.
func ConvertToURLValues(data map[string]string) string {
	v := url.Values{}

	for key, val := range data {
		v.Set(key, val)
	}

	return v.Encode()
}

// SendStationData will send an instance the provided url.Values to the
// provided URL.
func SendStationData(data url.Values) {
	resp, err := http.PostForm(dataPostURL, data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Body)
}
