package clocks

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	app        *tview.Application
	clockColl  ClockCollection
	dateFormat string
	timeFormat string
	settings   *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, nil, settings.common),

		app:        app,
		settings:   settings,
		dateFormat: settings.dateFormat,
		timeFormat: settings.timeFormat,
	}

	widget.clockColl = widget.buildClockCollection()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {
	widget.app.QueueUpdateDraw(func() {
		sortedClocks := widget.clockColl.Sorted(widget.settings.sort)
		widget.display(sortedClocks, widget.dateFormat, widget.timeFormat)
	})
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) buildClockCollection() ClockCollection {
	clockColl := ClockCollection{}

	clockColl.Clocks = widget.settings.locations

	return clockColl
}
