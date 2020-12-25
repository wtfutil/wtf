package covid

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

// Widget is the struct that defines this module widget
type Widget struct {
	view.TextWidget

	settings *Settings
	err      error
}

// NewWidget creates a new widget for this module
func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := &Widget{
		TextWidget: view.NewTextWidget(app, nil, settings.Common),

		settings: settings,
	}

	return widget
}

// Refresh checks if this module widget is disabled
func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.Redraw(widget.content)
}

// Render renders this module widget
func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	title := defaultTitle
	if widget.CommonSettings().Title != "" {
		title = widget.CommonSettings().Title
	}

	cases, err := LatestCases()
	var covidInfo string
	if err != nil {
		widget.err = err
	} else {
		covidInfo = fmt.Sprintf("[%s]Gobal[white]\n", widget.settings.Colors.Subheading)
		covidInfo += fmt.Sprintf("%s: %d\n", "Confirmed", cases.Latest.Confirmed)
		covidInfo += fmt.Sprintf("%s: %d\n", "Deaths", cases.Latest.Confirmed)
	}
	// Retrieve country stats if the country code is set in the config
	if widget.settings.country != "" {
		countryCases, err := widget.LatestCountryCases(widget.settings.country)
		if err != nil {
			widget.err = err
		} else {
			covidInfo += "\n"
			covidInfo += fmt.Sprintf("[%s]Country stats[white]: %s\n", widget.settings.Colors.Subheading, widget.settings.country)
			covidInfo += fmt.Sprintf("%s: %d\n", "Confirmed", countryCases.LatestCountryCases.Confirmed)
			covidInfo += fmt.Sprintf("%s: %d\n", "Deaths", countryCases.LatestCountryCases.Deaths)
		}
	}
	return title, covidInfo, true
}
