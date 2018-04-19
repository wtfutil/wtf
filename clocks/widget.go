package clocks

import (
	"time"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

const TimeFormat = "15:04 MST"
const DateFormat = "Jan 2"

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" 🕗 World Clocks ", "clocks"),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	clockColl := widget.buildClockCollection(Config.UMap("wtf.mods.clocks.locations"))

	widget.View.Clear()
	widget.display(clockColl.Sorted())
	widget.RefreshedAt = time.Now()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) buildClockCollection(locData map[string]interface{}) ClockCollection {
	clockColl := ClockCollection{}

	for label, locStr := range locData {
		timeLoc, err := time.LoadLocation(locStr.(string))
		if err != nil {
			continue
		}

		clock := Clock{
			Label:     label,
			LocalTime: time.Now().In(timeLoc),
			Timezone:  locStr.(string),
		}

		clockColl.Clocks = append(clockColl.Clocks, clock)
	}

	return clockColl
}
