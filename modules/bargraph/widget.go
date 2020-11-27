package bargraph

/**************
This is a demo bargraph that just populates some random date/val data
*/

import (
	"math/rand"
	"time"

	"github.com/rivo/tview"

	"github.com/wtfutil/wtf/view"
)

// Widget define wtf widget to register widget later
type Widget struct {
	view.BarGraph

	tviewApp *tview.Application
}

// NewWidget Make new instance of widget
func NewWidget(tviewApp *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		BarGraph: view.NewBarGraph(tviewApp, "Sample Bar Graph", settings.Common),

		tviewApp: tviewApp,
	}

	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// MakeGraph - Load the dead drop stats
func MakeGraph(widget *Widget) {

	//this could come from config
	const lineCount = 8
	var stats [lineCount]view.Bar

	barTime := time.Now()
	for i := 0; i < lineCount; i++ {
		barTime = barTime.Add(time.Minute)

		bar := view.Bar{
			Label:      barTime.Format("15:04"),
			Percent:    rand.Intn(100-5) + 5,
			LabelColor: "red",
		}

		stats[i] = bar
	}

	widget.BarGraph.BuildBars(stats[:])

}

// Refresh & update after interval time
func (widget *Widget) Refresh() {

	if widget.Disabled() {
		return
	}

	widget.View.Clear()

	widget.tviewApp.QueueUpdateDraw(func() {
		display(widget)
	})

}

/* -------------------- Unexported Functions -------------------- */

func display(widget *Widget) {
	MakeGraph(widget)
}
