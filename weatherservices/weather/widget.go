package weather

import (
	"os"

	owm "github.com/briandowns/openweathermap"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

const HelpText = `
  Keyboard commands for Weather:

    /: Show/hide this help window
    h: Previous weather location
    l: Next weather location

    arrow left:  Previous weather location
    arrow right: Next weather location
`

// Widget is the container for weather data.
type Widget struct {
	wtf.HelpfulWidget
	wtf.TextWidget

	APIKey string
	Data   []*owm.CurrentWeatherData
	Idx    int
}

// NewWidget creates and returns a new instance of the weather Widget.
func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	configKey := "weather"
	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget("Weather", configKey, true),

		Idx: 0,
	}

	widget.loadAPICredentials()

	widget.HelpfulWidget.SetView(widget.View)
	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Fetch retrieves OpenWeatherMap data from the OpenWeatherMap API.
// It takes a list of OpenWeatherMap city IDs.
// It returns a list of OpenWeatherMap CurrentWeatherData structs, one per valid city code.
func (widget *Widget) Fetch(cityIDs []int) []*owm.CurrentWeatherData {
	data := []*owm.CurrentWeatherData{}

	for _, cityID := range cityIDs {
		result, err := widget.currentWeather(widget.APIKey, cityID)
		if err == nil {
			data = append(data, result)
		}
	}

	return data
}

// Refresh fetches new data from the OpenWeatherMap API and loads the new data into the.
// widget's view for rendering
func (widget *Widget) Refresh() {
	if widget.apiKeyValid() {
		widget.Data = widget.Fetch(wtf.ToInts(wtf.Config.UList("wtf.mods.weather.cityids", widget.defaultCityCodes())))
	}

	widget.UpdateRefreshedAt()
	widget.display()
}

// Next displays data for the next city data in the list. If the current city is the last
// city, it wraps to the first city.
func (widget *Widget) Next() {
	widget.Idx = widget.Idx + 1
	if widget.Idx == len(widget.Data) {
		widget.Idx = 0
	}

	widget.display()
}

// Prev displays data for the previous city in the list. If the previous city is the first
// city, it wraps to the last city.
func (widget *Widget) Prev() {
	widget.Idx = widget.Idx - 1
	if widget.Idx < 0 {
		widget.Idx = len(widget.Data) - 1
	}

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) apiKeyValid() bool {
	if widget.APIKey == "" {
		return false
	}

	if len(widget.APIKey) != 32 {
		return false
	}

	return true
}

func (widget *Widget) currentData() *owm.CurrentWeatherData {
	if len(widget.Data) == 0 {
		return nil
	}

	if widget.Idx < 0 || widget.Idx >= len(widget.Data) {
		return nil
	}

	return widget.Data[widget.Idx]
}

func (widget *Widget) currentWeather(apiKey string, cityCode int) (*owm.CurrentWeatherData, error) {
	weather, err := owm.NewCurrent(
		wtf.Config.UString("wtf.mods.weather.tempUnit", "C"),
		wtf.Config.UString("wtf.mods.weather.language", "EN"),
		apiKey,
	)
	if err != nil {
		return nil, err
	}

	err = weather.CurrentByID(cityCode)
	if err != nil {
		return nil, err
	}

	return weather, nil
}

func (widget *Widget) defaultCityCodes() []interface{} {
	defaultArr := []int{3370352}

	var defaults = make([]interface{}, len(defaultArr))
	for i, d := range defaultArr {
		defaults[i] = d
	}

	return defaults
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "/":
		widget.ShowHelp()
		return nil
	case "h":
		widget.Prev()
		return nil
	case "l":
		widget.Next()
		return nil
	}

	switch event.Key() {
	case tcell.KeyLeft:
		widget.Prev()
		return nil
	case tcell.KeyRight:
		widget.Next()
		return nil
	default:
		return event
	}
}

// loadAPICredentials loads the API authentication credentials for this module
// First checks to see if they're in the config file. If not, checks the ENV var
func (widget *Widget) loadAPICredentials() {
	widget.APIKey = wtf.Config.UString(
		"wtf.mods.weather.apiKey",
		os.Getenv("WTF_OWM_API_KEY"),
	)
}
