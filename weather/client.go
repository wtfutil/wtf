package weather

import (
	"os"

	owm "github.com/briandowns/openweathermap"
)

/* -------------------- Exported Functions -------------------- */

func Fetch(cityID int) *owm.CurrentWeatherData {
	apiKey := os.Getenv("WTF_OWM_API_KEY")
	return currentWeather(apiKey, cityID)
}

/* -------------------- Unexported Functions -------------------- */

func currentWeather(apiKey string, cityCode int) *owm.CurrentWeatherData {
	weather, err := owm.NewCurrent(Config.UString("wtf.weather.tempUnit", "C"), Config.UString("wtf.weather.language", "EN"), apiKey)
	if err != nil {
		panic(err)
	}

	weather.CurrentByID(cityCode)

	return weather
}
