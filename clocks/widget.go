package clocks

import (
	"strings"
	"time"

	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	clockColl  ClockCollection
	dateFormat string
	timeFormat string
}

func NewWidget(app *tview.Application) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, "World Clocks", "clocks", false),
	}

	widget.clockColl = widget.buildClockCollection(wtf.Config.UMap("wtf.mods.clocks.locations"))

	widget.dateFormat = wtf.Config.UString("wtf.mods.clocks.dateFormat", wtf.SimpleDateFormat)
	widget.timeFormat = wtf.Config.UString("wtf.mods.clocks.timeFormat", wtf.SimpleTimeFormat)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.display(widget.clockColl.Sorted(), widget.dateFormat, widget.timeFormat)
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
