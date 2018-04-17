package weather

import (
	"time"

	owm "github.com/briandowns/openweathermap"
	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

var Config *config.Config

type Widget struct {
	wtf.TextWidget

	Idx int
	Data    []*owm.CurrentWeatherData
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Weather ", "weather"),
		Idx:    0,
	}

	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.Data = Fetch(wtf.ToInts(Config.UList("wtf.mods.weather.cityids", widget.defaultCityCodes())))

	widget.display(widget.Data)
	widget.RefreshedAt = time.Now()
}

func (widget *Widget) Next() {
	widget.Idx = widget.Idx + 1
	if widget.Idx == len(widget.Data) {
		widget.Idx = 0
	}

	widget.display(widget.Data)
}

func (widget *Widget) Prev() {
	widget.Idx = widget.Idx - 1
	if widget.Idx < 0 {
		widget.Idx = len(widget.Data) - 1
	}

	widget.display(widget.Data)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) currentCityData(data []*owm.CurrentWeatherData) *owm.CurrentWeatherData {
	return data[widget.Idx]
}

func (widget *Widget) defaultCityCodes() []interface{} {
	defaultArr := []int{6176823, 360630, 3413829}

	var defaults []interface{} = make([]interface{}, len(defaultArr))
	for i, d := range defaultArr {
		defaults[i] = d
	}

	return defaults
}

// icon returns an emoji for the current weather
// src: https://github.com/chubin/wttr.in/blob/master/share/translations/en.txt
// Note: these only work for English weather status. Sorry about that
//
// FIXME: Move these into a configuration file so they can be changed without a compile
func (widget *Widget) icon(data *owm.CurrentWeatherData) string {
	var icon string

	if len(data.Weather) == 0 {
		return ""
	}

	switch data.Weather[0].Description {
	case "broken clouds":
		icon = "☁️"
	case "clear":
		icon = "☀️"
	case "clear sky":
		icon = "☀️ "
	case "cloudy":
		icon = "⛅️"
	case "few clouds":
		icon = "🌤"
	case "fog":
		icon = "🌫"
	case "haze":
		icon = "🌫"
	case "heavy rain":
		icon = "💦"
	case "heavy snow":
		icon = "⛄️"
	case "light intensity shower rain":
		icon = "☔️"
	case "light rain":
		icon = "🌦"
	case "light snow":
		icon = "🌨"
	case "mist":
		icon = "🌬"
	case "moderate rain":
		icon = "🌧"
	case "moderate snow":
		icon = "🌨"
	case "overcast":
		icon = "🌥"
	case "overcast clouds":
		icon = "🌥"
	case "partly cloudy":
		icon = "🌤"
	case "scattered clouds":
		icon = "☁️"
	case "shower rain":
		icon = "☔️"
	case "snow":
		icon = "❄️"
	case "sunny":
		icon = "☀️"
	default:
		icon = "💥"
	}

	return icon
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyLeft:
		widget.Prev()
	case tcell.KeyRight:
		widget.Next()
	default:
		return event
	}

	return event
}

func (widget *Widget) refreshedAt() string {
	return widget.RefreshedAt.Format("15:04:05")
}
