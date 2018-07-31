package clocks

import (
	"strings"
	"time"

	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	clockColl ClockCollection
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget("World Clocks", "clocks", false),
	}

	widget.clockColl = widget.buildClockCollection(wtf.Config.UMap("wtf.mods.clocks.locations"))

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.UpdateRefreshedAt()
	widget.display(widget.clockColl.Sorted())
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
