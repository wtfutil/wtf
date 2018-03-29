package weather

import (
	"os"

	owm "github.com/briandowns/openweathermap"
)

func Fetch() *owm.CurrentWeatherData {
	w, err := owm.NewCurrent("C", "EN", os.Getenv("WTF_OWM_API_KEY"))
	if err != nil {
		panic(err)
	}

	//w.CurrentByName("Toronto,ON")
	w.CurrentByID(6173331)

	return w
}
