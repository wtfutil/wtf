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
	"net/http"
	"os"
)

// IconData holds the relevant info for linking icons to conditions.
type IconData struct {
	Condition string
	Day       string
	Night     string
}

// ConditionData holds data structure for weather conditions information.
type ConditionData struct {
	ID      int
	Meaning string
	Icon1   string
	Icon2   string
}

// RetrieveIcon will get the specified icon from the API.
func RetrieveIcon(destination, iconFile string) (int64, error) {
	fullFilePath := fmt.Sprintf("%s/%s", destination, iconFile)

	// Check to see if we've already gotten that icon file.  If so, use it
	// rather than getting it again.
	if _, err := os.Stat(fullFilePath); err != nil {
		response, err := http.Get(fmt.Sprintf(iconURL, iconFile))
		if err != nil {
			return 0, err
		}
		defer response.Body.Close()

		// Create the icon file
		out, err := os.Create(fullFilePath)
		if err != nil {
			return 0, err
		}
		defer out.Close()

		// Fill the empty file with the actual content
		n, err := io.Copy(out, response.Body)
		if err != nil {
			return 0, err
		}
		return n, nil
	}
	return 0, nil
}

// IconList is a slice of IconData pointers
var IconList = []*IconData{
	{Condition: "clear sky", Day: "01d.png", Night: "01n.png"},
	{Condition: "few clouds", Day: "02d.png", Night: "02n.png"},
	{Condition: "scattered clouds", Day: "03d.png", Night: "03n.png"},
	{Condition: "broken clouds", Day: "04d.png", Night: "04n.png"},
	{Condition: "shower rain", Day: "09d.png", Night: "09n.png"},
	{Condition: "rain", Day: "10d.png", Night: "10n.png"},
	{Condition: "thunderstorm", Day: "11d.png", Night: "11n.png"},
	{Condition: "snow", Day: "13d.png", Night: "13n.png"},
	{Condition: "mist", Day: "50d.png", Night: "50n.png"},
}

// ThunderstormConditions is a slice of ConditionData pointers
var ThunderstormConditions = []*ConditionData{
	{ID: 200, Meaning: "thunderstorm with light rain", Icon1: "11d.png"},
	{ID: 201, Meaning: "thunderstorm with rain", Icon1: "11d.png"},
	{ID: 202, Meaning: "thunderstorm with heavy rain", Icon1: "11d.png"},
	{ID: 210, Meaning: "light thunderstorm", Icon1: "11d.png"},
	{ID: 211, Meaning: "thunderstorm", Icon1: "11d.png"},
	{ID: 212, Meaning: "heavy thunderstorm", Icon1: "11d.png"},
	{ID: 221, Meaning: "ragged thunderstorm", Icon1: "11d.png"},
	{ID: 230, Meaning: "thunderstorm with light drizzle", Icon1: "11d.png"},
	{ID: 231, Meaning: "thunderstorm with drizzle", Icon1: "11d.png"},
	{ID: 232, Meaning: "thunderstorm with heavy drizzle", Icon1: "11d.png"},
}

// DrizzleConditions is a slice of ConditionData pointers
var DrizzleConditions = []*ConditionData{
	{ID: 300, Meaning: "light intensity drizzle", Icon1: "09d.png"},
	{ID: 301, Meaning: "drizzle", Icon1: "09d.png"},
	{ID: 302, Meaning: "heavy intensity drizzle", Icon1: "09d.png"},
	{ID: 310, Meaning: "light intensity drizzle rain", Icon1: "09d.png"},
	{ID: 311, Meaning: "drizzle rain", Icon1: "09d.png"},
	{ID: 312, Meaning: "heavy intensity drizzle rain", Icon1: "09d.png"},
	{ID: 313, Meaning: "shower rain and drizzle", Icon1: "09d.png"},
	{ID: 314, Meaning: "heavy shower rain and drizzle", Icon1: "09d.png"},
	{ID: 321, Meaning: "shower drizzle", Icon1: "09d.png"},
}

// RainConditions is a slice of ConditionData pointers
var RainConditions = []*ConditionData{
	{ID: 500, Meaning: "light rain", Icon1: "09d.png"},
	{ID: 501, Meaning: "moderate rain", Icon1: "09d.png"},
	{ID: 502, Meaning: "heavy intensity rain", Icon1: "09d.png"},
	{ID: 503, Meaning: "very heavy rain", Icon1: "09d.png"},
	{ID: 504, Meaning: "extreme rain", Icon1: "09d.png"},
	{ID: 511, Meaning: "freezing rain", Icon1: "13d.png"},
	{ID: 520, Meaning: "light intensity shower rain", Icon1: "09d.png"},
	{ID: 521, Meaning: "shower rain", Icon1: "09d.png"},
	{ID: 522, Meaning: "heavy intensity shower rain", Icon1: "09d.png"},
	{ID: 531, Meaning: "ragged shower rain", Icon1: "09d.png"},
}

