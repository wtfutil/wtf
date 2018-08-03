package weather

import (
	owm "github.com/briandowns/openweathermap"
)

var weatherEmoji = map[string]string{
	"default":                     "💥",
	"broken clouds":               "🌤",
	"clear":                       " ",
	"clear sky":                   " ",
	"cloudy":                      "⛅️",
	"few clouds":                  "🌤",
	"fog":                         "🌫",
	"haze":                        "🌫",
	"heavy intensity rain":        "💦",
	"heavy rain":                  "💦",
	"heavy snow":                  "⛄️",
	"light intensity shower rain": "☔️",
	"light rain":                  "🌦",
	"light shower snow":           "🌦⛄️",
	"light snow":                  "🌨",
	"mist":                        "🌬",
	"moderate rain":               "🌧",
	"moderate snow":               "🌨",
	"overcast":                    "🌥",
	"overcast clouds":             "🌥",
	"partly cloudy":               "🌤",
	"scattered clouds":            "🌤",
	"shower rain":                 "☔️",
	"smoke":                       "🔥",
	"snow":                        "❄️",
	"sunny":                       "☀️",
	"thunderstorm":                "⛈",
}

func (widget *Widget) emojiFor(data *owm.CurrentWeatherData) string {
	if len(data.Weather) == 0 {
		return ""
	}

	emoji := weatherEmoji[data.Weather[0].Description]
	if emoji == "" {
		emoji = weatherEmoji["default"]
	}

	return emoji
}
