package weather

import (
	"fmt"
	"strings"

	owm "github.com/briandowns/openweathermap"
	"github.com/senorprogrammer/wtf/wtf"
)

func (widget *Widget) display() {
	widget.View.Clear()

	cityData := widget.currentData()

	if len(cityData.Weather) == 0 {
		fmt.Fprintf(widget.View, "%s", " Weather data is unavailable.")
		return
	}

	widget.View.SetTitle(widget.title(cityData))

	str := widget.tickMarks(widget.Data) + "\n"
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
		" Rise: %s    Set: %s",
		wtf.UnixTime(int64(cityData.Sys.Sunrise)).Format("15:04 MST"),
		wtf.UnixTime(int64(cityData.Sys.Sunset)).Format("15:04 MST"),
	)
}

func (widget *Widget) temperatures(cityData *owm.CurrentWeatherData) string {
	tempUnit := Config.UString("wtf.mods.weather.tempUnit", "C")

	str := fmt.Sprintf("%8s: %4.1f° %s\n", "High", cityData.Main.TempMax, tempUnit)
	str = str + fmt.Sprintf("%8s: [green]%4.1f° %s[white]\n", "Current", cityData.Main.Temp, tempUnit)
	str = str + fmt.Sprintf("%8s: %4.1f° %s\n", "Low", cityData.Main.TempMin, tempUnit)

	return str
}

func (widget *Widget) tickMarks(data []*owm.CurrentWeatherData) string {
	str := ""

	if len(data) > 1 {
		marks := strings.Repeat("*", len(data))
		marks = marks[:widget.Idx] + "_" + marks[widget.Idx+1:]

		str = "[lightblue]" + fmt.Sprintf(wtf.RightAlignFormat(widget.View), marks) + "[white]"
	}

	return str
}

func (widget *Widget) title(cityData *owm.CurrentWeatherData) string {
	return fmt.Sprintf(" %s %s ", widget.icon(cityData), cityData.Name)
}
