package openweathermap

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var errInvalidUVIndex = errors.New("invalid UV index value")

// UVDataPoints holds the UV specific data
type UVDataPoints struct {
	DT    int64   `json:"dt"`
	Value float64 `json:"value"`
}

// UV contains the response from the OWM UV API
type UV struct {
	Coord []float64      `json:"coord"`
	Data  []UVDataPoints `json:"data,omitempty"`
	/*Data  []struct {
		DT    int64   `json:"dt"`
		Value float64 `json:"value"`
	} `json:"data,omitempty"`*/
	DT    int64   `json:"dt,omitempty"`
	Value float64 `json:"value,omitempty"`
	Key   string
	*Settings
}

// NewUV creates a new reference to UV
func NewUV(key string, options ...Option) (*UV, error) {
	k, err := setKey(key)
	if err != nil {
		return nil, err
	}
	u := &UV{
		Key:      k,
		Settings: NewSettings(),
	}

	if err := setOptions(u.Settings, options); err != nil {
		return nil, err
	}
	return u, nil
}

// Current gets the current UV data for the given coordinates
func (u *UV) Current(coord *Coordinates) error {
	response, err := u.client.Get(fmt.Sprintf("%suvi?lat=%f&lon=%f&appid=%s", uvURL, coord.Latitude, coord.Longitude, u.Key))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&u); err != nil {
		return err
	}

	return nil
}

// Historical gets the historical UV data for the coordinates and times
func (u *UV) Historical(coord *Coordinates, start, end time.Time) error {
	response, err := u.client.Get(fmt.Sprintf("%shistory?lat=%f&lon=%f&start=%d&end=%d&appid=%s", uvURL, coord.Latitude, coord.Longitude, start.Unix(), end.Unix(), u.Key))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&u); err != nil {
		return err
	}

	return nil
}

// UVIndexInfo
type UVIndexInfo struct {
	// UVIndex holds the range of the index
	UVIndex []float64

	// MGC represents the Media graphic color
	MGC string

	// Risk of harm from unprotected sun exposure, for the average adult
	Risk string

	// RecommendedProtection contains information on what a person should
	// do when outside in the associated UVIndex
	RecommendedProtection string
}

// UVData contains data in regards to UV index ranges, rankings, and steps for protection
var UVData = []UVIndexInfo{
	{
		UVIndex: []float64{0, 2.9},
		MGC:     "Green",
		Risk:    "Low",
		RecommendedProtection: "Wear sunglasses on bright days; use sunscreen if there is snow on the ground, which reflects UV radiation, or if you have particularly fair skin.",
	},
	{
		UVIndex: []float64{3, 5.9},
		MGC:     "Yellow",
		Risk:    "Moderate",
		RecommendedProtection: "Take precautions, such as covering up, if you will be outside. Stay in shade near midday when the sun is strongest.",
	},
	{
		UVIndex: []float64{6, 7.9},
		MGC:     "Orange",
		Risk:    "High",
		RecommendedProtection: "Cover the body with sun protective clothing, use SPF 30+ sunscreen, wear a hat, reduce time in the sun within three hours of solar noon, and wear sunglasses.",
	},
	{
		UVIndex: []float64{8, 10.9},
		MGC:     "Red",
		Risk:    "Very high",
		RecommendedProtection: "Wear SPF 30+ sunscreen, a shirt, sunglasses, and a wide-brimmed hat. Do not stay in the sun for too long.",
	},
	{
		UVIndex: []float64{11},
		MGC:     "Violet",
		Risk:    "Extreme",
		RecommendedProtection: "Take all precautions: Wear SPF 30+ sunscreen, a long-sleeved shirt and trousers, sunglasses, and a very broad hat. Avoid the sun within three hours of solar noon.",
	},
}

// UVInformation provides information on the given UV data which includes the severity
// and "Recommended protection"
func (u *UV) UVInformation() ([]UVIndexInfo, error) {
	switch {
	case u.Value != 0:
		switch {
		case u.Value < 2.9:
			return []UVIndexInfo{UVData[0]}, nil
		case u.Value > 3 && u.Value < 5.9:
			return []UVIndexInfo{UVData[1]}, nil
		case u.Value > 6 && u.Value < 7.9:
			return []UVIndexInfo{UVData[2]}, nil
		case u.Value > 8 && u.Value < 10.9:
			return []UVIndexInfo{UVData[3]}, nil
		case u.Value >= 11:
			return []UVIndexInfo{UVData[4]}, nil
		default:
			return nil, errInvalidUVIndex
		}

	case len(u.Data) > 0:
		var uvi []UVIndexInfo
		for _, i := range u.Data {
			switch {
			case i.Value < 2.9:
				uvi = append(uvi, UVData[0])
			case i.Value > 3 && u.Value < 5.9:
				uvi = append(uvi, UVData[1])
			case i.Value > 6 && u.Value < 7.9:
				uvi = append(uvi, UVData[2])
			case i.Value > 8 && u.Value < 10.9:
				uvi = append(uvi, UVData[3])
			case i.Value >= 11:
				uvi = append(uvi, UVData[4])
			default:
				return nil, errInvalidUVIndex
			}
		}
	}

	return nil, nil
}
