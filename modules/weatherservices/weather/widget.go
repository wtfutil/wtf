package weather

import (
	owm "github.com/briandowns/openweathermap"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

// Widget is the container for weather data.
type Widget struct {
	view.MultiSourceWidget
	view.TextWidget

	// APIKey   string
	Data []*owm.CurrentWeatherData

	pages    *tview.Pages
	settings *Settings
}

// NewWidget creates and returns a new instance of the weather Widget
func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		MultiSourceWidget: view.NewMultiSourceWidget(settings.Common, "cityid", "cityids"),
		TextWidget:        view.NewTextWidget(tviewApp, pages, settings.Common),

		pages:    pages,
		settings: settings,
	}

	widget.initializeKeyboardControls()

	widget.SetDisplayFunction(widget.display)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Fetch retrieves OpenWeatherMap data from the OpenWeatherMap API.
// It takes a list of OpenWeatherMap city IDs.
// It returns a list of OpenWeatherMap CurrentWeatherData structs, one per valid city code.
func (widget *Widget) Fetch(cityIDs []int) []*owm.CurrentWeatherData {
	data := []*owm.CurrentWeatherData{}

	for _, cityID := range cityIDs {
		result, err := widget.currentWeather(cityID)
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
		widget.Data = widget.Fetch(utils.ToInts(widget.settings.cityIDs))
	}

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) apiKeyValid() bool {
	if widget.settings.apiKey == "" {
		return false
	}

	if len(widget.settings.apiKey) != 32 {
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

func (widget *Widget) currentWeather(cityCode int) (*owm.CurrentWeatherData, error) {
	weather, err := owm.NewCurrent(
		widget.settings.tempUnit,
		widget.settings.language,
		widget.settings.apiKey,
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
