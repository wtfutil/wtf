package openweathermap

import (
	"encoding/json"
	"io"
)

// Forecast16WeatherList holds specific query data
type Forecast16WeatherList struct {
	Dt       int         `json:"dt"`
	Temp     Temperature `json:"temp"`
	Pressure float64     `json:"pressure"`
	Humidity int         `json:"humidity"`
	Weather  []Weather   `json:"weather"`
	Speed    float64     `json:"speed"`
	Deg      int         `json:"deg"`
	Clouds   int         `json:"clouds"`
	Snow     float64     `json:"snow"`
	Rain     float64     `json:"rain"`
}

// Forecast16WeatherData will hold returned data from queries
type Forecast16WeatherData struct {
	COD     int                     `json:"cod"`
	Message string                  `json:"message"`
	City    City                    `json:"city"`
	Cnt     int                     `json:"cnt"`
	List    []Forecast16WeatherList `json:"list"`
}

func (f *Forecast16WeatherData) Decode(r io.Reader) error {
	if err := json.NewDecoder(r).Decode(&f); err != nil {
		return err
	}
	return nil
}
