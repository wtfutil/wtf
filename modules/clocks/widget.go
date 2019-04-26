package clocks

import (
	"strings"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	app        *tview.Application
	clockColl  ClockCollection
	dateFormat string
	timeFormat string
	settings   *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, settings.common, false),

		app:        app,
		settings:   settings,
		dateFormat: settings.dateFormat,
		timeFormat: settings.timeFormat,
	}

	widget.clockColl = widget.buildClockCollection(settings.locations)

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

func (widget *Widget) buildClockCollection(locData map[string]interface{}) ClockCollection {
	clockColl := ClockCollection{}

	for label, locStr := range locData {
		timeLoc, err := time.LoadLocation(widget.sanitizeLocation(locStr.(string)))
		if err != nil {
			continue
		}

		clockColl.Clocks = append(clockColl.Clocks, NewClock(label, timeLoc))
	}

	return clockColl
}

func (widget *Widget) sanitizeLocation(locStr string) string {
	return strings.Replace(locStr, " ", "_", -1)
}
