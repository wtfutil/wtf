package weather

import (
	"fmt"
	"strings"

	owm "github.com/briandowns/openweathermap"
	"github.com/wtfutil/wtf/wtf"
)

func (widget *Widget) display() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	var err string

	if widget.apiKeyValid() == false {
		err = " Environment variable WTF_OWM_API_KEY is not set\n"
	}

	cityData := widget.currentData()
	if err == "" && cityData == nil {
		err += " Weather data is unavailable: no city data\n"
	}

	if err == "" && len(cityData.Weather) == 0 {
		err += " Weather data is unavailable: no weather data\n"
	}

	title := widget.CommonSettings().Title
	setWrap := false

	var content string
	if err != "" {
		setWrap = true
		content = err
	} else {
		title = widget.buildTitle(cityData)
		_, _, width, _ := widget.View.GetRect()
		content = widget.settings.common.SigilStr(len(widget.Data), widget.Idx, width) + "\n"
		content += widget.description(cityData) + "\n\n"
		content += widget.temperatures(cityData) + "\n"
		content += widget.sunInfo(cityData)
	}

	return title, content, setWrap
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
	str := fmt.Sprintf("%8s: %4.1f° %s\n", "High", cityData.Main.TempMax, widget.settings.tempUnit)

	str += fmt.Sprintf(
		"%8s: [%s]%4.1f° %s[white]\n",
		"Current",
		widget.settings.colors.current,
		cityData.Main.Temp,
		widget.settings.tempUnit,
	)

	str += fmt.Sprintf("%8s: %4.1f° %s\n", "Low", cityData.Main.TempMin, widget.settings.tempUnit)

	return str
}

func (widget *Widget) buildTitle(cityData *owm.CurrentWeatherData) string {
	return fmt.Sprintf("%s %s", widget.emojiFor(cityData), cityData.Name)
}
