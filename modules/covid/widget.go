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
func NewWidget(app *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	widget := &Widget{
		TextWidget: view.NewTextWidget(app, redrawChan, nil, settings.Common),

		settings: settings,
	}

	widget.View.SetScrollable(true)

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

// Display stats based on the user's locale
func (widget *Widget) displayStats(cases int) string {
	prntr, err := widget.settings.LocalizedPrinter()
	if err != nil {
		return err.Error()
	}

	return prntr.Sprintf("%d", cases)
}

func (widget *Widget) content() (string, string, bool) {
	title := defaultTitle
	if widget.CommonSettings().Title != "" {
		title = widget.CommonSettings().Title
	}

	cases, err := LatestCases()
	var covidStats string
	if err != nil {
		widget.err = err
	} else {
		// Display global stats
		covidStats = fmt.Sprintf("[%s]Global[white]\n", widget.settings.Colors.Subheading)
		covidStats += fmt.Sprintf("%s: %s\n", "Confirmed", widget.displayStats(cases.Latest.Confirmed))
		covidStats += fmt.Sprintf("%s: %s\n", "Deaths", widget.displayStats(cases.Latest.Deaths))
	}
	// Retrieve country stats if country codes are set in the config
	if len(widget.settings.countries) > 0 {
		countryCases, err := widget.LatestCountryCases(widget.settings.countries)
		if err != nil {
			widget.err = err
		} else {
			for i, name := range countryCases {
				covidStats += fmt.Sprintf("[%s]Country[white]: %s\n", widget.settings.Colors.Subheading, widget.settings.countries[i])
				covidStats += fmt.Sprintf("%s: %s\n", "Confirmed", widget.displayStats(name.Latest.Confirmed))
				covidStats += fmt.Sprintf("%s: %s\n", "Deaths", widget.displayStats(name.Latest.Deaths))
			}
		}
	}

	return title, covidStats, true
}
