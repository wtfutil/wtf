package weather

import (
	"os"

	owm "github.com/briandowns/openweathermap"
)

/* -------------------- Exported Functions -------------------- */

func Fetch() *owm.CurrentWeatherData {
	apiKey := os.Getenv("WTF_OWM_API_KEY")
	vancouver := 6173331

	return currentWeather(apiKey, vancouver)
}

/* -------------------- Unexported Functions -------------------- */

func currentWeather(apiKey string, cityCode int) *owm.CurrentWeatherData {
	weather, err := owm.NewCurrent("C", "EN", apiKey)
	if err != nil {
		panic(err)
	}

	weather.CurrentByID(cityCode)

	return weather
}
