package openweathermap

import (
	"encoding/json"
	"io"
	"strings"
	"time"
)

type DtTxt struct {
	time.Time
}

func (dt *DtTxt) UnmarshalJSON(b []byte) error {
	t, err := time.Parse("2006-01-02 15:04:05", strings.Trim(string(b), "\""))
	dt.Time = t
	return err
}

func (t *DtTxt) MarshalJSON() ([]byte, error) {
	return json.Marshal(t)
}

// Forecast5WeatherList holds specific query data
type Forecast5WeatherList struct {
	Dt      int       `json:"dt"`
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
	Clouds  Clouds    `json:"clouds"`
	Wind    Wind      `json:"wind"`
	Rain    Rain      `json:"rain"`
	Snow    Snow      `json:"snow"`
	DtTxt   DtTxt     `json:"dt_txt"`
}

// Forecast5WeatherData will hold returned data from queries
type Forecast5WeatherData struct {
	// COD     string                `json:"cod"`
	// Message float64               `json:"message"`
	City City                   `json:"city"`
	Cnt  int                    `json:"cnt"`
	List []Forecast5WeatherList `json:"list"`
}

func (f *Forecast5WeatherData) Decode(r io.Reader) error {
	if err := json.NewDecoder(r).Decode(&f); err != nil {
		return err
	}
	return nil
}
