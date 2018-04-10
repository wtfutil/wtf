package weather

import (
	//"fmt"
	"os"

	owm "github.com/briandowns/openweathermap"
)

/* -------------------- Exported Functions -------------------- */

func Fetch(cityID int) *owm.CurrentWeatherData {
	apiKey := os.Getenv("WTF_OWM_API_KEY")
	data, _ := currentWeather(apiKey, cityID)

	return data
}

/* -------------------- Unexported Functions -------------------- */

func currentWeather(apiKey string, cityCode int) (*owm.CurrentWeatherData, error) {
	weather, err := owm.NewCurrent(Config.UString("wtf.mods.weather.tempUnit", "C"), Config.UString("wtf.mods.weather.language", "EN"), apiKey)
	if err != nil {
		return nil, err
	}

	weather.CurrentByID(cityCode)

	return weather, nil
}
