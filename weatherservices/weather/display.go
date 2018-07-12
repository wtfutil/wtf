package weather

import (
	"fmt"
	"strings"

	owm "github.com/briandowns/openweathermap"
	"github.com/senorprogrammer/wtf/wtf"
)

func (widget *Widget) display() {

	if widget.apiKeyValid() == false {
		widget.View.SetText(" Environment variable WTF_OWM_API_KEY is not set")
		return
	}

	cityData := widget.currentData()
	if cityData == nil {
		widget.View.SetText(" Weather data is unavailable: no city data")
		return
	}

	if len(cityData.Weather) == 0 {
		widget.View.SetText(" Weather data is unavailable: no weather data")
		return
	}

	widget.View.SetTitle(widget.title(cityData))

	content := wtf.SigilStr(len(widget.Data), widget.Idx, widget.View) + "\n"
	content = content + widget.description(cityData) + "\n\n"
	content = content + widget.temperatures(cityData) + "\n"
	content = content + widget.sunInfo(cityData)

	widget.View.SetText(content)
}

func (widget *Widget) description(cityData *owm.CurrentWeatherData) string {
	descs := []string{}
	for _, weather := range cityData.Weather {
		descs = append(descs, fmt.Sprintf(" %s", weather.Description))
	}

	return strings.Join(descs, ",")
}

func (widget *Widget) sunInfo(cityData *owm.CurrentWeatherData) string {
	return fmt.Sprintf(
		" Rise: %s   Set: %s",
		wtf.UnixTime(int64(cityData.Sys.Sunrise)).Format("15:04 MST"),
		wtf.UnixTime(int64(cityData.Sys.Sunset)).Format("15:04 MST"),
	)
}

func (widget *Widget) temperatures(cityData *owm.CurrentWeatherData) string {
	tempUnit := wtf.Config.UString("wtf.mods.weather.tempUnit", "C")

	str := fmt.Sprintf("%8s: %4.1f° %s\n", "High", cityData.Main.TempMax, tempUnit)

	str = str + fmt.Sprintf(
		"%8s: [%s]%4.1f° %s[white]\n",
		"Current",
		wtf.Config.UString("wtf.mods.weather.colors.current", "green"),
		cityData.Main.Temp,
		tempUnit,
	)

	str = str + fmt.Sprintf("%8s: %4.1f° %s\n", "Low", cityData.Main.TempMin, tempUnit)

	return str
}

func (widget *Widget) title(cityData *owm.CurrentWeatherData) string {
	return fmt.Sprintf(" %s  %s ", widget.icon(cityData), cityData.Name)
}
