package weather

import (
	"os"

	owm "github.com/briandowns/openweathermap"
)

func Fetch() *owm.CurrentWeatherData {
	apiKey := os.Getenv("WTF_OWM_API_KEY")

	weather, err := owm.NewCurrent("C", "EN", apiKey)
	if err != nil {
		panic(err)
	}

	weather.CurrentByID(6173331)

	return weather
}
