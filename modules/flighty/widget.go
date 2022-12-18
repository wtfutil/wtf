package flighty

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.MultiSourceWidget
	view.TextWidget

	Data []*FlightData

	pages    *tview.Pages
	settings *Settings
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		MultiSourceWidget: view.NewMultiSourceWidget(settings.Common, "", "aircraft"),
		TextWidget:        view.NewTextWidget(tviewApp, redrawChan, pages, settings.Common),

		pages:    pages,
		settings: settings,
	}

	widget.initializeKeyboardControls()

	widget.SetDisplayFunction(widget.display)

	return widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Fetch(aircraft []string) []*FlightData {
	data := []*FlightData{}

	for _, plane := range aircraft {
		result, err := widget.flightData(plane)
		if err == nil {
			data = append(data, result)
		}
	}

	return data
}

// Refresh fetches new data from the OpenWeatherMap API and loads the new data into the.
// widget's view for rendering
func (widget *Widget) Refresh() {
	if widget.authenticationCredentialsValid() {
		widget.Data = widget.Fetch(widget.settings.aircraft)
	}

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) authenticationCredentialsValid() bool {
	if widget.settings.username == "" {
		return false
	}

	if len(widget.settings.password) != 32 {
		return false
	}

	return true
}

func (widget *Widget) flightData(aircraft string) (*FlightData, error) {
	flightData, err := NewFlightData(aircraft)
	return flightData, err
}
