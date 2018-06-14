package weather

import (
	"fmt"
	"strings"

	owm "github.com/briandowns/openweathermap"
	"github.com/senorprogrammer/wtf/wtf"
)

func (widget *Widget) display() {
	widget.View.Clear()

	if widget.apiKeyValid() == false {
		fmt.Fprintf(widget.View, "%s", " Environment variable WTF_OWM_API_KEY is not set")
		return
	}

	cityData := widget.currentData()
	if cityData == nil {
		fmt.Fprintf(widget.View, "%s", " Weather data is unavailable (1)")
		return
	}

	if len(cityData.Weather) == 0 {
		fmt.Fprintf(widget.View, "%s", " Weather data is unavailable (2)")
		return
	}

	widget.View.SetTitle(widget.title(cityData))

	str := wtf.SigilStr(len(widget.Data), widget.Idx, widget.View) + "\n"
	str = str + widget.description(cityData) + "\n\n"
	str = str + widget.temperatures(cityData) + "\n"
	str = str + widget.sunInfo(cityData)

	fmt.Fprintf(widget.View, "%s", str)
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
	tempUnit := Config.UString("wtf.mods.weather.tempUnit", "C")

	str := fmt.Sprintf("%8s: %4.1f° %s\n", "High", cityData.Main.TempMax, tempUnit)

	str = str + fmt.Sprintf(
		"%8s: [%s]%4.1f° %s[white]\n",
		"Current",
		Config.UString("wtf.mods.weather.colors.current", "green"),
		cityData.Main.Temp,
		tempUnit,
	)

	str = str + fmt.Sprintf("%8s: %4.1f° %s\n", "Low", cityData.Main.TempMin, tempUnit)

	return str
}

func (widget *Widget) title(cityData *owm.CurrentWeatherData) string {
	return fmt.Sprintf(" %s %s ", widget.icon(cityData), cityData.Name)
}
