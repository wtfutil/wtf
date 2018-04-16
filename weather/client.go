package weather

import (
	//"fmt"
	"os"

	owm "github.com/briandowns/openweathermap"
)

/* -------------------- Exported Functions -------------------- */

func Fetch(cityids []int) []*owm.CurrentWeatherData {
	apiKey := os.Getenv("WTF_OWM_API_KEY")

	data := []*owm.CurrentWeatherData{}

	for _, cityID := range cityids {
		result, err := currentWeather(apiKey, cityID)
		if err == nil {
			data = append(data, result)
		}
	}

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
