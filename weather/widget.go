package weather

import (
	"fmt"
	"strings"
	"time"

	owm "github.com/briandowns/openweathermap"
	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

var Config *config.Config

type Widget struct {
	wtf.BaseWidget
	View *tview.TextView
}

func NewWidget() *Widget {
	widget := Widget{
		BaseWidget: wtf.BaseWidget{
			Name:        "Weather",
			RefreshedAt: time.Now(),
			RefreshInt:  Config.UInt("wtf.weather.refreshInterval", 900),
		},
	}

	widget.addView()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	data := Fetch(Config.UInt("wtf.weather.cityId", 6176823))

	widget.View.SetTitle(fmt.Sprintf(" %s Weather - %s ", icon(data), data.Name))
	widget.RefreshedAt = time.Now()

	widget.View.Clear()
	fmt.Fprintf(widget.View, "%s", widget.contentFrom(data))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) addView() {
	view := tview.NewTextView()

	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorGray)
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name)
	view.SetWrap(false)

	widget.View = view
}

func (widget *Widget) contentFrom(data *owm.CurrentWeatherData) string {
	str := "\n"

	descs := []string{}
	for _, weather := range data.Weather {
		descs = append(descs, fmt.Sprintf(" %s", weather.Description))
	}

	str = str + strings.Join(descs, ",") + "\n\n"

	tempUnit := Config.UString("wtf.weather.tempUnit", "C")

	str = str + fmt.Sprintf("%8s: %4.1f° %s\n", "High", data.Main.TempMax, tempUnit)
	str = str + fmt.Sprintf("%8s: [green]%4.1f° %s[white]\n", "Current", data.Main.Temp, tempUnit)
	str = str + fmt.Sprintf("%8s: %4.1f° %s\n", "Low", data.Main.TempMin, tempUnit)

	str = str + "\n"
	str = str + fmt.Sprintf(
		" Sunrise: %s      Sunset: %s\n",
		wtf.UnixTime(int64(data.Sys.Sunrise)).Format("15:04"),
		wtf.UnixTime(int64(data.Sys.Sunset)).Format("15:04"),
	)

	return str
}

// icon returns an emoji for the current weather
// src: https://github.com/chubin/wttr.in/blob/master/share/translations/en.txt
// Note: these only work for English weather status. Sorry about that
func icon(data *owm.CurrentWeatherData) string {
	var icon string

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
	case "snow":
		icon = "❄️"
	case "sunny":
		icon = "☀️"
	default:
		icon = "💥"
	}

	return icon
}

func (widget *Widget) refreshedAt() string {
	return widget.RefreshedAt.Format("15:04:05")
}