// SnowConditions is a slice of ConditionData pointers
var SnowConditions = []*ConditionData{
	{ID: 600, Meaning: "light snow", Icon1: "13d.png"},
	{ID: 601, Meaning: "snow", Icon1: "13d.png"},
	{ID: 602, Meaning: "heavy snow", Icon1: "13d.png"},
	{ID: 611, Meaning: "sleet", Icon1: "13d.png"},
	{ID: 612, Meaning: "shower sleet", Icon1: "13d.png"},
	{ID: 615, Meaning: "light rain and snow", Icon1: "13d.png"},
	{ID: 616, Meaning: "rain and snow", Icon1: "13d.png"},
	{ID: 620, Meaning: "light shower snow", Icon1: "13d.png"},
	{ID: 621, Meaning: "shower snow", Icon1: "13d.png"},
	{ID: 622, Meaning: "heavy shower snow", Icon1: "13d.png"},
}

// AtmosphereConditions is a slice of ConditionData pointers
var AtmosphereConditions = []*ConditionData{
	{ID: 701, Meaning: "mist", Icon1: "50d.png"},
	{ID: 711, Meaning: "smoke", Icon1: "50d.png"},
	{ID: 721, Meaning: "haze", Icon1: "50d.png"},
	{ID: 731, Meaning: "sand, dust whirls", Icon1: "50d.png"},
	{ID: 741, Meaning: "fog", Icon1: "50d.png"},
	{ID: 751, Meaning: "sand", Icon1: "50d.png"},
	{ID: 761, Meaning: "dust", Icon1: "50d.png"},
	{ID: 762, Meaning: "volcanic ash", Icon1: "50d.png"},
	{ID: 771, Meaning: "squalls", Icon1: "50d.png"},
	{ID: 781, Meaning: "tornado", Icon1: "50d.png"},
}

// CloudConditions is a slice of ConditionData pointers
var CloudConditions = []*ConditionData{
	{ID: 800, Meaning: "clear sky", Icon1: "01d.png", Icon2: "01n.png"},
	{ID: 801, Meaning: "few clouds", Icon1: "02d.png", Icon2: " 02n.png"},
	{ID: 802, Meaning: "scattered clouds", Icon1: "03d.png", Icon2: "03d.png"},
	{ID: 803, Meaning: "broken clouds", Icon1: "04d.png", Icon2: "03d.png"},
	{ID: 804, Meaning: "overcast clouds", Icon1: "04d.png", Icon2: "04d.png"},
}

// ExtremeConditions is a slice of ConditionData pointers
var ExtremeConditions = []*ConditionData{
	{ID: 900, Meaning: "tornado", Icon1: ""},
	{ID: 901, Meaning: "tropical storm", Icon1: ""},
	{ID: 902, Meaning: "hurricane", Icon1: ""},
	{ID: 903, Meaning: "cold", Icon1: ""},
	{ID: 904, Meaning: "hot", Icon1: ""},
	{ID: 905, Meaning: "windy", Icon1: ""},
	{ID: 906, Meaning: "hail", Icon1: ""},
}

// AdditionalConditions is a slive of ConditionData pointers
var AdditionalConditions = []*ConditionData{
	{ID: 951, Meaning: "calm", Icon1: ""},
	{ID: 952, Meaning: "light breeze", Icon1: ""},
	{ID: 953, Meaning: "gentle breeze", Icon1: ""},
	{ID: 954, Meaning: "moderate breeze", Icon1: ""},
	{ID: 955, Meaning: "fresh breeze", Icon1: ""},
	{ID: 956, Meaning: "strong breeze", Icon1: ""},
	{ID: 957, Meaning: "high wind, near gale", Icon1: ""},
	{ID: 958, Meaning: "gale", Icon1: ""},
	{ID: 959, Meaning: "severe gale", Icon1: ""},
	{ID: 960, Meaning: "storm", Icon1: ""},
	{ID: 961, Meaning: "violent storm", Icon1: ""},
	{ID: 962, Meaning: "hurricane", Icon1: ""},
}
