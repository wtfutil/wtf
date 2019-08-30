package arpansagovau

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget
	location *location
	lastError error
	settings *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	locationData, err := GetLocationData(settings.city)
	widget := Widget {
		TextWidget: view.NewTextWidget(app, settings.common, false),
		location: locationData,
		lastError: err,
		settings: settings,
	}

	widget.View.SetWrap(true)

	return &widget
}

func (widget *Widget) content() (string, string, bool) {

	locationData, err := GetLocationData(widget.settings.city)
	widget.location = locationData
	widget.lastError = err

	if widget.lastError != nil {
		return widget.CommonSettings().Title, fmt.Sprintf("Err: %s", widget.lastError.Error()), true
	}

	return widget.CommonSettings().Title, formatLocationData(widget.location), true
}

func (widget *Widget) Refresh() {

	widget.Redraw(widget.content)
}

func formatLocationData(location *location) string {
	var level string
	var color string
	var content string

	if(location.name == "") {
		return "[red]No data?"
	}

	if(location.status != "ok") {
		content = "[red]Data unavailable for "
		content += location.name
		return content
	}

	switch {
		case location.index < 2.5:
			color = "[green]"
			level = " (LOW)"
		case location.index >= 2.5 && location.index < 5.5:
			color = "[yellow]"
			level = " (MODERATE)"
		case location.index >= 5.5 && location.index < 7.5:
			color = "[orange]"
			level = " (HIGH)"
		case location.index >= 7.5 && location.index < 10.5:
			color = "[red]"
			level = " (VERY HIGH)"
		case location.index >= 10.5:
			color = "[fuchsia]"
			level = " (EXTREME)"
	}

	content = "Location: "
	content += location.name
	content += "\nUV index: "
	content += color
	content += fmt.Sprintf("%.2f", location.index)
	content += level
	content += "[white]\nLocal time: "
	content += location.time
	content += " "
	content += location.date
	content += "\nDetector status: "
	content += location.status

	return content
}
